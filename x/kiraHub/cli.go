package kiraHub


import (
	constants "github.com/KiraCore/sekai/x/kiraHub/constants"
	"github.com/KiraCore/sekai/x/kiraHub/transactions/createOrderBook"
	"github.com/KiraCore/sekai/x/kiraHub/queries/listOrderBooks"
	"github.com/spf13/cobra"

	"github.com/KiraCore/cosmos-sdk/client"
	"github.com/KiraCore/cosmos-sdk/client/flags"
	"github.com/KiraCore/cosmos-sdk/codec"
)

func GetCLIRootTransactionCommand(codec *codec.Codec) *cobra.Command {
	rootTransactionCommand := &cobra.Command{
		Use:                        "testtoken",
		Short:                      "Asset root transaction command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	rootTransactionCommand.AddCommand(flags.PostCommands(
		createOrderBook.TransactionCommand(codec),
	)...)
	return rootTransactionCommand
}

func GetCLIRootQueryCommand(codec *codec.Codec) *cobra.Command {
	rootQueryCommand := &cobra.Command{
		Use:                        constants.QuerierRoute,
		Short:                      "Asset root query command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	rootQueryCommand.AddCommand(flags.GetCommands(
		listOrderBooks.GetOrderBooksCmd(codec),
		listOrderBooks.GetOrderBooksByTPCmd(codec),
	)...)
	return rootQueryCommand
}