package main

//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//// ShopeeProduct represents the Shopee product data structure
//type ShopeeProduct struct {
//	FacetCounts []struct {
//		Counts []struct {
//			Count       int    `json:"count"`
//			Highlighted string `json:"highlighted"`
//			Value       string `json:"value"`
//		} `json:"counts"`
//		FieldName string `json:"field_name"`
//		Stats     struct {
//			TotalValues int `json:"total_values"`
//		} `json:"stats"`
//	} `json:"facet_counts"`
//	Found int `json:"found"`
//	Hits  []struct {
//		Document struct {
//			AccessibleToOrganizationID string   `json:"accessible_to_organization_id"`
//			Attributes                 []string `json:"attributes"`
//			Brand                      string   `json:"brand"`
//			BrandID                    int      `json:"brand_id"`
//			Categories                 []string `json:"categories"`
//			Condition                  int      `json:"condition"`
//			CopyTo                     string   `json:"copy_to"`
//			CTime                      int      `json:"ctime"`
//			CurrentBrand               struct {
//				ID   int    `json:"id"`
//				Name string `json:"name"`
//			} `json:"current_brand"`
//			CurrentCate         string   `json:"current_cate"`
//			Description         string   `json:"description"`
//			Errors              []string `json:"errors"`
//			FlexImageIDSelected string   `json:"flex_image_id_selected"`
//			FlexImages          []struct {
//				ID     string   `json:"id"`
//				Images []string `json:"images"`
//			} `json:"flex_images"`
//		} `json:"document"`
//	} `json:"hits"`
//}
//
//// WooCommerceProduct represents the WooCommerce product data structure
//type WooCommerceProduct struct {
//	Name         string                 `json:"name"`
//	Type         string                 `json:"type"`
//	RegularPrice string                 `json:"regular_price"`
//	Description  string                 `json:"description"`
//	ShortDesc    string                 `json:"short_description"`
//	Images       []WooCommerceImage     `json:"images"`
//	Categories   []WooCommerceCategory  `json:"categories"`
//	Attributes   []WooCommerceAttribute `json:"attributes"`
//	//Dimensions   WooCommerceDimensions  `json:"dimensions"`
//	VideoURL string `json:"video_url,omitempty"`
//}
//
//// WooCommerceImage represents the WooCommerce image data structure
//type WooCommerceImage struct {
//	Src string `json:"src"`
//}
//
//// WooCommerceCategory represents the WooCommerce category data structure
//type WooCommerceCategory struct {
//	ID int `json:"id"`
//}
//
//// WooCommerceAttribute represents the WooCommerce attribute data structure
//type WooCommerceAttribute struct {
//	Name   string   `json:"name"`
//	Option []string `json:"option"`
//}
//
//// WooCommerceDimensions represents the WooCommerce dimensions data structure
////type WooCommerceDimensions struct {
////	Length string `json:"length"`
////	Width  string `json:"width"`
////	Height string `json:"height"`
////}
//
//func main() {
//	// Example JSON string representing the Shopee product data
//	inputData := `{"results":[{"facet_counts":[{"counts":[{"count":189,"highlighted":"true","value":"true"}],"field_name":"valid_to_copy","stats":{"total_values":1}},{"counts":[{"count":189,"highlighted":"0","value":"0"}],"field_name":"targets","stats":{"total_values":1}},{"counts":[{"count":189,"highlighted":"woo","value":"woo"}],"field_name":"copy_to","stats":{"total_values":1}}],"found":189,"hits":[{"document":{"accessible_to_organization_id":"NEHGPh7H6OWYvYHs67ocUmQ3NZh1","attributes":[],"brand":"Natureby","brand_id":2435644,"categories":[],"condition":1,"copy_to":"woo","ctime":1685109646,"current_brand":{"id":0,"name":"No Brand"},"current_cate":"15 - Ch\u01b0a ph\u00e2n lo\u1ea1i","description":"<p><p>M\u1eb7t N\u1ea1 Natureby H\u00e0n Qu\u1ed1c Ch\u00ednh H\u00e3ng</p><p>Qui C\u00e1ch: 1 Mi\u1ebfng </p><p>Xu\u1ea5t x\u1ee9: 100% H\u00e0n Qu\u1ed1c</p><p>HSD: 2024 (t\u00f9y t\u1eebng lo\u1ea1i m\u00e0 NG\u00c0Y TH\u00c1NG s\u1ea3n xu\u1ea5t kh\u00e1c nhau \u1ea1)</p><p>M\u1eb7t n\u1ea1 d\u01b0\u1ee1ng da Natureby ch\u00ednh h\u00e3ng H\u00e0n Qu\u1ed1c</p><p>C\u00f3 ph\u1ea3i b\u1ea1n r\u1ea5t b\u1eadn r\u1ed9n v\u00e0 kh\u00f4ng c\u00f3 th\u1eddi gian d\u00e0nh cho b\u1ea3n th\u00e2n ?</p><p>C\u00f3 ph\u1ea3i b\u1ea1n r\u1ea5t mu\u1ed1n l\u00e0m \u0111\u1eb9p m\u00e0 s\u1ee3 t\u1ed1n th\u1eddi gian ?</p><p>C\u00f3 ph\u1ea3i b\u1ea1n r\u1ea5t mu\u1ed1n s\u1ee1 h\u1eefu m\u1ed9t g\u01b0\u01a1ng m\u1eb7t tr\u1eafng s\u00e1ng m\u1ecbn m\u00e0ng v\u00e0 kh\u00f4ng b\u1ecb n\u1ed5i m\u1ee5n</p><p>V\u00e0 \u0111\u00e2y l\u00e0 gi\u1ea3i ph\u00e1p d\u00e0nh cho b\u1ea1n ====> [S\u1ec9 Gi\u00e1 T\u1ed1t] 14 Lo\u1ea1i M\u1eb7t N\u1ea1 Natureby H\u00e0n Qu\u1ed1c</p><p>-Natureby l\u00e0 m\u1ed9t trong nh\u1eefng m\u1ef9 ph\u1ea9m H\u00e0n Qu\u1ed1c \u0111\u01b0\u1ee3c ph\u00e1i n\u1eef tin d\u00f9ng. Natureby \u0111\u01b0\u1ee3c chi\u1ebft  100% t\u1eeb thi\u00ean nhi\u00ean \u0111\u1eb7t bi\u1ec7t an to\u00e0n cho m\u1ecdi lo\u1ea1i da </p><p>-H\u1ed2NG S\u00c2M: Chi\u1ebft xu\u1ea5t h\u1ed3ng s\u00e2m H\u00e0n Qu\u1ed1c gi\u00fap ph\u1ee5c da b\u1ecb h\u01b0 t\u1ed5n, nu\u00f4i d\u01b0\u1ee1ng da, d\u01b0\u1ee1ng \u1ea9m gi\u00fap da m\u1ec1m m\u1ea1i c\u0103ng b\u00f3ng v\u00e0 ng\u00e0y c\u00e0ng h\u1ed3ng h\u00e0o h\u01a1n.</p><p>-RONG BI\u1ec2N: chi\u1ebft xu\u1ea5t t\u1eeb rong bi\u1ec3n gi\u00fap da c\u0103ng m\u01b0\u1ee3t, r\u1ea1ng r\u1ee1 v\u00e0 \u0111\u1ea7y s\u1ee9c s\u1ed1ng. S\u1ea3n ph\u1ea9m v\u1edbi t\u00e1c d\u1ee5ng c\u1ea5p n\u01b0\u1edbc th\u1ea7n k\u00ec, kh\u00f4ng ch\u1ee9a c\u00e1c ch\u1ea5t c\u00f3 h\u1ea1i cho da, th\u00e0nh ph\u1ea7n chi\u1ebft xu\u1ea5t thi\u00ean nhi\u00ean l\u00e0nh t\u00ednh, mang l\u1ea1i \u0111\u1ed9 \u1ea9m l\u00fd t\u01b0\u1edfng, l\u00e0n da c\u0103ng m\u01b0\u1ee3t v\u00e0 \u0111\u01b0\u1ee3c c\u1ea3i thi\u1ec7n r\u00f5 r\u1ec7t.</p><p>-NHA \u0110AM : M\u1eb7t n\u1ea1 chi\u1ebft xu\u1ea5t nha \u0111am l\u00e0m d\u1ecbu da, c\u00f3 kh\u1ea3 n\u0103ng l\u00e0m x\u1eb9p m\u1ee5n, gi\u1ea3m c\u00e1c v\u1ebft s\u01b0ng t\u1ea5y do m\u1ee5n g\u00e2y ra, ch\u1ea5t gel c\u1ee7a nha \u0111am c\u00f3 nhi\u1ec7m v\u1ee5 t\u1ea9y c\u00e1c t\u1ebf b\u00e0o ch\u1ebft v\u00e0 gi\u1eef cho da lu\u00f4n t\u01b0\u01a1i tr\u1ebb. Gi\u1eef \u0111\u01b0\u1ee3c \u0111\u1ed9 \u1ea9m c\u1ea9n thi\u1ebft tr\u00e1nh hi\u1ec7n t\u01b0\u1ee3ng n\u1ee9t n\u1ebb, x\u00f3a v\u1ebft nh\u0103n v\u00e0 th\u00e2m quanh v\u00f9ng m\u00ed m\u1eaft sau nh\u1eefng ng\u00e0y l\u00e0m vi\u1ec7c m\u1ec7t m\u1ecfi.</p><p>-D\u01afA LEO: d\u01b0\u1ee1ng tr\u1eafng s\u00e1ng da, gi\u1ea3m s\u01b0ng t\u1ea5y v\u00f9ng da d\u01b0\u1edbi m\u1eaft.</p><p>-L\u1ef1u: d\u01b0\u1ee1ng tr\u1eafng s\u00e1ng da v\u00e0 ch\u1ed1ng l\u00e3o h\u00f3a</p><p>-Hoa h\u1ed3ng: d\u01b0\u1ee1ng s\u00e1ng tr\u1eafng da, x\u00f3a m\u1edd v\u1ebft th\u00e2m, se kh\u00edt l\u1ed7 ch\u00e2n l\u00f4ng</p><p>-G\u1ea1o :  b\u1ed9t g\u1ea1o c\u0169ng gi\u00fap l\u00e0m m\u1edd n\u00e1m, x\u00f3a t\u00e0n nhang h\u1eefu hi\u1ec7u ,T\u00ecnh tr\u1ea1ng c\u00e1c l\u1ed7 ch\u00e2n l\u00f4ng to, x\u1ec9n m\u00e0u d\u1ec5 d\u00e0ng \u0111\u01b0\u1ee3c kh\u1eafc ph\u1ee5c</p><p>-\u1ed0c s\u00ean : l\u00e0m m\u1edd v\u1ebft nh\u0103n hi\u1ec7u qu\u1ea3, \u0111\u1ed3ng th\u1eddi nu\u00f4i d\u01b0\u1ee1ng v\u00e0 t\u0103ng \u0111\u1ed9 \u0111\u00e0n h\u1ed3i cho da, ph\u1ee5c h\u1ed3i s\u1ef1 t\u01b0\u01a1i tr\u1ebb cho l\u00e0n da l\u00e3o h\u00f3a.</p><p>-H\u01b0\u1edbng d\u1eabn s\u1eed d\u1ee5ng:</p><p>+L\u00e0m s\u1ea1ch da m\u1eb7t b\u1eb1ng n\u01b0\u1edbc \u1ea5m ho\u1eb7c s\u1eefa r\u1eeda m\u1eb7t</p><p>+B\u1eaft \u0111\u1ea7u \u0111\u1eafp m\u1eb7t n\u1ea1 t\u1eeb v\u00f9ng tr\u00e1n sao cho ph\u1ee7 \u0111\u1ec1u h\u1ebft g\u01b0\u01a1ng m\u1eb7t</p><p>+Th\u01b0 gi\u00e3n kho\u1ea3ng 20 - 30 ph\u00fat</p><p>+L\u1ed9t b\u1ecf m\u1eb7t n\u1ea1, nh\u1eb9 nh\u00e0ng thoa tinh ch\u1ea5t c\u00f2n s\u00f3t l\u1ea1i tr\u00ean m\u1eb7t n\u1ea1 \u0111\u1ec3 cho d\u01b0\u1ee1ng ch\u1ea5t \u0111\u01b0\u1ee3c h\u1ea5p th\u1ee5 s\u00e2u v\u00e0o b\u00ean trong da r\u1ed3i r\u1eeda l\u1ea1i b\u1eb1ng n\u01b0\u1edbc s\u1ea1ch.</p><p>+1 tu\u1ea7n n\u00ean \u0111\u1eafp m\u1eb7t n\u1ea1 2 - 3 l\u1ea7n \u0111\u1ec3 c\u00f3 1 l\u00e0n da \u0111\u1eb9p, duy tr\u00ec tu\u1ed5i thanh xu\u00e2n nh\u00e1 \ud83d\ude1a\ud83d\ude1a</p><p>-H\u00c3Y T\u01af\u1edeNG T\u01af\u1ee2NG m\u1ecdi ng\u01b0\u1eddi tr\u1ea7m tr\u1ed3 v\u00ec b\u1ea1n c\u00f3 m\u1ed9t khu\u00f4n m\u1eb7t m\u1ecbn m\u00e0ng tr\u1eafng s\u00e1ng v\u00e0 n\u00f3i kh\u00f4ng v\u1edbi m\u1ee5n </p><p>#matna #matnaduongda #matnatrangda #natureby #naturebyhanquoc #naturebychinhhang #naturebynhadam</p></p>","errors":[],"flex_image_id_selected":"1685110188","flex_images":[{"id":"1685110188","images":["https://cf.shopee.vn/file/d2f3d7921f17f77e1c6c9b4719aaf950","https://cf.shopee.vn/file/9e8f0e9a46789b0abc7da109f8131ace","https://cf.shopee.vn/file/da12dcd4ecb76b413912fdd2229d89be","https://cf.shopee.vn/file/6145402ce6159dc93117a9e4f3b85026","https://cf.shopee.vn/file/39805027cd5024f38866eb66fe1fbb1a","https://cf.shopee.vn/file/812858846d2e52361c02530d1123e608","https://cf.shopee.vn/file/bf512017f8f7eb0b5c403199b9efbc19"],"is_default":true,"name":"ver 1","tags":["sx_default_desc"],"title_tags":["sx_default_title"]}],"highlights":"<HOMI BEAUTY> M\u1eb7t N\u1ea1 Natureby H\u00e0n Qu\u1ed1c Ch\u00ednh H\u00e3ng HMB018","history_post":[{"auto":false,"ctime":1685110866,"shop_id":"woo_homikids.vn","shop_name":"homikids.vn","success_id":1607},{"auto":false,"ctime":1685110649,"shop_id":"woo_homikids.vn","shop_name":"homikids.vn","success_id":1414}],"id":"shopee.vn_woo_402282834_14159000497_li1eqmxq","images":["https://cf.shopee.vn/file/d2f3d7921f17f77e1c6c9b4719aaf950","https://cf.shopee.vn/file/9e8f0e9a46789b0abc7da109f8131ace","https://cf.shopee.vn/file/da12dcd4ecb76b413912fdd2229d89be","https://cf.shopee.vn/file/6145402ce6159dc93117a9e4f3b85026","https://cf.shopee.vn/file/39805027cd5024f38866eb66fe1fbb1a","https://cf.shopee.vn/file/812858846d2e52361c02530d1123e608","https://cf.shopee.vn/file/bf512017f8f7eb0b5c403199b9efbc19"],"images_size":[],"is_official_shop":false,"itemid":14159000497,"models":[{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[0]},"id":"Nh\u00e2n s\u00e2m","input_price":0,"itemid":14159000497,"modelid":"160584535243","name":"Nh\u00e2n s\u00e2m","price":20000,"price_before_discount":20000,"sku":"14159000497_10","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[1]},"id":"D\u01b0a leo","input_price":0,"itemid":14159000497,"modelid":"152817866085","name":"D\u01b0a leo","price":20000,"price_before_discount":20000,"sku":"14159000497_11","stock":9},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[2]},"id":"C\u00e0 chua","input_price":0,"itemid":14159000497,"modelid":"152817866086","name":"C\u00e0 chua","price":20000,"price_before_discount":20000,"sku":"14159000497_6","stock":9},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[3]},"id":"\u1ed0c s\u00ean","input_price":0,"itemid":14159000497,"modelid":"152817866080","name":"\u1ed0c s\u00ean","price":20000,"price_before_discount":20000,"sku":"14159000497_1","stock":6},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[4]},"id":"Khoai t\u00e2y","input_price":0,"itemid":14159000497,"modelid":"160584535244","name":"Khoai t\u00e2y","price":20000,"price_before_discount":20000,"sku":"14159000497_8","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[5]},"id":"Rong bi\u1ec3n","input_price":0,"itemid":14159000497,"modelid":"152817866087","name":"Rong bi\u1ec3n","price":20000,"price_before_discount":20000,"sku":"14159000497_12","stock":10},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[6]},"id":"B\u01a1","input_price":0,"itemid":14159000497,"modelid":"160584535240","name":"B\u01a1","price":20000,"price_before_discount":20000,"sku":"14159000497_7","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[7]},"id":"Vi\u1ec7t qu\u1ea5t","input_price":0,"itemid":14159000497,"modelid":"160584535241","name":"Vi\u1ec7t qu\u1ea5t","price":20000,"price_before_discount":20000,"sku":"14159000497_13","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[8]},"id":"M\u1eadt ong","input_price":0,"itemid":14159000497,"modelid":"160584535242","name":"M\u1eadt ong","price":20000,"price_before_discount":20000,"sku":"14159000497_9","stock":10},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[9]},"id":"Colagen","input_price":0,"itemid":14159000497,"modelid":"152817866081","name":"Colagen","price":20000,"price_before_discount":20000,"sku":"14159000497_4","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[10]},"id":"Tr\u00e0 xanh","input_price":0,"itemid":14159000497,"modelid":"152817866082","name":"Tr\u00e0 xanh","price":20000,"price_before_discount":20000,"sku":"14159000497_2","stock":10},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[11]},"id":"Nha \u0111am","input_price":0,"itemid":14159000497,"modelid":"152817866083","name":"Nha \u0111am","price":20000,"price_before_discount":20000,"sku":"14159000497_3","stock":20},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[12]},"id":"G\u1ea1o","input_price":0,"itemid":14159000497,"modelid":"152817866084","name":"G\u1ea1o","price":20000,"price_before_discount":20000,"sku":"14159000497_5","stock":10}],"mtime":1685111769,"name":"<HOMI BEAUTY> M\u1eb7t N\u1ea1 Natureby H\u00e0n Qu\u1ed1c Ch\u00ednh H\u00e3ng HMB018","objectID":"shopee.vn_woo_402282834_14159000497_li1eqmxq","origin_cate":"100898 - S\u1eafc \u0110\u1eb9p > Ch\u0103m s\u00f3c da m\u1eb7t > M\u1eb7t n\u1ea1","origin_images":["https://cf.shopee.vn/file/d2f3d7921f17f77e1c6c9b4719aaf950","https://cf.shopee.vn/file/9e8f0e9a46789b0abc7da109f8131ace","https://cf.shopee.vn/file/da12dcd4ecb76b413912fdd2229d89be","https://cf.shopee.vn/file/6145402ce6159dc93117a9e4f3b85026","https://cf.shopee.vn/file/39805027cd5024f38866eb66fe1fbb1a","https://cf.shopee.vn/file/812858846d2e52361c02530d1123e608","https://cf.shopee.vn/file/bf512017f8f7eb0b5c403199b9efbc19"],"origin_shop":"homikids29","origin_url":"https://shopee.vn/!-i.402282834.14159000497","price":20000,"ptime":0,"shipxanh_ctime":1685109646,"shipxanh_height":15,"shipxanh_include_video":false,"shipxanh_length":15,"shipxanh_mtime":1685109647,"shipxanh_sku":"14159000497","shipxanh_unit":"1","shipxanh_weight":1000,"shipxanh_width":25,"shopid":402282834,"source":"shopee.vn","stock":84,"targets":["0"],"text_match":100,"tier_variations":[{"name":"ph\u00e2n lo\u1ea1i","options":["Nh\u00e2n s\u00e2m","D\u01b0a leo","C\u00e0 chua","\u1ed0c s\u00ean","Khoai t\u00e2y","Rong bi\u1ec3n","B\u01a1","Vi\u1ec7t qu\u1ea5t","M\u1eadt ong","Colagen","Tr\u00e0 xanh","Nha \u0111am","G\u1ea1o"],"properties":[],"summed_stocks":null,"type":0}],"valid_to_copy":true,"video_info_list":[{"video_id":""}]},"highlights":[],"text_match":100},{"document":{"accessible_to_organization_id":"NEHGPh7H6OWYvYHs67ocUmQ3NZh1","attributes":[],"brand":"Natureby","brand_id":2435644,"categories":[],"condition":1,"copy_to":"woo","ctime":1685109646,"current_brand":{"id":0,"name":"No Brand"},"current_cate":"15 - Ch\u01b0a ph\u00e2n lo\u1ea1i","description":"<p><p>M\u1eb7t N\u1ea1 Natureby H\u00e0n Qu\u1ed1c Ch\u00ednh H\u00e3ng</p><p>Qui C\u00e1ch: 1 Mi\u1ebfng </p><p>Xu\u1ea5t x\u1ee9: 100% H\u00e0n Qu\u1ed1c</p><p>HSD: 2024 (t\u00f9y t\u1eebng lo\u1ea1i m\u00e0 NG\u00c0Y TH\u00c1NG s\u1ea3n xu\u1ea5t kh\u00e1c nhau \u1ea1)</p><p>M\u1eb7t n\u1ea1 d\u01b0\u1ee1ng da Natureby ch\u00ednh h\u00e3ng H\u00e0n Qu\u1ed1c</p><p>C\u00f3 ph\u1ea3i b\u1ea1n r\u1ea5t b\u1eadn r\u1ed9n v\u00e0 kh\u00f4ng c\u00f3 th\u1eddi gian d\u00e0nh cho b\u1ea3n th\u00e2n ?</p><p>C\u00f3 ph\u1ea3i b\u1ea1n r\u1ea5t mu\u1ed1n l\u00e0m \u0111\u1eb9p m\u00e0 s\u1ee3 t\u1ed1n th\u1eddi gian ?</p><p>C\u00f3 ph\u1ea3i b\u1ea1n r\u1ea5t mu\u1ed1n s\u1ee1 h\u1eefu m\u1ed9t g\u01b0\u01a1ng m\u1eb7t tr\u1eafng s\u00e1ng m\u1ecbn m\u00e0ng v\u00e0 kh\u00f4ng b\u1ecb n\u1ed5i m\u1ee5n</p><p>V\u00e0 \u0111\u00e2y l\u00e0 gi\u1ea3i ph\u00e1p d\u00e0nh cho b\u1ea1n ====> [S\u1ec9 Gi\u00e1 T\u1ed1t] 14 Lo\u1ea1i M\u1eb7t N\u1ea1 Natureby H\u00e0n Qu\u1ed1c</p><p>-Natureby l\u00e0 m\u1ed9t trong nh\u1eefng m\u1ef9 ph\u1ea9m H\u00e0n Qu\u1ed1c \u0111\u01b0\u1ee3c ph\u00e1i n\u1eef tin d\u00f9ng. Natureby \u0111\u01b0\u1ee3c chi\u1ebft  100% t\u1eeb thi\u00ean nhi\u00ean \u0111\u1eb7t bi\u1ec7t an to\u00e0n cho m\u1ecdi lo\u1ea1i da </p><p>-H\u1ed2NG S\u00c2M: Chi\u1ebft xu\u1ea5t h\u1ed3ng s\u00e2m H\u00e0n Qu\u1ed1c gi\u00fap ph\u1ee5c da b\u1ecb h\u01b0 t\u1ed5n, nu\u00f4i d\u01b0\u1ee1ng da, d\u01b0\u1ee1ng \u1ea9m gi\u00fap da m\u1ec1m m\u1ea1i c\u0103ng b\u00f3ng v\u00e0 ng\u00e0y c\u00e0ng h\u1ed3ng h\u00e0o h\u01a1n.</p><p>-RONG BI\u1ec2N: chi\u1ebft xu\u1ea5t t\u1eeb rong bi\u1ec3n gi\u00fap da c\u0103ng m\u01b0\u1ee3t, r\u1ea1ng r\u1ee1 v\u00e0 \u0111\u1ea7y s\u1ee9c s\u1ed1ng. S\u1ea3n ph\u1ea9m v\u1edbi t\u00e1c d\u1ee5ng c\u1ea5p n\u01b0\u1edbc th\u1ea7n k\u00ec, kh\u00f4ng ch\u1ee9a c\u00e1c ch\u1ea5t c\u00f3 h\u1ea1i cho da, th\u00e0nh ph\u1ea7n chi\u1ebft xu\u1ea5t thi\u00ean nhi\u00ean l\u00e0nh t\u00ednh, mang l\u1ea1i \u0111\u1ed9 \u1ea9m l\u00fd t\u01b0\u1edfng, l\u00e0n da c\u0103ng m\u01b0\u1ee3t v\u00e0 \u0111\u01b0\u1ee3c c\u1ea3i thi\u1ec7n r\u00f5 r\u1ec7t.</p><p>-NHA \u0110AM : M\u1eb7t n\u1ea1 chi\u1ebft xu\u1ea5t nha \u0111am l\u00e0m d\u1ecbu da, c\u00f3 kh\u1ea3 n\u0103ng l\u00e0m x\u1eb9p m\u1ee5n, gi\u1ea3m c\u00e1c v\u1ebft s\u01b0ng t\u1ea5y do m\u1ee5n g\u00e2y ra, ch\u1ea5t gel c\u1ee7a nha \u0111am c\u00f3 nhi\u1ec7m v\u1ee5 t\u1ea9y c\u00e1c t\u1ebf b\u00e0o ch\u1ebft v\u00e0 gi\u1eef cho da lu\u00f4n t\u01b0\u01a1i tr\u1ebb. Gi\u1eef \u0111\u01b0\u1ee3c \u0111\u1ed9 \u1ea9m c\u1ea9n thi\u1ebft tr\u00e1nh hi\u1ec7n t\u01b0\u1ee3ng n\u1ee9t n\u1ebb, x\u00f3a v\u1ebft nh\u0103n v\u00e0 th\u00e2m quanh v\u00f9ng m\u00ed m\u1eaft sau nh\u1eefng ng\u00e0y l\u00e0m vi\u1ec7c m\u1ec7t m\u1ecfi.</p><p>-D\u01afA LEO: d\u01b0\u1ee1ng tr\u1eafng s\u00e1ng da, gi\u1ea3m s\u01b0ng t\u1ea5y v\u00f9ng da d\u01b0\u1edbi m\u1eaft.</p><p>-L\u1ef1u: d\u01b0\u1ee1ng tr\u1eafng s\u00e1ng da v\u00e0 ch\u1ed1ng l\u00e3o h\u00f3a</p><p>-Hoa h\u1ed3ng: d\u01b0\u1ee1ng s\u00e1ng tr\u1eafng da, x\u00f3a m\u1edd v\u1ebft th\u00e2m, se kh\u00edt l\u1ed7 ch\u00e2n l\u00f4ng</p><p>-G\u1ea1o :  b\u1ed9t g\u1ea1o c\u0169ng gi\u00fap l\u00e0m m\u1edd n\u00e1m, x\u00f3a t\u00e0n nhang h\u1eefu hi\u1ec7u ,T\u00ecnh tr\u1ea1ng c\u00e1c l\u1ed7 ch\u00e2n l\u00f4ng to, x\u1ec9n m\u00e0u d\u1ec5 d\u00e0ng \u0111\u01b0\u1ee3c kh\u1eafc ph\u1ee5c</p><p>-\u1ed0c s\u00ean : l\u00e0m m\u1edd v\u1ebft nh\u0103n hi\u1ec7u qu\u1ea3, \u0111\u1ed3ng th\u1eddi nu\u00f4i d\u01b0\u1ee1ng v\u00e0 t\u0103ng \u0111\u1ed9 \u0111\u00e0n h\u1ed3i cho da, ph\u1ee5c h\u1ed3i s\u1ef1 t\u01b0\u01a1i tr\u1ebb cho l\u00e0n da l\u00e3o h\u00f3a.</p><p>-H\u01b0\u1edbng d\u1eabn s\u1eed d\u1ee5ng:</p><p>+L\u00e0m s\u1ea1ch da m\u1eb7t b\u1eb1ng n\u01b0\u1edbc \u1ea5m ho\u1eb7c s\u1eefa r\u1eeda m\u1eb7t</p><p>+B\u1eaft \u0111\u1ea7u \u0111\u1eafp m\u1eb7t n\u1ea1 t\u1eeb v\u00f9ng tr\u00e1n sao cho ph\u1ee7 \u0111\u1ec1u h\u1ebft g\u01b0\u01a1ng m\u1eb7t</p><p>+Th\u01b0 gi\u00e3n kho\u1ea3ng 20 - 30 ph\u00fat</p><p>+L\u1ed9t b\u1ecf m\u1eb7t n\u1ea1, nh\u1eb9 nh\u00e0ng thoa tinh ch\u1ea5t c\u00f2n s\u00f3t l\u1ea1i tr\u00ean m\u1eb7t n\u1ea1 \u0111\u1ec3 cho d\u01b0\u1ee1ng ch\u1ea5t \u0111\u01b0\u1ee3c h\u1ea5p th\u1ee5 s\u00e2u v\u00e0o b\u00ean trong da r\u1ed3i r\u1eeda l\u1ea1i b\u1eb1ng n\u01b0\u1edbc s\u1ea1ch.</p><p>+1 tu\u1ea7n n\u00ean \u0111\u1eafp m\u1eb7t n\u1ea1 2 - 3 l\u1ea7n \u0111\u1ec3 c\u00f3 1 l\u00e0n da \u0111\u1eb9p, duy tr\u00ec tu\u1ed5i thanh xu\u00e2n nh\u00e1 \ud83d\ude1a\ud83d\ude1a</p><p>-H\u00c3Y T\u01af\u1edeNG T\u01af\u1ee2NG m\u1ecdi ng\u01b0\u1eddi tr\u1ea7m tr\u1ed3 v\u00ec b\u1ea1n c\u00f3 m\u1ed9t khu\u00f4n m\u1eb7t m\u1ecbn m\u00e0ng tr\u1eafng s\u00e1ng v\u00e0 n\u00f3i kh\u00f4ng v\u1edbi m\u1ee5n </p><p>#matna #matnaduongda #matnatrangda #natureby #naturebyhanquoc #naturebychinhhang #naturebynhadam</p></p>","errors":[],"flex_image_id_selected":"1685110188","flex_images":[{"id":"1685110188","images":["https://cf.shopee.vn/file/d2f3d7921f17f77e1c6c9b4719aaf950","https://cf.shopee.vn/file/9e8f0e9a46789b0abc7da109f8131ace","https://cf.shopee.vn/file/da12dcd4ecb76b413912fdd2229d89be","https://cf.shopee.vn/file/6145402ce6159dc93117a9e4f3b85026","https://cf.shopee.vn/file/39805027cd5024f38866eb66fe1fbb1a","https://cf.shopee.vn/file/812858846d2e52361c02530d1123e608","https://cf.shopee.vn/file/bf512017f8f7eb0b5c403199b9efbc19"],"is_default":true,"name":"ver 1","tags":["sx_default_desc"],"title_tags":["sx_default_title"]}],"highlights":"<HOMI BEAUTY> M\u1eb7t N\u1ea1 Natureby H\u00e0n Qu\u1ed1c Ch\u00ednh H\u00e3ng HMB018","history_post":[{"auto":false,"ctime":1685110866,"shop_id":"woo_homikids.vn","shop_name":"homikids.vn","success_id":1607},{"auto":false,"ctime":1685110649,"shop_id":"woo_homikids.vn","shop_name":"homikids.vn","success_id":1414}],"id":"shopee.vn_woo_402282834_14159000497_li1eqmxq","images":["https://cf.shopee.vn/file/d2f3d7921f17f77e1c6c9b4719aaf950","https://cf.shopee.vn/file/9e8f0e9a46789b0abc7da109f8131ace","https://cf.shopee.vn/file/da12dcd4ecb76b413912fdd2229d89be","https://cf.shopee.vn/file/6145402ce6159dc93117a9e4f3b85026","https://cf.shopee.vn/file/39805027cd5024f38866eb66fe1fbb1a","https://cf.shopee.vn/file/812858846d2e52361c02530d1123e608","https://cf.shopee.vn/file/bf512017f8f7eb0b5c403199b9efbc19"],"images_size":[],"is_official_shop":false,"itemid":14159000497,"models":[{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[0]},"id":"Nh\u00e2n s\u00e2m","input_price":0,"itemid":14159000497,"modelid":"160584535243","name":"Nh\u00e2n s\u00e2m","price":20000,"price_before_discount":20000,"sku":"14159000497_10","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[1]},"id":"D\u01b0a leo","input_price":0,"itemid":14159000497,"modelid":"152817866085","name":"D\u01b0a leo","price":20000,"price_before_discount":20000,"sku":"14159000497_11","stock":9},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[2]},"id":"C\u00e0 chua","input_price":0,"itemid":14159000497,"modelid":"152817866086","name":"C\u00e0 chua","price":20000,"price_before_discount":20000,"sku":"14159000497_6","stock":9},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[3]},"id":"\u1ed0c s\u00ean","input_price":0,"itemid":14159000497,"modelid":"152817866080","name":"\u1ed0c s\u00ean","price":20000,"price_before_discount":20000,"sku":"14159000497_1","stock":6},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[4]},"id":"Khoai t\u00e2y","input_price":0,"itemid":14159000497,"modelid":"160584535244","name":"Khoai t\u00e2y","price":20000,"price_before_discount":20000,"sku":"14159000497_8","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[5]},"id":"Rong bi\u1ec3n","input_price":0,"itemid":14159000497,"modelid":"152817866087","name":"Rong bi\u1ec3n","price":20000,"price_before_discount":20000,"sku":"14159000497_12","stock":10},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[6]},"id":"B\u01a1","input_price":0,"itemid":14159000497,"modelid":"160584535240","name":"B\u01a1","price":20000,"price_before_discount":20000,"sku":"14159000497_7","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[7]},"id":"Vi\u1ec7t qu\u1ea5t","input_price":0,"itemid":14159000497,"modelid":"160584535241","name":"Vi\u1ec7t qu\u1ea5t","price":20000,"price_before_discount":20000,"sku":"14159000497_13","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[8]},"id":"M\u1eadt ong","input_price":0,"itemid":14159000497,"modelid":"160584535242","name":"M\u1eadt ong","price":20000,"price_before_discount":20000,"sku":"14159000497_9","stock":10},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[9]},"id":"Colagen","input_price":0,"itemid":14159000497,"modelid":"152817866081","name":"Colagen","price":20000,"price_before_discount":20000,"sku":"14159000497_4","stock":0},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[10]},"id":"Tr\u00e0 xanh","input_price":0,"itemid":14159000497,"modelid":"152817866082","name":"Tr\u00e0 xanh","price":20000,"price_before_discount":20000,"sku":"14159000497_2","stock":10},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[11]},"id":"Nha \u0111am","input_price":0,"itemid":14159000497,"modelid":"152817866083","name":"Nha \u0111am","price":20000,"price_before_discount":20000,"sku":"14159000497_3","stock":20},{"extinfo":{"estimated_days":2,"group_buy_info":null,"is_pre_order":false,"tier_index":[12]},"id":"G\u1ea1o","input_price":0,"itemid":14159000497,"modelid":"152817866084","name":"G\u1ea1o","price":20000,"price_before_discount":20000,"sku":"14159000497_5","stock":10}],"mtime":1685111769,"name":"<HOMI BEAUTY> M\u1eb7t N\u1ea1 Natureby H\u00e0n Qu\u1ed1c Ch\u00ednh H\u00e3ng HMB018","objectID":"shopee.vn_woo_402282834_14159000497_li1eqmxq","origin_cate":"100898 - S\u1eafc \u0110\u1eb9p > Ch\u0103m s\u00f3c da m\u1eb7t > M\u1eb7t n\u1ea1","origin_images":["https://cf.shopee.vn/file/d2f3d7921f17f77e1c6c9b4719aaf950","https://cf.shopee.vn/file/9e8f0e9a46789b0abc7da109f8131ace","https://cf.shopee.vn/file/da12dcd4ecb76b413912fdd2229d89be","https://cf.shopee.vn/file/6145402ce6159dc93117a9e4f3b85026","https://cf.shopee.vn/file/39805027cd5024f38866eb66fe1fbb1a","https://cf.shopee.vn/file/812858846d2e52361c02530d1123e608","https://cf.shopee.vn/file/bf512017f8f7eb0b5c403199b9efbc19"],"origin_shop":"homikids29","origin_url":"https://shopee.vn/!-i.402282834.14159000497","price":20000,"ptime":0,"shipxanh_ctime":1685109646,"shipxanh_height":15,"shipxanh_include_video":false,"shipxanh_length":15,"shipxanh_mtime":1685109647,"shipxanh_sku":"14159000497","shipxanh_unit":"1","shipxanh_weight":1000,"shipxanh_width":25,"shopid":402282834,"source":"shopee.vn","stock":84,"targets":["0"],"text_match":100,"tier_variations":[{"name":"ph\u00e2n lo\u1ea1i","options":["Nh\u00e2n s\u00e2m","D\u01b0a leo","C\u00e0 chua","\u1ed0c s\u00ean","Khoai t\u00e2y","Rong bi\u1ec3n","B\u01a1","Vi\u1ec7t qu\u1ea5t","M\u1eadt ong","Colagen","Tr\u00e0 xanh","Nha \u0111am","G\u1ea1o"],"properties":[],"summed_stocks":null,"type":0}],"valid_to_copy":true,"video_info_list":[{"video_id":""}]},"highlights":[],"text_match":100}],"out_of":1970357,"page":1,"request_params":{"collection_name":"products_copy","per_page":250,"q":"*"},"search_cutoff":false,"search_time_ms":88}]}`
//
//	// Parse the input data
//	var shopeeProduct ShopeeProduct
//	err := json.Unmarshal([]byte(inputData), &shopeeProduct)
//	if err != nil {
//		fmt.Printf("Failed to parse input data: %v", err)
//		return
//	}
//
//	// Convert Shopee product data to WooCommerce product data
//	wooCommerceProduct := ConvertToWooCommerceProduct(shopeeProduct)
//
//	// Print the converted WooCommerce product data
//	convertedData, err := json.MarshalIndent(wooCommerceProduct, "", "  ")
//	if err != nil {
//		fmt.Printf("Failed to marshal converted data: %v", err)
//		return
//	}
//	fmt.Println(string(convertedData))
//}
//
//// ConvertToWooCommerceProduct converts Shopee product data to WooCommerce product data
//func ConvertToWooCommerceProduct(shopeeProduct ShopeeProduct) WooCommerceProduct {
//	wooCommerceProduct := WooCommerceProduct{
//		Name:         shopeeProduct.Name,
//		Type:         "simple",
//		RegularPrice: fmt.Sprintf("%.2f", float64(shopeeProduct.Price)/100),
//		Description:  shopeeProduct.Description,
//		Images:       make([]WooCommerceImage, 0),
//		Categories:   make([]WooCommerceCategory, 0),
//		Attributes:   make([]WooCommerceAttribute, 0),
//		//Weight:       fmt.Sprintf("%.2f", float64(shopeeProduct.Weight)/1000),
//		//Dimensions: WooCommerceDimensions{
//		//	Length: fmt.Sprintf("%.2f", float64(shopeeProduct.Length)/10),
//		//	Width:  fmt.Sprintf("%.2f", float64(shopeeProduct.Width)/10),
//		//	Height: fmt.Sprintf("%.2f", float64(shopeeProduct.Height)/10),
//		//},
//		VideoURL: shopeeProduct.VideoURL,
//	}
//
//	for _, image := range shopeeProduct.Images {
//		wooCommerceProduct.Images = append(wooCommerceProduct.Images, WooCommerceImage{Src: image.URL})
//	}
//
//	for _, category := range shopeeProduct.Categories {
//		wooCommerceProduct.Categories = append(wooCommerceProduct.Categories, WooCommerceCategory{ID: convertCategoryToID(category)})
//	}
//
//	for _, attribute := range shopeeProduct.Attributes {
//		wooCommerceProduct.Attributes = append(wooCommerceProduct.Attributes, WooCommerceAttribute{
//			Name:   attribute.Name,
//			Option: attribute.Options,
//		})
//	}
//
//	return wooCommerceProduct
//}
//
//// Helper function to convert a category name to its corresponding ID in WooCommerce
//func convertCategoryToID(category string) int {
//	// Perform the necessary mapping or logic to obtain the appropriate category ID in WooCommerce
//	// Replace the following example logic with your actual implementation
//	switch category {
//	case "Category 1":
//		return 1
//	case "Category 2":
//		return 2
//	default:
//		return 0
//	}
//}
