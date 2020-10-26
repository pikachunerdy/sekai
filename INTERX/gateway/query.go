package gateway

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

const (
	queryStatus     = "/api/cosmos/status"
	queryRPCMethods = "/api/rpc_methods"
	queryFunctions  = "/api/functions"
)

// RegisterQueryRoutes registers query routers.
func RegisterQueryRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(queryStatus, QueryStatusRequest(rpcAddr)).Methods(GET)
	r.HandleFunc(queryRPCMethods, QueryRPCMethods(rpcAddr)).Methods(GET)
	r.HandleFunc(queryFunctions, QueryFunctions(rpcAddr)).Methods(GET)

	AddRPCMethod(GET, queryStatus, "This is an API to query status.", true)
	AddRPCMethod(GET, queryFunctions, "This is an API to list functions and metadata.", true)
}

func queryStatusHandle(rpcAddr string) (interface{}, interface{}, int) {
	return MakeGetRequest(rpcAddr, "/status", "")
}

// QueryStatusRequest is a function to query status.
func QueryStatusRequest(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := GetInterxRequest(r)
		response := GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		if !rpcMethods[GET][queryStatus].Enabled {
			response.Response, response.Error, statusCode = ServeError(0, "", "", http.StatusForbidden)
		} else {
			if rpcMethods[GET][queryStatus].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					WrapResponse(w, request, *response, statusCode, false)
					return
				}
			}

			response.Response, response.Error, statusCode = queryStatusHandle(rpcAddr)
		}

		WrapResponse(w, request, *response, statusCode, rpcMethods[GET][queryStatus].CachingEnabled)
	}
}

// QueryRPCMethods is a function to query RPC methods.
func QueryRPCMethods(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := GetInterxRequest(r)
		response := GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		response.Response = rpcMethods

		WrapResponse(w, request, *response, statusCode, false)
	}
}

// QueryFunctions is a function to list functions and metadata.
func QueryFunctions(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := GetInterxRequest(r)
		response := GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		response.Response = rpcMethods

		WrapResponse(w, request, *response, statusCode, false)
	}
}
