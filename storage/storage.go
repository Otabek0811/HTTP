package storage

import "app/models"

type StorageI interface {
	Close()
	Category() CategoryRepoI
	Product() ProductRepoI
}

type CategoryRepoI interface {
	CreateCategory(*models.CreateCategory) (string, error)
	GetListCategory(*models.CategoryGetListRequest)(*models.CategoryGetListResponse, error)
	GetCategoryByID(*models.CategoryPrimaryKey) (*models.Category, error)
	UpdateCategory(*models.UpdateCategory)(string, error)
	DeleteCategory(*models.CategoryPrimaryKey)(error)
}

type ProductRepoI interface{
	CreateProduct(*models.CreateProduct) (string, error)
	GetListProduct(*models.ProductGetListRequest)(*models.ProductGetListResponse, error)
	GetProductByID(req *models.ProductPrimaryKey) (*models.Product, error)
	UpdateProduct(*models.UpdateProduct)(string, error)
	DeleteProduct(*models.ProductPrimaryKey)(error)
}
