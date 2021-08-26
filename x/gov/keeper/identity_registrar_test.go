package keeper_test

import (
	"testing"
	"time"

	simapp "github.com/KiraCore/sekai/app"
	"github.com/KiraCore/sekai/x/gov/keeper"
	"github.com/KiraCore/sekai/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestKeeper_CheckIfWithinAddressArray(t *testing.T) {
	addr1 := sdk.AccAddress("foo1________________")
	addr2 := sdk.AccAddress("foo2________________")
	addr3 := sdk.AccAddress("foo3________________")
	addr4 := sdk.AccAddress("foo4________________")

	array := []sdk.AccAddress{addr1, addr2, addr3}

	require.True(t, keeper.CheckIfWithinAddressArray(addr1, array))
	require.True(t, keeper.CheckIfWithinAddressArray(addr2, array))
	require.True(t, keeper.CheckIfWithinAddressArray(addr3, array))
	require.False(t, keeper.CheckIfWithinAddressArray(addr4, array))
}

func TestKeeper_LastIdentityRecordId(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	lastRecordId := app.CustomGovKeeper.GetLastIdentityRecordId(ctx)
	require.Equal(t, lastRecordId, uint64(0))

	app.CustomGovKeeper.SetLastIdentityRecordId(ctx, 5)

	lastRecordId = app.CustomGovKeeper.GetLastIdentityRecordId(ctx)
	require.Equal(t, lastRecordId, uint64(5))
}

func TestKeeper_LastIdRecordVerifyRequestId(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	lastRecordId := app.CustomGovKeeper.GetLastIdRecordVerifyRequestId(ctx)
	require.Equal(t, lastRecordId, uint64(0))

	app.CustomGovKeeper.SetLastIdRecordVerifyRequestId(ctx, 5)

	lastRecordId = app.CustomGovKeeper.GetLastIdRecordVerifyRequestId(ctx)
	require.Equal(t, lastRecordId, uint64(5))
}

func TestKeeper_IdentityRecordBasicFlow(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	// try to get non existent record
	record := app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.Nil(t, record)

	// create a new record and check if set correctly
	addr1 := sdk.AccAddress("foo1________________")
	addr2 := sdk.AccAddress("foo2________________")
	addr3 := sdk.AccAddress("foo3________________")
	newRecord := types.IdentityRecord{
		Id:        1,
		Address:   addr1,
		Key:       "key",
		Value:     "value",
		Date:      time.Now().UTC(),
		Verifiers: []sdk.AccAddress{addr2, addr3},
	}
	app.CustomGovKeeper.SetIdentityRecord(ctx, newRecord)
	record = app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.NotNil(t, record)
	require.Equal(t, *record, newRecord)

	// check no panics
	app.CustomGovKeeper.DeleteIdentityRecordById(ctx, 0)

	// remove existing id and check
	app.CustomGovKeeper.DeleteIdentityRecordById(ctx, 1)
	record = app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.Nil(t, record)
}

