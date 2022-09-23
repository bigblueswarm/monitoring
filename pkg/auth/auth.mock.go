// Package auth provides the application authentication
package auth

// ProviderMock is a mock implementation of Provider struct
type ProviderMock struct{}

var (
	// AuthenticateProviderMockFunction is the function that will be called when mock provider is called
	AuthenticateProviderMockFunction func(pwd string) bool
	// GenerateSessionProviderMockFunction generate a new authentication session
	GenerateSessionProviderMockFunction func() *Session
	// SaveSessionProviderMockFunction save session in system
	SaveSessionProviderMockFunction func(s *Session) error
	// RetrieveSessionProviderMockFunction retrieve session in system
	RetrieveSessionProviderMockFunction func(id string) (*Session, error)
	// DropSessionProviderMockFunction remove session from system
	DropSessionProviderMockFunction func(id string) error
)

// Authenticate is a mock implementation that check if provided password equals provider password
func (p *ProviderMock) Authenticate(pwd string) bool {
	return AuthenticateProviderMockFunction(pwd)
}

// GenerateSession is a mock implementation that generate a new authentication session
func (p *ProviderMock) GenerateSession() *Session {
	return GenerateSessionProviderMockFunction()
}

// SaveSession is a mock implementation that save session in system
func (p *ProviderMock) SaveSession(s *Session) error {
	return SaveSessionProviderMockFunction(s)
}

// RetrieveSession is a mock implementation that retrieve session in system
func (p *ProviderMock) RetrieveSession(id string) (*Session, error) {
	return RetrieveSessionProviderMockFunction(id)
}

// DropSession is a mock implementation that remove session from system
func (p *ProviderMock) DropSession(id string) error {
	return DropSessionProviderMockFunction(id)
}
