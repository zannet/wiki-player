package interfaces

// UserModel is the Interface for User models
type UserModel interface {
	User(field, value string) (UserModel, error)
	Update() error
	Create() (string, error)
	Delete(nonce string) error
}
