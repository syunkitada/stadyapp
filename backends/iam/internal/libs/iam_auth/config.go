package iam_auth

type Config struct {
	PublicKeyDir  string
	PrivateKeyDir string
	ExpiresSec    int
}

func GetDefaultConfig() Config {
	return Config{
		PublicKeyDir:  "/etc/iam/token_keys/public",
		PrivateKeyDir: "/etc/iam/token_keys/private",
		ExpiresSec:    43200, //nolint:mnd // 12 hours
	}
}
