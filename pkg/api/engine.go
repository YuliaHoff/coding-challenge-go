package api

import (
	"coding-challenge-go/pkg/api/product"
	"coding-challenge-go/pkg/api/seller"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// CreateAPIEngine creates engine instance that serves API endpoints,
// consider it as a router for incoming requests.
func CreateAPIEngine(db *sql.DB) (*gin.Engine, error) {
	r := gin.New()
	v1 := r.Group("api/v1")
	v2 := r.Group("api/v2")
	productRepository := product.NewRepository(db)
	sellerRepository := seller.NewRepository(db)
	emailProvider := seller.NewEmailProvider()
	productController := product.NewController(productRepository, sellerRepository, emailProvider)
	v1.GET("products", productController.List)
	v2.GET("products", productController.List)
	v1.GET("product", productController.Get)
	v2.GET("product", productController.Get)
	v1.POST("product", productController.Post)
	v1.PUT("product", productController.Put)
	v1.DELETE("product", productController.Delete)
	sellerController := seller.NewController(sellerRepository)
	v1.GET("sellers", sellerController.List)

	return r, nil
}
