package seller

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func NewSmsProvider(productName string, sellerPhone string, sellerUUID string) *SmsProvider {
	return &SmsProvider{
		ProductName: productName,
		SellerPhone: sellerPhone,
		SellerUUID:  sellerUUID,
	}
}

type SmsProvider struct {
	ProductName string `json:"product_name"`
	SellerPhone string `json:"seller_phone"`
	SellerUUID  string `json:"seller_uuid"`
}

func (sp *SmsProvider) StockChanged(oldStock int, newStock int) {
	log.Info().Msg(fmt.Sprintf("SMS Warning sent to %s (Phone: %s): %s Product stock changed", sp.SellerUUID, sp.SellerPhone, sp.ProductName))
}
