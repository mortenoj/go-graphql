package config

// GraphQL holds graphql configs
type GraphQL struct {
	ServerPath        string `env:"GQL_SERVER_GRAPHQL_PATH" envDefault:"/graphql"`
	PlaygroundEnabled bool   `env:"GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED" envDefault:"true"`
	PlaygroundPath    string `env:"GQL_SERVER_GRAPHQL_PLAYGROUND_PATH" envDefault:"/"`
}
