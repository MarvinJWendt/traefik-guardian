package secure

type PlaintextResolver struct{}

func (p *PlaintextResolver) Check(h string, password string) bool {
	return h == password
}
