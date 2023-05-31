package repository

import (
	"context"

	"github.com/AbdulrahmanDaud10/commerce-backend-golang/pkg/api"
)

type ProductStorer interface {
	Insert(context.Context, *api.Product) error
	GetByID(context.Context, string) (*api.Product, error)
	GetAll(context.Context) ([]*api.Product, error)
}
