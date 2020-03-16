package config

// Server holds server configs
type Server struct {
	Host        string `env:"SERVER_HOST" envDefault:"localhost"`
	Port        string `env:"SERVER_PORT" envDefault:"8080"`
	PathVersion string `env:"SERVER_PATH_VERSION" envDefault:"v1"`
	URISchema   string `env:"SERVER_URI_SCHEMA" envDefault:"http://"`

	GinMode       string `env:"GIN_MODE" envDefault:"debug"`
	SessionSecret string `env:"SESSION_SECRET" envDefault:"supersecret"`
}

// ListenEndpoint builds the endpoint string (host + port)
func (s *Server) ListenEndpoint() string {
	if s.Port == "80" {
		return s.Host
	}

	return s.Host + ":" + s.Port
}

// VersionedEndpoint builds the endpoint string (host + port + version)
func (s *Server) VersionedEndpoint(path string) string {
	return "/" + s.PathVersion + path
}

// SchemaVersionedEndpoint builds the schema endpoint string (schema + host + port + version)
func (s *Server) SchemaVersionedEndpoint(path string) string {
	if s.Port == "80" {
		return s.URISchema + s.Host + "/" + s.PathVersion + path
	}

	return s.URISchema + s.Host + ":" + s.Port + "/" + s.PathVersion + path
}
