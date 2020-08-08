package staking

import (
	"encoding/json"

	"github.com/gogo/protobuf/grpc"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/KiraCore/cosmos-sdk/client"
	"github.com/KiraCore/cosmos-sdk/codec"
	types2 "github.com/KiraCore/cosmos-sdk/codec/types"
	sdk "github.com/KiraCore/cosmos-sdk/types"
	"github.com/KiraCore/cosmos-sdk/types/module"
	"github.com/KiraCore/cosmos-sdk/x/staking"
	stakingkeeper "github.com/KiraCore/cosmos-sdk/x/staking/keeper"
	"github.com/KiraCore/cosmos-sdk/x/staking/types"
	"github.com/KiraCore/sekai/x/staking/keeper"
	cumstomtypes "github.com/KiraCore/sekai/x/staking/types"

	"github.com/KiraCore/sekai/x/staking/client/cli"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct {
}

func (b AppModuleBasic) Name() string {
	return cumstomtypes.ModuleName
}

func (b AppModuleBasic) RegisterInterfaces(registry types2.InterfaceRegistry) {
	cumstomtypes.RegisterInterfaces(registry)
}

func (b AppModuleBasic) DefaultGenesis(marshaler codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (b AppModuleBasic) ValidateGenesis(marshaler codec.JSONMarshaler, config client.TxEncodingConfig, message json.RawMessage) error {
	return nil
}

func (b AppModuleBasic) RegisterRESTRoutes(context client.Context, router *mux.Router) {
}

func (b AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxClaimValidatorCmd()
}

func (b AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

// RegisterCodec registers the staking module's types for the given codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	cumstomtypes.RegisterCodec(cdc)
}

// AppModule extends the cosmos SDK staking.
type AppModule struct {
	staking.AppModule

	stakingKeeper       stakingkeeper.Keeper
	customStakingKeeper keeper.Keeper
}

// NewHandler returns an sdk.Handler for the staking module.
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.stakingKeeper, am.customStakingKeeper)
}

// RegisterQueryService registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterQueryService(server grpc.Server) {
	querier := Querier{}
	cumstomtypes.RegisterQueryServer(server, querier)
}

// NewAppModule returns a new Custom Staking module.
func NewAppModule(
	cdc codec.Marshaler,
	keeper stakingkeeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
) AppModule {
	return AppModule{
		AppModule:     staking.NewAppModule(cdc, keeper, ak, bk),
		stakingKeeper: keeper,
	}
}
