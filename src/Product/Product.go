package Product

import "CrawlerBot/ProductDetail"

type Product struct {
	PersianName    string                       `json:"persian_name"`
	EnglishName    string                       `json:"english_name"`
	BrandOwner     string                       `json:"brand_owner"`
	LicenseHolder  string                       `json:"license_holder"`
	Price          string                       `json:"price"`
	Packaging      string                       `json:"packaging"`
	ProductCode    string                       `json:"product_code"`
	GenericCode    string                       `json:"generic_code"`
	ProductDetails ProductDetail.ProductDetails `json:"product_details"`
}
