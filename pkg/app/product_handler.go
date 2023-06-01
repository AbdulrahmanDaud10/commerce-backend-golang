package app

import (
	"encoding/json"
	"net/http"

	"github.com/AbdulrahmanDaud10/commerce-backend-golang/pkg/api"
	"github.com/AbdulrahmanDaud10/commerce-backend-golang/pkg/repository"
	"github.com/anthdm/weavebox"
)

type ProductHandler struct {
	Store repository.ProductStorer
}

func NewProductHandler(productStore repository.ProductStorer) *ProductHandler {
	return &ProductHandler{
		Store: productStore,
	}
}

func (h *ProductHandler) HandlePostProduct(ctx *weavebox.Context) error {
	productRequest := &api.CreateProductRequest{}
	if err := json.NewDecoder(ctx.Request().Body).Decode(productRequest); err != nil {
		return err
	}

	product, err := api.NewProductFromRequest(productRequest)
	if err != nil {
		return err
	}

	if err := h.Store.Insert(ctx.Context, product); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) HandleGetProducts(ctx *weavebox.Context) error {
	products, err := h.Store.GetAll(ctx.Context)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) HandleGetProductsByID(ctx *weavebox.Context) error {
	id := ctx.Param("id")

	product, err := h.Store.GetByID(ctx.Context, id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, product)
}
