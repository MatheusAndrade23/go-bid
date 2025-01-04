package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/matheusandrade23/go-bid/internal/jsonutils"
	"github.com/matheusandrade23/go-bid/internal/usecases/product"
)

func (api *Api) handleCreateProduct(w http.ResponseWriter, r *http.Request){
	data, problemns, err := jsonutils.DecodeValidJson[product.CreateProductReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problemns)
		return
	}
	userID, ok := api.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)
	if !ok {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected error, try again later",
		})
		return
	}
	id, err := api.ProductService.CreateProduct(
		r.Context(),
		userID,
		data.ProductName,
		data.Description,
		data.BasePrice,
		data.AuctionEnd,
	)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "failed to create product auction try again later",
		})
		return
	}
	jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"message":    "product created with success",
		"product_id": id,
	})
}