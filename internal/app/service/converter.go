package service

import (
	"ecommerce/internal/app/domain"
	"fmt"
)

// Function to convert ShipXanhShopeeRaw to Product
func ConvertShipXanhShopeeRawToProduct(shipXanhShopeeRaw domain.ShipXanhShopeeRaw) (*[]domain.WooCommerceProduct, error) {
	rawProduct := shipXanhShopeeRaw.Results[0].Hits
	print(rawProduct)
	var products []domain.WooCommerceProduct
	for _, item := range rawProduct {
		product := domain.WooCommerceProduct{
			Name:         item.Document.Name,
			Type:         "simple",
			RegularPrice: fmt.Sprintf("%.2f", float64(item.Document.Price)/100000),
			Description:  item.Document.Description,
			//ShortDescription: item.Document.Highlights,
		}

		// Map categories
		//for _, category := range item.Document.Categories {
		//	product.Categories = append(product.Categories, Category{
		//		ID: category.(int64),
		//	})
		//}

		// Map images
		for _, image := range item.Document.Images {
			fileName, errrr := getFileNameFromURL(image)
			if errrr != nil {
				fmt.Printf("Error getting file name: %v\n", errrr)
				return products, errrr
			}
			// Remove the .png extension from the file path
			fileName = removeExtension(fileName, ".png")
			fileName = fmt.Sprintf("https://homikids.vn/wp-content/uploads/product_images/%s.jpeg", fileName)

			product.Images = append(product.Images, Image{

				Src: fileName,
			})
		}
		products = append(products, product)
	}

	return products, nil
}

//func createProductFromRaw() {
//	// Assume you have the ShipXanhShopeeRaw data stored in the variable "shipXanhShopeeRaw"
//
//	// Convert to Product
//	product, err := ConvertShipXanhShopeeRawToProduct(shipXanhShopeeRaw)
//	if err != nil {
//		fmt.Println("Error converting ShipXanhShopeeRaw to Product:", err)
//		return
//	}
//
//	// Marshal the Product to JSON
//	data, err := product.Marshal()
//	if err != nil {
//		fmt.Println("Error marshaling Product to JSON:", err)
//		return
//	}
//
//	// Print the JSON data
//	fmt.Println(string(data))
//}
