package iam_auth

import (
	"context"
	"crypto"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/syunkitada/stadyapp/backends/libs/pkg/tlog"
)

type IAMAuth struct {
	conf    *Config
	Expires time.Duration
}

func New(conf *Config) *IAMAuth {
	return &IAMAuth{
		conf:    conf,
		Expires: time.Duration(conf.ExpiresSec) * time.Second,
	}
}

func (self *IAMAuth) GetPublicKey(ctx context.Context, key string) (crypto.PublicKey, error) {
	bytes, err := os.ReadFile(path.Join(self.conf.PublicKeyDir, key))
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to read public key")
	}

	publicKey, err := jwt.ParseEdPublicKeyFromPEM(bytes)
	if err != nil {
		return nil, tlog.WrapErr(ctx, err, "failed to parse public key")
	}

	return publicKey, nil
}

func (self *IAMAuth) AuthUserID(ctx context.Context, header http.Header) error {
	xuserID := header.Get("x-user-id")
	if xuserID == "" {
		return tlog.Err(ctx, echo.NewHTTPError(http.StatusUnauthorized, "x-user-id is not found"))
	}

	xdomainID := header.Get("x-domain-id")
	if xdomainID == "" {
		return tlog.Err(ctx, echo.NewHTTPError(http.StatusUnauthorized, "x-domain-id is not found"))
	}

	return nil
}

func (self *IAMAuth) AuthToken(ctx context.Context, header http.Header) error {
	xauthToken := header.Get("x-auth-token")
	if xauthToken == "" {
		return tlog.Err(ctx, echo.NewHTTPError(http.StatusUnauthorized, "x-auth-token is not found"))
	}

	tokenData, err := self.VerifyToken(ctx, xauthToken, header)
	if err != nil {
		return tlog.WrapErr(ctx, err, "failed to verify token")
	}

	header.Set("x-user-id", tokenData.UserID)
	header.Set("x-domain-id", tokenData.DomainID)
	header.Set("x-project-id", tokenData.ProjectID)
	header.Set("x-roles", tokenData.Roles)
	header.Set("x-catalog", tokenData.Catalog)
	header.Set("x-inherit", strconv.FormatBool(tokenData.Inherit))

	return nil
}

func (self *IAMAuth) VerifyToken(ctx context.Context, payload string, header http.Header) (*AuthData, error) {
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

func (self *IAMAuth) NewToken(ctx context.Context, authData AuthData) (string, error) {
	keyName, privateKey, err := self.GetPrivateKey(ctx)
	if err != nil {
		return "", tlog.WrapErr(ctx, err, "failed to get private key")
	}

	roles := []string{}
	for key := range authData.RoleSet {
		roles = append(roles, key)
	}
	authData.Roles = strings.Join(roles, ",")

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(self.Expires)),
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

	payload, err := token.SignedString(privateKey)
	if err != nil {
		return "", tlog.Err(ctx, err)
	}

	return payload, nil
}

func (self *IAMAuth) GetPrivateKey(ctx context.Context) (string, crypto.PrivateKey, error) {
	files, _ := os.ReadDir(self.conf.PrivateKeyDir)

	if len(files) == 0 {
		err := echo.NewHTTPError(http.StatusInternalServerError, "private key is not found")

		return "", nil, tlog.Err(ctx, err)
	}

	fileName := files[0].Name()
	bytes, err := os.ReadFile(path.Join(self.conf.PrivateKeyDir, fileName))

	if err != nil {
		return "", nil, tlog.Err(ctx, err)
	}

	privateKey, err := jwt.ParseEdPrivateKeyFromPEM(bytes)
	if err != nil {
		return "", nil, tlog.Err(ctx, err)
	}

	return fileName, privateKey, nil
}
