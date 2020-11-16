package config

type Config struct {
	Seller struct {
		NotificationTypes []string `yaml:"notification_types"`
	} `yaml:"seller"`
}
