package api

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"golang.org/x/oauth2/jws"

	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
)

var defaultHeader jws.Header = jws.Header{Algorithm: "RS256", Typ: "JWT"}

var dummyPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAx4fm7dngEmOULNmAs1IGZ9Apfzh+BkaQ1dzkmbUgpcoghucE
DZRnAGd2aPyB6skGMXUytWQvNYav0WTR00wFtX1ohWTfv68HGXJ8QXCpyoSKSSFY
fuP9X36wBSkSX9J5DVgiuzD5VBdzUISSmapjKm+DcbRALjz6OUIPEWi1Tjl6p5RK
1w41qdbmt7E5/kGhKLDuT7+M83g4VWhgIvaAXtnhklDAggilPPa8ZJ1IFe31lNlr
k4DRk38nc6sEutdf3RL7QoH7FBusI7uXV03DC6dwN1kP4GE7bjJhcRb/7jYt7CQ9
/E9Exz3c0yAp0yrTg0Fwh+qxfH9dKwN52S7SBwIDAQABAoIBAQCaCs26K07WY5Jt
3a2Cw3y2gPrIgTCqX6hJs7O5ByEhXZ8nBwsWANBUe4vrGaajQHdLj5OKfsIDrOvn
2NI1MqflqeAbu/kR32q3tq8/Rl+PPiwUsW3E6Pcf1orGMSNCXxeducF2iySySzh3
nSIhCG5uwJDWI7a4+9KiieFgK1pt/Iv30q1SQS8IEntTfXYwANQrfKUVMmVF9aIK
6/WZE2yd5+q3wVVIJ6jsmTzoDCX6QQkkJICIYwCkglmVy5AeTckOVwcXL0jqw5Kf
5/soZJQwLEyBoQq7Kbpa26QHq+CJONetPP8Ssy8MJJXBT+u/bSseMb3Zsr5cr43e
DJOhwsThAoGBAPY6rPKl2NT/K7XfRCGm1sbWjUQyDShscwuWJ5+kD0yudnT/ZEJ1
M3+KS/iOOAoHDdEDi9crRvMl0UfNa8MAcDKHflzxg2jg/QI+fTBjPP5GOX0lkZ9g
z6VePoVoQw2gpPFVNPPTxKfk27tEzbaffvOLGBEih0Kb7HTINkW8rIlzAoGBAM9y
1yr+jvfS1cGFtNU+Gotoihw2eMKtIqR03Yn3n0PK1nVCDKqwdUqCypz4+ml6cxRK
J8+Pfdh7D+ZJd4LEG6Y4QRDLuv5OA700tUoSHxMSNn3q9As4+T3MUyYxWKvTeu3U
f2NWP9ePU0lV8ttk7YlpVRaPQmc1qwooBA/z/8AdAoGAW9x0HWqmRICWTBnpjyxx
QGlW9rQ9mHEtUotIaRSJ6K/F3cxSGUEkX1a3FRnp6kPLcckC6NlqdNgNBd6rb2rA
cPl/uSkZP42Als+9YMoFPU/xrrDPbUhu72EDrj3Bllnyb168jKLa4VBOccUvggxr
Dm08I1hgYgdN5huzs7y6GeUCgYEAj+AZJSOJ6o1aXS6rfV3mMRve9bQ9yt8jcKXw
5HhOCEmMtaSKfnOF1Ziih34Sxsb7O2428DiX0mV/YHtBnPsAJidL0SdLWIapBzeg
KHArByIRkwE6IvJvwpGMdaex1PIGhx5i/3VZL9qiq/ElT05PhIb+UXgoWMabCp84
OgxDK20CgYAeaFo8BdQ7FmVX2+EEejF+8xSge6WVLtkaon8bqcn6P0O8lLypoOhd
mJAYH8WU+UAy9pecUnDZj14LAGNVmYcse8HFX71MoshnvCTFEPVo4rZxIAGwMpeJ
5jgQ3slYLpqrGlcbLgUXBUgzEO684Wk/UV9DFPlHALVqCfXQ9dpJPg==
-----END RSA PRIVATE KEY-----`)

func ParseKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block != nil {
		key = block.Bytes
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		parsedKey, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("private key should be a PEM or plain PKCS1 or PKCS8; parse error: %v", err)
		}
	}
	parsed, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is invalid")
	}
	return parsed, nil
}

func (self *API) CreateKeystoneToken(ctx context.Context, input *oapi.CreateKeystoneTokenInput) (*oapi.KeystoneToken, string, error) {
	// dbProjects, err := self.db.FindProjects(ctx, &db.FindProjectsInput{})

	// if err != nil {
	// 	return nil, tlog.WrapError(ctx, err, "failed to self.db.FindProjects")
	// }

	// projects := []oapi.Project{}
	// for _, dbProject := range dbProjects {
	// 	projects = append(projects, oapi.Project{
	// 		Id:   dbProject.ID,
	// 		Name: dbProject.Name,
	// 	})
	// }
	pk, err := ParseKey(dummyPrivateKey)
	if err != nil {
		return nil, "", err
	}

	claimSet := &jws.ClaimSet{
		Iss:           "hoge@example.com",
		Scope:         "user",
		Aud:           "tokenurl",
		Exp:           time.Now().Add(time.Duration(10) * time.Second).Unix(),
		PrivateClaims: map[string]interface{}{"sub": "hoge"},
	}

	payload, err := jws.Encode(&defaultHeader, claimSet, pk)
	if err != nil {
		return nil, "", err
	}

	time.Sleep(1 * time.Second)

	if err := jws.Verify(payload, &pk.PublicKey); err != nil {
		return nil, "", err
	}

	if data, err := jws.Decode(payload); err != nil {
		return nil, "", err
	} else {
		fmt.Println("DEBUG decoded data", time.Unix(data.Exp, 0).Round(0).Before(time.Now()))
	}

	fmt.Println("DEBUG paylowd", payload)

	token := oapi.KeystoneToken{
		Token: oapi.KeystoneTokenData{
			AuditIds:  []string{"audit_id1", "audit_id2"},
			Methods:   []string{"password"},
			ExpiresAt: time.Now(),
			IssuedAt:  time.Now(),
			User: oapi.KeystoneTokenUser{
				Domain: oapi.KeystoneTokenDomain{
					Id:   "domain_id",
					Name: "domain_name",
				},
				Id:                "user_id",
				Name:              "user_name",
				PasswordExpiresAt: time.Now(),
			},
			Project: oapi.KeystoneTokenProject{
				Domain: oapi.KeystoneTokenDomain{
					Id:   "domain_id",
					Name: "domain_name",
				},
				Id:   "project_id",
				Name: "project_name",
			},
			Roles: []oapi.KeystoneTokenRole{
				{
					Id:   "role_id",
					Name: "role_name",
				},
			},
			Catalog: []oapi.KeystoneCatalog{
				{
					Id:   "catalog_id",
					Name: "keystone",
					Type: "identity",
					Endpoints: []oapi.KeystoneEndpoint{
						{
							Id:        "endpoint_id",
							Interface: "public",
							Region:    "region1",
							Url:       "http://localhost:10080/api/iam/keystone/v3",
						},
					},
				},
			},
		},
	}
	tokenStr := "token_id"
	return &token, tokenStr, nil
}
