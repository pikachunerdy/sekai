package interx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/KiraCore/sekai/INTERX/common"
	"github.com/KiraCore/sekai/INTERX/config"
	"github.com/KiraCore/sekai/INTERX/types"
	"github.com/KiraCore/sekai/INTERX/types/kira"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	tmRPCTypes "github.com/tendermint/tendermint/rpc/core/types"
)

// RegisterValidatorsQueryRoutes registers validators query routers.
func RegisterValidatorsQueryRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QueryConsensus, QueryConsensus(gwCosmosmux, rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryValidators, QueryValidators(gwCosmosmux, rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryValidatorInfos, QueryValidatorInfos(gwCosmosmux, rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryConsensus, "This is an API to query consensus.", true)
	common.AddRPCMethod("GET", config.QueryValidators, "This is an API to query validators.", true)
	common.AddRPCMethod("GET", config.QueryValidatorInfos, "This is an API to query validator infos.", true)
}

const (
	// Undefined status
	Undefined string = "UNDEFINED"
	// Active status
	Active string = "ACTIVE"
	// Inactive status
	Inactive string = "INACTIVE"
	// Paused status
	Paused string = "PAUSED"
	// Jailed status
	Jailed string = "JAILED"
)

func queryValidatorsHandle(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	tempRequest := r.Clone(r.Context())

	queries := r.URL.Query()
	address := queries["address"]
	valkey := queries["valkey"]
	pubkey := queries["pubkey"]
	moniker := queries["moniker"]
	status := queries["status"]
	key := queries["key"]
	offset := queries["offset"]
	limit := queries["limit"]
	proposer := queries["proposer"]
	countTotal := queries["count_total"]
	all := queries["all"]

	isQueryAll := false

	var events = make([]string, 0, 9)
	if len(key) == 1 {
		events = append(events, fmt.Sprintf("pagination.key=%s", key[0]))
	}
	if len(offset) == 1 {
		events = append(events, fmt.Sprintf("pagination.offset=%s", offset[0]))
	}
	if len(limit) == 1 {
		events = append(events, fmt.Sprintf("pagination.limit=%s", limit[0]))
	}
	if len(countTotal) == 1 {
		events = append(events, fmt.Sprintf("pagination.count_total=%s", countTotal[0]))
	}
	if len(address) == 1 {
		events = append(events, fmt.Sprintf("address=%s", address[0]))
	}
	if len(valkey) == 1 {
		events = append(events, fmt.Sprintf("valkey=%s", valkey[0]))
	}
	if len(pubkey) == 1 {
		events = append(events, fmt.Sprintf("pubkey=%s", pubkey[0]))
	}
	if len(proposer) == 1 {
		events = append(events, fmt.Sprintf("proposer=%s", proposer[0]))
	}
	if len(moniker) == 1 {
		events = append(events, fmt.Sprintf("moniker=%s", moniker[0]))
	}
	if len(status) == 1 {
		events = append(events, fmt.Sprintf("status=%s", status[0]))
	}
	if len(all) == 1 {
		events = append(events, fmt.Sprintf("all=%s", all[0]))
		isQueryAll = all[0] == "true"
	}

	r.URL.RawQuery = strings.Join(events, "&")

	success, failure, statusCode := common.ServeGRPC(r, gwCosmosmux)
	if success != nil {
		result := struct {
			Validators []types.QueryValidator `json:"validators,omitempty"`
			Actors     []string               `json:"actors,omitempty"`
			Pagination interface{}            `json:"pagination,omitempty"`
		}{}

		byteData, err := json.Marshal(success)
		if err != nil {
			common.GetLogger().Error("[query-validators] Invalid response format: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		err = json.Unmarshal(byteData, &result)
		if err != nil {
			common.GetLogger().Error("[query-validators] Invalid response format: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		for index, validator := range result.Validators {
			pubkey, _ := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, validator.Pubkey)

			newReq := tempRequest.Clone(tempRequest.Context())
			newReq.URL.Path = config.QueryValidatorInfos + "/" + sdk.GetConsAddress(pubkey).String()

			signInfoRes, _, _ := common.ServeGRPC(newReq, gwCosmosmux)

			if signInfoRes != nil {
				signInfoResponse := struct {
					ValSigningInfo types.ValidatorSigningInfo `json:"val_signing_info,omitempty"`
				}{}

				byteData, err := json.Marshal(signInfoRes)
				if err != nil {
					common.GetLogger().Error("[query-validator-signinginfo] Invalid response format: ", err)
					return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
				}

				err = json.Unmarshal(byteData, &signInfoResponse)
				if err != nil {
					common.GetLogger().Error("[query-validator-signinginfo] Invalid response format: ", err)
					return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
				}

				result.Validators[index].Mischance = signInfoResponse.ValSigningInfo.MissedBlocksCounter
				result.Validators[index].StartHeight = signInfoResponse.ValSigningInfo.StartHeight
				result.Validators[index].IndexOffset = signInfoResponse.ValSigningInfo.IndexOffset
				result.Validators[index].InactiveUntil = signInfoResponse.ValSigningInfo.InactiveUntil
				result.Validators[index].Tombstoned = signInfoResponse.ValSigningInfo.Tombstoned
				result.Validators[index].MissedBlocksCounter = signInfoResponse.ValSigningInfo.MissedBlocksCounter
			}
		}

		sort.Sort(types.QueryValidators(result.Validators))
		for index := range result.Validators {
			result.Validators[index].Top = index + 1
		}

		if isQueryAll {

			allValidators := types.AllValidators{}

			allValidators.Validators = result.Validators
			allValidators.Waiting = make([]string, 0)
			for _, actor := range result.Actors {
				isWaiting := true
				for _, validator := range result.Validators {
					if validator.Address == actor {
						isWaiting = false
						break
					}
				}

				if isWaiting {
					allValidators.Waiting = append(allValidators.Waiting, actor)
				}
			}

			allValidators.Status.TotalValidators = len(result.Validators)
			allValidators.Status.WaitingValidators = len(allValidators.Waiting)

			allValidators.Status.ActiveValidators = 0
			allValidators.Status.PausedValidators = 0
			allValidators.Status.InactiveValidators = 0
			allValidators.Status.JailedValidators = 0
			for _, validator := range result.Validators {
				if validator.Status == Active {
					allValidators.Status.ActiveValidators++
				}
				if validator.Status == Inactive {
					allValidators.Status.InactiveValidators++
				}
				if validator.Status == Paused {
					allValidators.Status.PausedValidators++
				}
				if validator.Status == Jailed {
					allValidators.Status.JailedValidators++
				}
			}

			return allValidators, nil, statusCode
		}

		return result, nil, statusCode
	}

	return success, failure, statusCode
}

// QueryValidators is a function to list validators.
func QueryValidators(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		if !common.RPCMethods["GET"][config.QueryValidators].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryValidators].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[encode-transaction] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryValidatorsHandle(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, true)
	}
}

func queryValidatorInfosHandle(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	queries := r.URL.Query()
	key := queries["key"]
	offset := queries["offset"]
	limit := queries["limit"]
	countTotal := queries["count_total"]
	all := queries["all"]

	var events = make([]string, 0, 9)
	if len(key) == 1 {
		events = append(events, fmt.Sprintf("pagination.key=%s", key[0]))
	}
	if len(offset) == 1 {
		events = append(events, fmt.Sprintf("pagination.offset=%s", offset[0]))
	}
	if len(limit) == 1 {
		events = append(events, fmt.Sprintf("pagination.limit=%s", limit[0]))
	}
	if len(countTotal) == 1 {
		events = append(events, fmt.Sprintf("pagination.count_total=%s", countTotal[0]))
	}
	if len(all) == 1 {
		events = append(events, fmt.Sprintf("all=%s", all[0]))
	}

	r.URL.RawQuery = strings.Join(events, "&")

	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryValidatorInfos is a function to list validators information.
func QueryValidatorInfos(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		if !common.RPCMethods["GET"][config.QueryValidatorInfos].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryValidatorInfos].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[encode-transaction] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryValidatorInfosHandle(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, true)
	}
}

