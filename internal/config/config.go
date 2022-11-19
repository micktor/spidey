package config

// Defines the application configuration.
// Add to and create structs as needed.

// Database defines the settings
// to connect to a database.
type Database struct {
	Hostname string
	Port     string
	Database string
	Username string
	Password string
}

// HTTPServer defines the settings
// for running a http server, and it's middleware.
type HTTPServer struct {
	Debug    bool
	Hostname string
	Port     string
}

// Config contains all the application configuration.
type Config struct {
	Database   Database
	HTTPServer HTTPServer
}
