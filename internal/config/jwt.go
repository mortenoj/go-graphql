package config

// JWT holds JSON web token configs
type JWT struct {
	Secret    string `env:"AUTH_JWT_SECRET"`
	Algorithm string `env:"AUTH_JWT_SIGNING_ALGORITHM" envDefault:"HS512"`
}
