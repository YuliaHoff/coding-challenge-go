package seller

func NewEmailProvider(email string) *EmailProvider {
	return &EmailProvider{SellerEmail: email}
}

type EmailProvider struct {
	SellerEmail string `json:"seller_email"`
}

func (ep *EmailProvider) StockChanged(oldStock int, newStock int) {

}
