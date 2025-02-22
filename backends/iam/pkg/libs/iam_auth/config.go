package iam_auth

type Config struct {
	PublicKeyDir  string
	PrivateKeyDir string
}

func GetDefaultConfig() Config {
	return Config{
		PublicKeyDir:  "/etc/iam/token_keys/public",
		PrivateKeyDir: "/etc/iam/token_keys/private",
	}
}
