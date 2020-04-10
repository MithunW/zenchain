package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/okex/okchain/x/dex/types"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/okex/okchain/x/common"
	govRest "github.com/okex/okchain/x/gov/client/rest"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/products", productsHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/deposits", depositsHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/match_order", matchOrderHandler(cliCtx)).Methods("GET")
}

func productsHandler(cliContext context.CLIContext) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ownerAddress := r.URL.Query().Get("address")
		pageStr := r.URL.Query().Get("page")
		perPageStr := r.URL.Query().Get("per_page")

		var params = &types.QueryDexInfoParams{}
		err := params.SetPageAndPerPage(ownerAddress, pageStr, perPageStr)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusBadRequest, common.ErrorInvalidParam)
			return
		}
		bz, err := cliContext.Codec.MarshalJSON(&params)

		if err != nil {
			common.HandleErrorMsg(w, cliContext, err.Error())
			return
		}

		res, _, err := cliContext.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryProducts), bz)
		if err != nil {
			common.HandleErrorMsg(w, cliContext, err.Error())
			return
		}

		result := common.GetBaseResponse("hello")
		result2, err2 := json.Marshal(result)
		if err2 != nil {
			common.HandleErrorMsg(w, cliContext, err2.Error())
			return
		}
		result2 = []byte(strings.Replace(string(result2), "\"hello\"", string(res), 1))
		rest.PostProcessResponse(w, cliContext, result2)
	}

}

func depositsHandler(cliContext context.CLIContext) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		pageStr := r.URL.Query().Get("page")
		perPageStr := r.URL.Query().Get("per_page")

		if len(address) == 0 || address == "" {
			common.HandleErrorResponseV2(w, http.StatusBadRequest, common.ErrorMissingRequiredParam)
			return
		}
		var params = &types.QueryDexInfoParams{}
		err := params.SetPageAndPerPage(address, pageStr, perPageStr)

		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusBadRequest, common.ErrorInvalidParam)
			return
		}
		bz, err := cliContext.Codec.MarshalJSON(&params)
		if err != nil {
			common.HandleErrorMsg(w, cliContext, err.Error())
			return
		}

		res, _, err := cliContext.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDeposits), bz)
		if err != nil {
			common.HandleErrorMsg(w, cliContext, err.Error())
			return
		}

		result := common.GetBaseResponse("hello")
		result2, err2 := json.Marshal(result)
		if err2 != nil {
			common.HandleErrorMsg(w, cliContext, err2.Error())
			return
		}
		result2 = []byte(strings.Replace(string(result2), "\"hello\"", string(res), 1))
		rest.PostProcessResponse(w, cliContext, result2)
	}

}

func matchOrderHandler(cliContext context.CLIContext) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		perPageStr := r.URL.Query().Get("per_page")

		var params = &types.QueryDexInfoParams{}
		err := params.SetPageAndPerPage("", pageStr, perPageStr)
		if err != nil {
			common.HandleErrorResponseV2(w, http.StatusBadRequest, common.ErrorInvalidParam)
			return
		}
		bz, err := cliContext.Codec.MarshalJSON(&params)

		if err != nil {
			common.HandleErrorMsg(w, cliContext, err.Error())
			return
		}

		res, _, err := cliContext.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryMatchOrder), bz)
		if err != nil {
			common.HandleErrorMsg(w, cliContext, err.Error())
			return
		}

		result := common.GetBaseResponse("hello")
		result2, err2 := json.Marshal(result)
		if err2 != nil {
			common.HandleErrorMsg(w, cliContext, err2.Error())
			return
		}
		result2 = []byte(strings.Replace(string(result2), "\"hello\"", string(res), 1))
		rest.PostProcessResponse(w, cliContext, result2)
	}

}

// DelistProposalRESTHandler defines dex proposal handler
func DelistProposalRESTHandler(context.CLIContext) govRest.ProposalRESTHandler {
	return govRest.ProposalRESTHandler{}
}
