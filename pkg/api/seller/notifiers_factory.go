package seller

type NotifiersFactory struct {
	NotifierTypes []string `json:"notifier_types"`
}

func (factory *NotifiersFactory) CreateNotifiers(productName string, seller Seller) (listAllNotifiers []Notifier) {
	for _, currType := range factory.NotifierTypes {
		if currType == "sms" {
			listAllNotifiers = append(listAllNotifiers, NewSmsProvider(productName, seller.Phone, seller.UUID))
		} else if currType == "email" {
			listAllNotifiers = append(listAllNotifiers, NewEmailProvider(seller.Email))
		}
	}
	return listAllNotifiers
}
