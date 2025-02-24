package compute

import (
	"context"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "compute",
	Short: "CLI for mycloudstack",
}

type Controller struct {
	token string
}

func NewController() IController {
	token := os.Getenv("OS_TOKEN")

	return &Controller{token: token}
}

func (self *Controller) RequestEditorFn(ctx context.Context, req *http.Request) error {
	req.Header.Set("x-auth-token", self.token)
	return nil
}

func (self *Controller) GetBaseURL() string {
	return "http://localhost:11080/api/compute"
}
