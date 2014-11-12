package interfaces

// UserModelInterface is the Interface for User models
type UserModelInterface interface {
	User(field, value string) (UserModelInterface, error)
	Update() error
	Create() (string, error)
	Delete(nonce string) error
}
