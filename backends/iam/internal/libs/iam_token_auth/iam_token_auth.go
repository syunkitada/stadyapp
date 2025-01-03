package iam_token_auth

import (
	"context"
	"crypto"
	"errors"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

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

type IAMTokenAuth struct {
	conf *Config
}

func New(conf *Config) *IAMTokenAuth {
	return &IAMTokenAuth{
		conf: conf,
	}
}

type AuthData struct {
	Domain  string `json:"1"`
	User    string `json:"2"`
	Project string `json:"3"`
	Roles   string `json:"4"`
	Catalog string `json:"5"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	KeyName  string   `json:"0"`
	AuthData AuthData `json:"1"`
}

func (self *IAMTokenAuth) GetPublicKey(ctx context.Context, key string) (crypto.PublicKey, error) {
	bytes, err := os.ReadFile(path.Join(self.conf.PublicKeyDir, key))
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to read public key")
	}

	pk, err := jwt.ParseEdPublicKeyFromPEM(bytes)
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to parse public key")
	}
	return pk, nil
}

func (self *IAMTokenAuth) GetPrivateKey(ctx context.Context) (string, crypto.PrivateKey, error) {
	files, _ := os.ReadDir(self.conf.PrivateKeyDir)

	for _, f := range files {
		bytes, err := os.ReadFile(path.Join(self.conf.PrivateKeyDir, f.Name()))
		if err != nil {
			return "", nil, err
		}

		pk, err := jwt.ParseEdPrivateKeyFromPEM(bytes)
		if err != nil {
			return "", nil, err
		}
		return f.Name(), pk, nil
	}

	return "", nil, errors.New("private key is not found")
}

func (self *IAMTokenAuth) AuthToken(ctx context.Context, header http.Header) error {
	xauthToken := header.Get("x-auth-token")
	if xauthToken == "" {
		return tlog.Err(ctx, echo.NewHTTPError(http.StatusUnauthorized, "x-auth-token is not found"))
	}

	tokenData, err := self.VerifyToken(ctx, xauthToken, header)
	if err != nil {
		return tlog.WrapErr(ctx, err, "failed to verify token")
	}

	header.Set("x-user-id", tokenData.User)

	return nil
}

func (self *IAMTokenAuth) VerifyToken(ctx context.Context, payload string, header http.Header) (*AuthData, error) {
	token, err := jwt.ParseWithClaims(payload, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusUnauthorized, "token is invalid"))
		}
		return self.GetPublicKey(ctx, claims.KeyName)
	})
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to parse token")
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		return &claims.AuthData, nil
	} else {
		return nil, tlog.Err(ctx, echo.NewHTTPError(http.StatusUnauthorized, "token is invalid"))
	}
}

const TokenScope = "iam.user"
const PrivateClaimsDataKey = "d"

func (self *IAMTokenAuth) NewToken(ctx context.Context, authData AuthData) (string, error) {
	keyName, pk, err := self.GetPrivateKey(ctx)
	if err != nil {
		return "", tlog.WrapErr(ctx, err, "failed to get private key")
	}

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
		KeyName:  keyName,
		AuthData: authData,
	}

	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)

	payload, err := token.SignedString(pk)

	return payload, nil
}

type AuthContext struct {
	UserID    string
	ProjectID string
}

type key int

const KeyAuthContext key = iota

func WithEchoContext(ectx echo.Context) context.Context {
	ctx := tlog.WithEchoContext(ectx)
	xuser := ectx.Request().Header.Get("x-user-id")
	AuthContext := AuthContext{UserID: xuser}
	ctx = context.WithValue(ctx, KeyAuthContext, &AuthContext)

	return ctx
}

func GetAuthContext(ctx context.Context) (*AuthContext, error) {
	authCtx, ok := ctx.Value(KeyAuthContext).(*AuthContext)
	if !ok {
		return nil, errors.New("auth context is not found")
	}
	return authCtx, nil
}
