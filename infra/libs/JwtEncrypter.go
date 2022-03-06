package libs

type JwtEncrypter struct {
}

func (*JwtEncrypter) Generate() (string, error) {
	return "mock-jwt", nil
}
