package domain

type WooCommerceProduct struct {
	Name         string                 `json:"name"`
	Type         string                 `json:"type"`
	RegularPrice string                 `json:"regular_price"`
	Description  string                 `json:"description"`
	ShortDesc    string                 `json:"short_description"`
	Images       []WooCommerceImage     `json:"images"`
	Categories   []WooCommerceCategory  `json:"categories"`
	Attributes   []WooCommerceAttribute `json:"attributes"`
	//Dimensions   WooCommerceDimensions  `json:"dimensions"`
	VideoURL string `json:"video_url,omitempty"`
}

// WooCommerceImage represents the WooCommerce image data structure
type WooCommerceImage struct {
	Src string `json:"src"`
}

// WooCommerceCategory represents the WooCommerce category data structure
type WooCommerceCategory struct {
	ID int `json:"id"`
}

// WooCommerceAttribute represents the WooCommerce attribute data structure
type WooCommerceAttribute struct {
	Name   string   `json:"name"`
	Option []string `json:"option"`
}

// WooCommerceDimensions represents the WooCommerce dimensions data structure
type WooCommerceDimensions struct {
	Length string `json:"length"`
	Width  string `json:"width"`
	Height string `json:"height"`
}
