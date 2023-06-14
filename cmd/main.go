// main.go
package main

import (
	"ecommerce/internal/app/api"
	"ecommerce/internal/app/config"
	"ecommerce/internal/app/domain"
	"ecommerce/internal/app/service"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	//middleware
	r.Use(logger.SetLogger())

	err := config.LoadConfig("/Users/thaitran/WorkingSpace/GitHub/ecomerce/config.yaml")
	if err != nil {
		fmt.Println("Load config error.....")
	}

	cfg := config.GetConfig()

	//Routers
	r.GET("/products", func(c *gin.Context) {
		baseURL := cfg.Woocommerce.WebsiteUrl + "/wp-json/wc/v3/products"
		client := api.NewAPIClient(baseURL, nil)

		// Make an API call
		products, err := client.GetProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"products": products,
		})
	})

	r.POST("/products", func(c *gin.Context) {
		woocommerceProducts := prepareProducts()
		baseURL := cfg.Woocommerce.WebsiteUrl + "/wp-json/wc/v3/products"
		client := api.NewAPIClient(baseURL, nil)

		// Make an API call
		products, err := client.NewProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"products": products,
		})
	})

	// Start server
	r.Run(":8080")

	//// Load configuration
	//configFilePath := viper.GetString("config.yaml")
	//err := config.LoadConfig(configFilePath)
	//if err != nil {
	//	fmt.Println("Error loading config:", err)
	//	return
	//}
	//
	//repo := &repository.InitWooProductRepository{
	//	Products: make(map[int]domain.WooCommerceProduct),
	//}
	//
	//productService := service.NewProductService(*repo)
	//products, err := productService.GetProducts()
	//if err != nil {
	//	_ = fmt.Errorf(err.Error())
	//}
	//
	//fmt.Println(products)
}

func prepareProducts() (*[]domain.WooCommerceProduct, error) {
	var err error
	var products []domain.WooCommerceProduct
	//read raw products from local
	//transfer the products - woocommerce
	shipXanhShopeeRaw, err := mappingRawProduct()
	if err != nil {
		return &products, err
	}

	// Convert to Product
	products, err = service.ConvertShipXanhShopeeRawToProduct(*shipXanhShopeeRaw)
	if err != nil {
		fmt.Println("Error converting ShipXanhShopeeRaw to Product:", err)
		return err
	}
	return &products, err
}

func mappingRawProduct() (*domain.ShipXanhShopeeRaw, error) {
	//read data from example
	file, err := os.Open("example.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return &domain.ShipXanhShopeeRaw{}, err
	}

	defer file.Close()

	//Create a decoder
	decoder := json.NewDecoder(file)
	//Map json data to struct
	var rawProduct domain.ShipXanhShopeeRaw
	err = decoder.Decode(&rawProduct)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return &domain.ShipXanhShopeeRaw{}, err
	}

	return &rawProduct, err
}
