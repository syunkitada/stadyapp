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

type CLIController struct {
	token string
}

func NewCLIController() ICLIController {
	token := os.Getenv("OS_TOKEN")

	return &CLIController{token: token}
}

func (self *CLIController) RequestEditorFn(ctx context.Context, req *http.Request) error {
	req.Header.Set("x-auth-token", self.token)
	return nil
}

func (self *CLIController) GetBaseURL() string {
	return "http://localhost:11080/api/compute"
}