func GetValidator(consAddrHex string) string {
	bytes, err := hex.DecodeString(consAddrHex)
	if err != nil {
		return ""
	}

	return sdk.AccAddress(bytes).String()
}

func queryConsensusHandle(r *http.Request, gwCosmosmux *runtime.ServeMux, rpcAddr string) (interface{}, interface{}, int) {
	// Query All Validators
	var AllValidators []types.QueryValidator = make([]types.QueryValidator, 0)

	r.URL.RawQuery = "all=true"
	r.URL.Path = config.QueryValidators
	success, failure, statusCode := common.ServeGRPC(r, gwCosmosmux)
	if success != nil {
		result := struct {
			Validators []types.QueryValidator `json:"validators,omitempty"`
			Actors     []string               `json:"actors,omitempty"`
			Pagination interface{}            `json:"pagination,omitempty"`
		}{}

		byteData, err := json.Marshal(success)
		if err != nil {
			common.GetLogger().Error("[query-validators] Invalid response format: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		err = json.Unmarshal(byteData, &result)
		if err != nil {
			common.GetLogger().Error("[query-validators] Invalid response format: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		AllValidators = result.Validators
	} else {
		return success, failure, statusCode
	}

	// Query consensus
	success, failure, statusCode = common.MakeGetRequest(rpcAddr, "/dump_consensus_state", "")
	if success != nil {
		consensus := tmRPCTypes.ResultDumpConsensusState{}

		byteData, err := json.Marshal(success)
		if err != nil {
			common.GetLogger().Error("[query-consensus] Invalid response format: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		err = json.Unmarshal(byteData, &consensus)
		if err != nil {
			common.GetLogger().Error("[query-consensus] Invalid response format: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		type Vote struct {
			Round              int64             `json:"round"`
			Prevotes           map[string]string `json:"prevotes"`
			PrevotesBitArray   string            `json:"prevotes_bit_array"`
			Precommits         map[string]string `json:"precommits"`
			PrecommitsBitArray string            `json:"precommits_bit_array"`
		}

		type Validator struct {
			Address  string `json:"address"`
			Valkey   string `json:"valkey"`
			Pubkey   string `json:"pubkey"`
			Proposer string `json:"proposer"`
		}

		response := struct {
			RawData          tmRPCTypes.ResultDumpConsensusState `json:"raw_data,omitempty"`
			ConsensusStopped bool                                `json:"consensus_stopped"`
			Validators       []Validator                         `json:"validators"`
			Proposer         Validator                           `json:"proposer"`
			Votes            []Vote                              `json:"votes"`
		}{}

		response.RawData = consensus

		roundState := kira.RoundState{}

		err = json.Unmarshal(consensus.RoundState, &roundState)
		if err != nil {
			common.GetLogger().Error("[query-consensus] Invalid round state: ", err)
			return common.ServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		response.ConsensusStopped = common.IsConsensusStopped(len(roundState.Validators.Validators))
		response.Validators = make([]Validator, 0)
		for i := range roundState.Validators.Validators {
			for _, validator := range AllValidators {
				if validator.Proposer == roundState.Validators.Validators[i].Address {
					response.Validators = append(response.Validators, Validator{
						Address:  validator.Address,
						Pubkey:   validator.Pubkey,
						Valkey:   validator.Valkey,
						Proposer: validator.Proposer,
					})
					break
				}
			}
		}
		for _, validator := range AllValidators {
			if validator.Proposer == roundState.Validators.Proposer.Address {
				response.Proposer = Validator{
					Address:  validator.Address,
					Pubkey:   validator.Pubkey,
					Valkey:   validator.Valkey,
					Proposer: validator.Proposer,
				}
				break
			}
		}

		response.Votes = make([]Vote, 0)

		for _, vote := range roundState.Votes {
			newVote := Vote{}

			newVote.Round = vote.Round
			newVote.PrevotesBitArray = vote.PrevotesBitArray
			newVote.PrecommitsBitArray = vote.PrecommitsBitArray

			newVote.Prevotes = make(map[string]string)
			newVote.Precommits = make(map[string]string)

			for i := range vote.Prevotes {
				newVote.Prevotes[response.Validators[i].Address] = vote.Prevotes[i]
			}
			for i := range vote.Precommits {
				newVote.Precommits[response.Validators[i].Address] = vote.Precommits[i]
			}

			response.Votes = append(response.Votes, newVote)
		}

		return response, failure, statusCode
	}

	return success, failure, statusCode
}

// QueryConsensus is a function to query consensus.
func QueryConsensus(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		if !common.RPCMethods["GET"][config.QueryConsensus].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryConsensus].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[encode-transaction] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryConsensusHandle(r, gwCosmosmux, rpcAddr)
		}

		common.WrapResponse(w, request, *response, statusCode, true)
	}
}
