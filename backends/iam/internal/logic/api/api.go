package api

import (
	"encoding/json"

	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/api"
	"github.com/syunkitada/stadyapp/backends/iam/internal/domain/db"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/config"
	"github.com/syunkitada/stadyapp/backends/iam/internal/iam-api/spec/oapi"
	"github.com/syunkitada/stadyapp/backends/iam/internal/libs/iam_auth"
)

type API struct {
	conf               *config.Config
	db                 db.IDB
	iamAuth            *iam_auth.IAMAuth
	keystoneCatalog    []oapi.KeystoneCatalog
	keystoneCatalogStr string
}

func New(conf *config.Config, db db.IDB, iamAuth *iam_auth.IAMAuth) api.IAPI { //nolint:ireturn
	keystoneCatalog, keystoneCatalogStr := convertKeystoneCatalog(conf)
	// keystoneCatalogBytes, err := json.Marshal(keystoneCatalog)
	// if err != nil {
	// 	panic(err)
	// }

	return &API{
		conf:               conf,
		db:                 db,
		iamAuth:            iamAuth,
		keystoneCatalog:    keystoneCatalog,
		keystoneCatalogStr: keystoneCatalogStr,
	}
}

func convertKeystoneCatalog(conf *config.Config) ([]oapi.KeystoneCatalog, string) {
	catalog2 := []map[string]interface{}{}

	catalog := []oapi.KeystoneCatalog{}
	for _, catalogConf := range conf.IAM.Catalogs {
		endpoints := []oapi.KeystoneEndpoint{}
		endpoints2 := []map[string]interface{}{}
		for _, endpoint := range catalogConf.Endpoints {
			endpoints = append(endpoints, oapi.KeystoneEndpoint{
				Interface: endpoint.Interface,
				Region:    endpoint.Region,
				Url:       endpoint.URL,
			})

			endpoints2 = append(endpoints2, map[string]interface{}{
				"interface": endpoint.Interface,
				"region":    endpoint.Region,
				"publicURL": endpoint.URL,
			})
		}

		catalog = append(catalog, oapi.KeystoneCatalog{
			Type:      catalogConf.Type,
			Name:      catalogConf.Name,
			Endpoints: endpoints,
		})

		catalog2 = append(catalog2, map[string]interface{}{
			"type":      catalogConf.Type,
			"endpoints": endpoints2,
		})
	}

	keystoneCatalogBytes, err := json.Marshal(&catalog2)
	if err != nil {
		panic(err)
	}

	return catalog, string(keystoneCatalogBytes)
}
