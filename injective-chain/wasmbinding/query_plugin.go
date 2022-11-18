package wasmbinding

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	OracleRoute       = "oracle"
	ExchangeRoute     = "exchange"
	TokenFactoryRoute = "tokenfactory"
)

type InjectiveQueryWrapper struct {
	// specifies which module handler should handle the query
	Route string `json:"route,omitempty"`
	// The query data that should be parsed into the module query
	QueryData json.RawMessage `json:"query_data,omitempty"`
}

// CustomQuerier dispatches custom CosmWasm bindings queries.
func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery InjectiveQueryWrapper
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrap(err, "Error parsing request data")
		}

		switch contractQuery.Route {
		case OracleRoute:
			return qp.HandleOracleQuery(ctx, contractQuery.QueryData)
		case ExchangeRoute:
			return qp.HandleExchangeQuery(ctx, contractQuery.QueryData)
		case TokenFactoryRoute:
			return qp.HandleTokenFactoryQuery(ctx, contractQuery.QueryData)
		default:
			return nil, wasmvmtypes.UnsupportedRequest{Kind: "Unknown Injective Query Route"}
		}
	}
}
