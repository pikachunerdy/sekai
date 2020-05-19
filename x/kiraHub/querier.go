package kiraHub

import (
	sdk "github.com/KiraCore/cosmos-sdk/types"
	constants "github.com/KiraCore/sekai/x/kiraHub/constants"
	"github.com/KiraCore/sekai/x/kiraHub/queries/listOrderBooks"
	"github.com/KiraCore/cosmos-sdk/types/errors"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryListOrderBooks = "listOrderBooks"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(context sdk.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
		switch path[0] {

		case QueryListOrderBooks:
			return listOrderBooks.QueryGetOrderBooks(context, path[1:], requestQuery, keeper.getCreateOrderBookKeeper())

		default:
			return nil, errors.Wrapf(constants.UnknownQueryCode, "%v", path[0])
		}
	}
}