func TestKeeper_IdentityRecordAddEditRemove(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	// create a new record and check if set correctly
	addr1 := sdk.AccAddress("foo1________________")
	addr2 := sdk.AccAddress("foo2________________")
	addr3 := sdk.AccAddress("foo3________________")
	infos := make(map[string]string)
	infos["key"] = "value"
	now := time.Now().UTC()
	ctx = ctx.WithBlockTime(now)
	err := app.CustomGovKeeper.RegisterIdentityRecords(ctx, addr1, types.WrapInfos(infos))
	require.NoError(t, err)

	record := app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.NotNil(t, record)
	expectedRecord := types.IdentityRecord{
		Id:      1,
		Address: addr1,
		Key:     "key",
		Value:   "value",
		Date:    now,
	}
	require.Equal(t, *record, expectedRecord)

	cacheCtx, _ := ctx.CacheContext()
	err = app.CustomGovKeeper.RegisterIdentityRecords(cacheCtx, addr1, types.WrapInfos(infos))
	require.Error(t, err)
	err = app.CustomGovKeeper.RegisterIdentityRecords(ctx, addr2, types.WrapInfos(infos))
	require.NoError(t, err)

	records := app.CustomGovKeeper.GetAllIdentityRecords(ctx)
	require.Len(t, records, 2)
	records = app.CustomGovKeeper.GetIdRecordsByAddress(ctx, addr1)
	require.NotNil(t, records)
	records = app.CustomGovKeeper.GetIdRecordsByAddress(ctx, addr2)
	require.NotNil(t, records)

	// remove existing id and check
	app.CustomGovKeeper.DeleteIdentityRecordById(ctx, 2)
	records = app.CustomGovKeeper.GetAllIdentityRecords(ctx)
	require.Len(t, records, 1)
	records = app.CustomGovKeeper.GetIdRecordsByAddress(ctx, addr1)
	require.NotNil(t, record)
	records = app.CustomGovKeeper.GetIdRecordsByAddress(ctx, addr2)
	require.Nil(t, record)

	infos["key1"] = "value1"
	now = now.Add(time.Second)
	ctx = ctx.WithBlockTime(now)

	// try deleting with other owner
	err = app.CustomGovKeeper.DeleteIdentityRecords(ctx, addr3, []string{"key1"})
	require.Error(t, err)

	// set verifier of identity record
	app.CustomGovKeeper.RegisterIdentityRecords(ctx, addr1, types.WrapInfos(infos))

	err = app.CustomGovKeeper.RegisterIdentityRecords(ctx, addr1, types.WrapInfos(infos))
	require.NoError(t, err)
	record = app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.NotNil(t, record)
	require.Equal(t, *record, types.IdentityRecord{
		Id:      1,
		Address: addr1,
		Key:     "key1",
		Value:   "value1",
		Date:    now,
	})
	records = app.CustomGovKeeper.GetAllIdentityRecords(ctx)
	require.Len(t, records, 1)
	records = app.CustomGovKeeper.GetIdRecordsByAddress(ctx, addr1)
	require.NotNil(t, record)
	records = app.CustomGovKeeper.GetIdRecordsByAddress(ctx, addr2)
	require.Nil(t, record)
}

