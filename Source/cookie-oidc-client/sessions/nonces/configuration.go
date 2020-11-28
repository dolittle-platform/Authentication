package nonces

// Configuration for nonce generation
type Configuration interface {
	// Length returns the size in bytes for generated nonces
	Size() int
}
