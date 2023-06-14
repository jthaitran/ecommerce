// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    ShipXanhShopeeRaw, err := UnmarshalShipXanhShopeeRaw(bytes)
//    bytes, err = ShipXanhShopeeRaw.Marshal()

package domain

import "encoding/json"

func UnmarshalShipXanhShopeeRaw(data []byte) (ShipXanhShopeeRaw, error) {
	var r ShipXanhShopeeRaw
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ShipXanhShopeeRaw) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ShipXanhShopeeRaw struct {
	Results []Result `json:"results"`
}

type Result struct {
	FacetCounts   []FacetCount  `json:"facet_counts"`
	Found         int64         `json:"found"`
	Hits          []Hit         `json:"hits"`
	OutOf         int64         `json:"out_of"`
	Page          int64         `json:"page"`
	RequestParams RequestParams `json:"request_params"`
	SearchCutoff  bool          `json:"search_cutoff"`
	SearchTimeMS  int64         `json:"search_time_ms"`
}

type FacetCount struct {
	Counts    []Count `json:"counts"`
	FieldName string  `json:"field_name"`
	Stats     Stats   `json:"stats"`
}

type Count struct {
	Count       int64  `json:"count"`
	Highlighted string `json:"highlighted"`
	Value       string `json:"value"`
}

type Stats struct {
	TotalValues int64 `json:"total_values"`
}

type Hit struct {
	Document   Document      `json:"document"`
	Highlights []interface{} `json:"highlights"`
	TextMatch  int64         `json:"text_match"`
}

type Document struct {
	AccessibleToOrganizationID AccessibleToOrganizationID `json:"accessible_to_organization_id"`
	Attributes                 []interface{}              `json:"attributes"`
	Brand                      *string                    `json:"brand,omitempty"`
	BrandID                    int64                      `json:"brand_id"`
	Categories                 []interface{}              `json:"categories"`
	Condition                  int64                      `json:"condition"`
	CopyTo                     CopyTo                     `json:"copy_to"`
	Ctime                      int64                      `json:"ctime"`
	CurrentBrand               CurrentBrand               `json:"current_brand"`
	CurrentCate                CurrentCate                `json:"current_cate"`
	Description                string                     `json:"description"`
	Errors                     []interface{}              `json:"errors"`
	FlexImageIDSelected        *string                    `json:"flex_image_id_selected,omitempty"`
	FlexImages                 []FlexImage                `json:"flex_images,omitempty"`
	Highlights                 string                     `json:"highlights"`
	HistoryPost                []HistoryPost              `json:"history_post,omitempty"`
	ID                         string                     `json:"id"`
	Images                     []string                   `json:"images"`
	ImagesSize                 []interface{}              `json:"images_size"`
	IsOfficialShop             bool                       `json:"is_official_shop"`
	Itemid                     int64                      `json:"itemid"`
	Models                     []Model                    `json:"models"`
	Mtime                      int64                      `json:"mtime"`
	Name                       string                     `json:"name"`
	ObjectID                   string                     `json:"objectID"`
	OriginCate                 string                     `json:"origin_cate"`
	OriginImages               []string                   `json:"origin_images"`
	OriginShop                 OriginShop                 `json:"origin_shop"`
	OriginURL                  string                     `json:"origin_url"`
	Price                      int64                      `json:"price"`
	Ptime                      int64                      `json:"ptime"`
	ShipxanhCtime              int64                      `json:"shipxanh_ctime"`
	ShipxanhHeight             int64                      `json:"shipxanh_height"`
	ShipxanhIncludeVideo       bool                       `json:"shipxanh_include_video"`
	ShipxanhLength             int64                      `json:"shipxanh_length"`
	ShipxanhMtime              int64                      `json:"shipxanh_mtime"`
	ShipxanhSku                string                     `json:"shipxanh_sku"`
	ShipxanhUnit               string                     `json:"shipxanh_unit"`
	ShipxanhWeight             int64                      `json:"shipxanh_weight"`
	ShipxanhWidth              int64                      `json:"shipxanh_width"`
	Shopid                     int64                      `json:"shopid"`
	Source                     Source                     `json:"source"`
	Stock                      int64                      `json:"stock"`
	Targets                    []string                   `json:"targets"`
	TextMatch                  int64                      `json:"text_match"`
	TierVariations             []TierVariation            `json:"tier_variations"`
	ValidToCopy                bool                       `json:"valid_to_copy"`
	VideoInfoList              []VideoInfoList            `json:"video_info_list"`
	Warranty                   *Warranty                  `json:"warranty,omitempty"`
}

