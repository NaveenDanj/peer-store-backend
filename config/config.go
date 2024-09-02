package config

type Config struct {
	FileCWD            string
	PrivatePEMFilePath string
	PublicPEMFilePath  string
	PassPhrase         string
}

var CONFIG = Config{
	FileCWD:            "./test",
	PrivatePEMFilePath: "./private_key.pem",
	PublicPEMFilePath:  "./public_key.pem",
	PassPhrase:         "SunnyDayInJuly2024WithABreeze123",
}
