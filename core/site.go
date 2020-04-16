package core

// Site contains site-specific parameters.
type Site struct {
	Config *Config

	JWTSecret []byte
}
