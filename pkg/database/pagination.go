package database

// Option - Option type
type Option func(*Pagination)

// Size -
func Size(size uint32) Option {
	return func(c *Pagination) {
		c.size = size
	}
}

// Token -
func Token(token string) Option {
	return func(c *Pagination) {
		c.token = token
	}
}

// Pagination -
type Pagination struct {
	size  uint32
	token string
}

// NewPagination -
func NewPagination(opts ...Option) Pagination {
	pagination := &Pagination{}

	// Custom options
	for _, opt := range opts {
		opt(pagination)
	}

	if pagination.size == 0 {
		pagination.size = _defaultPageSize
	}

	return *pagination
}

// PageSize -
func (p Pagination) PageSize() uint32 {
	return p.size
}

// Token -
func (p Pagination) Token() string {
	return p.token
}

// EncodedContinuousToken -
type EncodedContinuousToken interface {
	// String returns the string representation of the continuous token.
	String() string
	// Decode decodes the continuous token from a string
	Decode() (ContinuousToken, error)
}

// ContinuousToken -
type ContinuousToken interface {
	// Encode encodes the continuous token to a string.
	Encode() EncodedContinuousToken
}