func TestKeeper_IdentityRecordApproveFlow(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	// create a new record and check if set correctly
	addr1 := sdk.AccAddress("foo1________________")
	addr2 := sdk.AccAddress("foo2________________")
	addr3 := sdk.AccAddress("foo3________________")
	app.BankKeeper.SetBalance(ctx, addr1, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000))
	app.BankKeeper.SetBalance(ctx, addr2, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000))
	app.BankKeeper.SetBalance(ctx, addr3, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000))

	infos := make(map[string]string)
	infos["key"] = "value"
	now := time.Now().UTC()
	ctx = ctx.WithBlockTime(now)
	err := app.CustomGovKeeper.RegisterIdentityRecords(ctx, addr1, types.WrapInfos(infos))
	require.NoError(t, err)

	record := app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.NotNil(t, record)
	expectedRecord := types.IdentityRecord{
		Id:      1,
		Address: addr1,
		Key:     "key",
		Value:   "value",
		Date:    now,
	}
	require.Equal(t, *record, expectedRecord)

	ctxCache, _ := ctx.CacheContext()
	err = app.CustomGovKeeper.RegisterIdentityRecords(ctxCache, addr1, types.WrapInfos(infos))
	require.Error(t, err)
	err = app.CustomGovKeeper.RegisterIdentityRecords(ctx, addr2, types.WrapInfos(infos))
	require.NoError(t, err)

	// bigger tip than balance
	ctxCache, _ = ctx.CacheContext()
	reqId, err := app.CustomGovKeeper.RequestIdentityRecordsVerify(ctxCache, addr1, addr3, []uint64{1}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))
	require.Equal(t, reqId, uint64(0))
	require.Error(t, err)

	// request id record 1 to addr3 by addr1
	reqId, err = app.CustomGovKeeper.RequestIdentityRecordsVerify(ctx, addr1, addr3, []uint64{1}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10))
	require.Equal(t, reqId, uint64(1))
	require.NoError(t, err)
	request := app.CustomGovKeeper.GetIdRecordsVerifyRequest(ctx, 1)
	require.NotNil(t, request)
	require.Equal(t, *request, types.IdentityRecordsVerify{
		Id:        1,
		Address:   addr1,
		Verifier:  addr3,
		RecordIds: []uint64{1},
		Tip:       sdk.NewInt64Coin(sdk.DefaultBondDenom, 10),
	})
	coins := app.BankKeeper.GetAllBalances(ctx, addr1)
	require.Equal(t, coins, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 9990)})
	coins = app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress(types.ModuleName))
	require.Equal(t, coins, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 10)})
	err = app.CustomGovKeeper.HandleIdentityRecordsVerifyRequest(ctx, addr3, 1, true)
	require.NoError(t, err)
	coins = app.BankKeeper.GetAllBalances(ctx, addr3)
	require.Equal(t, coins, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 10010)})
	coins = app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress(types.ModuleName))
	require.Equal(t, coins, sdk.Coins(nil))
	record = app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.NotNil(t, record)
	require.Equal(t, *record, types.IdentityRecord{
		Id:        1,
		Address:   addr1,
		Key:       "key",
		Value:     "value",
		Date:      now,
		Verifiers: []sdk.AccAddress{addr3},
	})
	request = app.CustomGovKeeper.GetIdRecordsVerifyRequest(ctx, 1)
	require.Nil(t, request)

	// request id record 1 to addr3 by addr1 again
	reqId, err = app.CustomGovKeeper.RequestIdentityRecordsVerify(ctx, addr1, addr3, []uint64{1}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10))
	require.Equal(t, reqId, uint64(2))
	require.NoError(t, err)
	coins = app.BankKeeper.GetAllBalances(ctx, addr1)
	require.Equal(t, coins, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 9980)})
	err = app.CustomGovKeeper.HandleIdentityRecordsVerifyRequest(ctx, addr3, 2, true)
	require.NoError(t, err)
	coins = app.BankKeeper.GetAllBalances(ctx, addr3)
	require.Equal(t, coins, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 10020)})
	record = app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.NotNil(t, record)
	require.Equal(t, *record, types.IdentityRecord{
		Id:        1,
		Address:   addr1,
		Key:       "key",
		Value:     "value",
		Date:      now,
		Verifiers: []sdk.AccAddress{addr3},
	})

	// request id record 1 and 2 to addr3 by addr1 again
	reqId, err = app.CustomGovKeeper.RequestIdentityRecordsVerify(ctx, addr1, addr3, []uint64{1, 2}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10))
	require.Equal(t, reqId, uint64(3))
	require.NoError(t, err)
	coins = app.BankKeeper.GetAllBalances(ctx, addr1)
	require.Equal(t, coins, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 9970)})
	err = app.CustomGovKeeper.HandleIdentityRecordsVerifyRequest(ctx, addr3, 3, true)
	require.NoError(t, err)
	coins = app.BankKeeper.GetAllBalances(ctx, addr3)
	require.Equal(t, coins, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 10030)})
	record = app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.NotNil(t, record)
	require.Equal(t, *record, types.IdentityRecord{
		Id:        1,
		Address:   addr1,
		Key:       "key",
		Value:     "value",
		Date:      now,
		Verifiers: []sdk.AccAddress{addr3},
	})
	record = app.CustomGovKeeper.GetIdentityRecordById(ctx, 2)
	require.NotNil(t, record)
	require.Equal(t, *record, types.IdentityRecord{
		Id:        2,
		Address:   addr2,
		Key:       "key",
		Value:     "value",
		Date:      now,
		Verifiers: []sdk.AccAddress{addr3},
	})

	// request id record 2 to addr3 by addr2
	reqId, err = app.CustomGovKeeper.RequestIdentityRecordsVerify(ctx, addr2, addr3, []uint64{2}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 0))
	require.Equal(t, reqId, uint64(4))
	require.NoError(t, err)
	err = app.CustomGovKeeper.HandleIdentityRecordsVerifyRequest(ctx, addr3, 4, true)
	require.NoError(t, err)
	record = app.CustomGovKeeper.GetIdentityRecordById(ctx, 2)
	require.NotNil(t, record)
	require.Equal(t, *record, types.IdentityRecord{
		Id:        2,
		Address:   addr2,
		Key:       "key",
		Value:     "value",
		Date:      now,
		Verifiers: []sdk.AccAddress{addr3},
	})

	// request non-exist identity record
	ctxCache, _ = ctx.CacheContext()
	reqId, err = app.CustomGovKeeper.RequestIdentityRecordsVerify(ctxCache, addr2, addr3, []uint64{5}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10))
	require.Equal(t, reqId, uint64(5))
	require.Error(t, err)

	// approve with non-approver
	reqId, err = app.CustomGovKeeper.RequestIdentityRecordsVerify(ctx, addr1, addr3, []uint64{1}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10))
	require.Equal(t, reqId, uint64(5))
	require.NoError(t, err)
	ctxCache, _ = ctx.CacheContext()
	err = app.CustomGovKeeper.HandleIdentityRecordsVerifyRequest(ctxCache, addr2, 5, true)
	require.Error(t, err)

	// approve not existing request id
	ctxCache, _ = ctx.CacheContext()
	err = app.CustomGovKeeper.HandleIdentityRecordsVerifyRequest(ctxCache, addr2, 0xFFFFF, true)
	require.Error(t, err)

	// try edit and check if verification records all gone
	infos["key1"] = "value1"
	err = app.CustomGovKeeper.DeleteIdentityRecords(ctx, addr1, []string{"key", "key1"})
	require.NoError(t, err)
	record = app.CustomGovKeeper.GetIdentityRecordById(ctx, 1)
	require.NotNil(t, record)
	require.Equal(t, *record, types.IdentityRecord{
		Id:      1,
		Address: addr1,
		Key:     "key",
		Value:   "value",
		Date:    now,
	})

	// check get queries
	requests := app.CustomGovKeeper.GetIdRecordsVerifyRequestsByRequester(ctx, addr1)
	require.Len(t, requests, 1)
	requests = app.CustomGovKeeper.GetIdRecordsVerifyRequestsByApprover(ctx, addr1)
	require.Len(t, requests, 1)
	requests = app.CustomGovKeeper.GetAllIdRecordsVerifyRequests(ctx)
	require.Len(t, requests, 1)

	// remove all and query again
	app.CustomGovKeeper.DeleteIdRecordsVerifyRequest(ctx, 5)
	requests = app.CustomGovKeeper.GetAllIdRecordsVerifyRequests(ctx)
	require.Len(t, requests, 0)

	// try to cancel request and check coin moves correctly
	reqId, err = app.CustomGovKeeper.RequestIdentityRecordsVerify(ctx, addr2, addr3, []uint64{2}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10))
	require.Equal(t, reqId, uint64(6))
	require.NoError(t, err)
	coins = app.BankKeeper.GetAllBalances(ctx, addr2)
	require.Equal(t, coins, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 9990)})
	cacheCtx, _ := ctx.CacheContext()
	err = app.CustomGovKeeper.CancelIdentityRecordsVerifyRequest(cacheCtx, addr3, 6)
	require.Error(t, err)
	err = app.CustomGovKeeper.CancelIdentityRecordsVerifyRequest(ctx, addr2, 6)
	require.NoError(t, err)
	request = app.CustomGovKeeper.GetIdRecordsVerifyRequest(ctx, 6)
	require.Nil(t, request)
	coins = app.BankKeeper.GetAllBalances(ctx, addr2)
	require.Equal(t, coins, sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)})

	// try deleting request after request creation
	reqId, err = app.CustomGovKeeper.RequestIdentityRecordsVerify(ctx, addr2, addr3, []uint64{2}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10))
	require.Equal(t, reqId, uint64(7))
	app.CustomGovKeeper.DeleteIdRecordsVerifyRequest(ctx, 7)
	request = app.CustomGovKeeper.GetIdRecordsVerifyRequest(ctx, 7)
	require.Nil(t, request)
	requests = app.CustomGovKeeper.GetIdRecordsVerifyRequestsByRequester(ctx, addr2)
	require.Len(t, requests, 0)
	requests = app.CustomGovKeeper.GetIdRecordsVerifyRequestsByApprover(ctx, addr3)
	require.Len(t, requests, 0)

	// try deleting id record after request creation
	reqId, err = app.CustomGovKeeper.RequestIdentityRecordsVerify(ctx, addr2, addr3, []uint64{2}, sdk.NewInt64Coin(sdk.DefaultBondDenom, 10))
	require.Equal(t, reqId, uint64(8))
	app.CustomGovKeeper.DeleteIdentityRecords(ctx, addr1, []string{})
	cacheCtx, _ = ctx.CacheContext()
	err = app.CustomGovKeeper.HandleIdentityRecordsVerifyRequest(cacheCtx, addr3, 8, true)
	require.Error(t, err)
}
