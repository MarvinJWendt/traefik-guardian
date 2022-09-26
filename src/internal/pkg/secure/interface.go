package secure

type HashResolver interface {
	Check(h string, password string) bool
}