type CurrentBrand struct {
	ID   int64            `json:"id"`
	Name CurrentBrandName `json:"name"`
}

type FlexImage struct {
	ID        string   `json:"id"`
	Images    []string `json:"images"`
	IsDefault bool     `json:"is_default"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
	TitleTags []string `json:"title_tags"`
}

type HistoryPost struct {
	Auto      bool     `json:"auto"`
	Ctime     int64    `json:"ctime"`
	ShopID    ShopID   `json:"shop_id"`
	ShopName  ShopName `json:"shop_name"`
	SuccessID int64    `json:"success_id"`
}

type Model struct {
	Extinfo             Extinfo `json:"extinfo"`
	ID                  string  `json:"id"`
	InputPrice          int64   `json:"input_price"`
	Itemid              int64   `json:"itemid"`
	Modelid             string  `json:"modelid"`
	Name                string  `json:"name"`
	Price               int64   `json:"price"`
	PriceBeforeDiscount *int64  `json:"price_before_discount,omitempty"`
	Sku                 string  `json:"sku"`
	Stock               int64   `json:"stock"`
}

type Extinfo struct {
	EstimatedDays int64       `json:"estimated_days"`
	GroupBuyInfo  interface{} `json:"group_buy_info"`
	IsPreOrder    bool        `json:"is_pre_order"`
	TierIndex     []int64     `json:"tier_index"`
}

type TierVariation struct {
	Name         TierVariationName `json:"name"`
	Options      []string          `json:"options"`
	Properties   []interface{}     `json:"properties"`
	SummedStocks []int64           `json:"summed_stocks"`
	Type         int64             `json:"type"`
	Images       []string          `json:"images,omitempty"`
}

type VideoInfoList struct {
	VideoID       string      `json:"video_id"`
	DefaultFormat *Format     `json:"default_format,omitempty"`
	Duration      *int64      `json:"duration,omitempty"`
	Formats       []Format    `json:"formats,omitempty"`
	ItemCover     interface{} `json:"item_cover"`
	MmsData       *string     `json:"mms_data,omitempty"`
	ThumbURL      *string     `json:"thumb_url,omitempty"`
	Version       *int64      `json:"version,omitempty"`
	Vid           *string     `json:"vid,omitempty"`
}

type Format struct {
	Defn    Defn    `json:"defn"`
	Format  int64   `json:"format"`
	Height  int64   `json:"height"`
	Path    string  `json:"path"`
	Profile Profile `json:"profile"`
	URL     string  `json:"url"`
	Width   int64   `json:"width"`
}

type Warranty struct {
	WarrantyPeriod interface{} `json:"warranty_period"`
	WarrantyPolicy string      `json:"warranty_policy"`
}

type RequestParams struct {
	CollectionName string `json:"collection_name"`
	PerPage        int64  `json:"per_page"`
	Q              string `json:"q"`
}

type AccessibleToOrganizationID string

const (
	NEHGPh7H6OWYvYHs67OcUmQ3NZh1 AccessibleToOrganizationID = "NEHGPh7H6OWYvYHs67ocUmQ3NZh1"
)

type CopyTo string

const (
	Woo CopyTo = "woo"
)

type CurrentBrandName string

const (
	NoBrand CurrentBrandName = "No Brand"
)

type CurrentCate string

const (
	The15ChưaPhânLoại CurrentCate = "15 - Chưa phân loại"
)

type ShopID string

const (
	WooHomikidsVn ShopID = "woo_homikids.vn"
)

type ShopName string

const (
	HomikidsVn ShopName = "homikids.vn"
)

type OriginShop string

const (
	Homikids29 OriginShop = "homikids29"
)

type Source string

const (
	ShopeeVn Source = "shopee.vn"
)

type TierVariationName string

const (
	Empty       TierVariationName = ""
	KichCỡ      TierVariationName = "Kich cỡ"
	KíchCỡ      TierVariationName = "kích cỡ"
	KíchThước   TierVariationName = "Kích thước"
	LoạiSảnPhẩm TierVariationName = "Loại sản phẩm"
	MàuSắc      TierVariationName = "màu sắc"
	Mẫu         TierVariationName = "Mẫu"
	NameKíchCỡ  TierVariationName = "Kích cỡ"
	NameMàuSắc  TierVariationName = "Màu sắc"
	PhânLoại    TierVariationName = "phân loại"
)

type Defn string

const (
	Ori   Defn = "ORI"
	V480P Defn = "V480P"
	V720P Defn = "V720P"
)

type Profile string

const (
	Mp4 Profile = "MP4"
)
