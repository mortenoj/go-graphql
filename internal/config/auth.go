package config

// AuthProviders holds the different authentication provider configs
type AuthProviders struct {
	Google
	Auth0
}

// Google authentication provider config
type Google struct {
	Key    string   `env:"PROVIDER_GOOGLE_KEY"`
	Secret string   `env:"PROVIDER_GOOGLE_SECRET"`
	Scopes []string `env:"PROVIDER_GOOGLE_SCOPES" envSeparator:","`
}

// Auth0 authentication provider config
type Auth0 struct {
	Key    string   `env:"PROVIDER_AUTH0_KEY"`
	Secret string   `env:"PROVIDER_AUTH0_SECRET"`
	Scopes []string `env:"PROVIDER_AUTH0_SCOPES" envSeparator:","`
	Domain string   `env:"PROVIDER_AUTH0_DOMAIN"`
}
