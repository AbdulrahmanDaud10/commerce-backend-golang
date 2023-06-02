package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AbdulrahmanDaud10/commerce-backend-golang/pkg/app"
	"github.com/AbdulrahmanDaud10/commerce-backend-golang/pkg/repository"
	"github.com/anthdm/weavebox"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handleAPIError(ctx *weavebox.Context, err error) {
	fmt.Println("API error:", err)
	ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
}

func main() {
	application := weavebox.New()
	application.ErrorHandler = handleAPIError

	adminMW := &app.AdminAuthMiddleware{}
	adminRoute := application.Box("/admin")
	adminRoute.Use(adminMW.Authenticate)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	productStore := repository.NewMongoProductStore(client.Database("golangcommerce"))
	productHandler := app.NewProductHandler(productStore)

	// admin/product
	adminProductRoute := adminRoute.Box("/product")
	adminProductRoute.Get("/:id", productHandler.HandleGetProductByID)
	adminProductRoute.Get("/", productHandler.HandleGetProducts)
	adminProductRoute.Post("/", productHandler.HandlePostProduct)

	application.Serve(3001)
}
