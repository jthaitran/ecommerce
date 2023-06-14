package api

import (
	"ecommerce/internal/app/config"
	"ecommerce/internal/app/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

type APIClient struct {
	BaseURL             string
	WooCommerceProducts []domain.WooCommerceProduct
}

type ResponseData struct {
	Message string `json:"message"`
}

func NewAPIClient(baseURL string, wooCommerceProducts []domain.WooCommerceProduct) *APIClient {
	return &APIClient{
		BaseURL:             baseURL,
		WooCommerceProducts: wooCommerceProducts,
	}
}

func (c *APIClient) GetExampleData() (*ResponseData, error) {
	resp, err := http.Get(c.BaseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data ResponseData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *APIClient) GetProducts() (*[]domain.WooCommerceProduct, error) {
	var products []domain.WooCommerceProduct
	req, err := http.NewRequest("GET", c.BaseURL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.GetConfig().Woocommerce.ConsumerKey, config.GetConfig().Woocommerce.ConsumerSecret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve products: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (c *APIClient) NewProducts() (*[]domain.WooCommerceProduct, error) {
	var products []domain.WooCommerceProduct
	req, err := http.NewRequest("POST", c.BaseURL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.GetConfig().Woocommerce.ConsumerKey, config.GetConfig().Woocommerce.ConsumerSecret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create products: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return nil, err
	}

	return &products, nil
}
