package product

type product struct {
	ProductID  int    `json:"-"`
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Brand      string `json:"brand"`
	Stock      int    `json:"stock"`
	SellerUUID string `json:"seller_uuid"`
}

type productV2 struct {
	ProductID int           `json:"-"`
	UUID      string        `json:"uuid"`
	Name      string        `json:"name"`
	Brand     string        `json:"brand"`
	Stock     int           `json:"stock"`
	Seller    productSeller `json:"seller"`
}

type productSeller struct {
	UUID  string          `json:"uuid"`
	Links map[string]link `json:"_links"`
}

type link struct {
	Href string `json:"href"`
}
