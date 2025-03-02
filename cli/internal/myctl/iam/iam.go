package iam

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "iam",
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
	return "http://localhost:11080/api/iam"
}

func (self *CLIController) GetKeystoneVersion(ctx context.Context, res *GetKeystoneVersionResponse, err error) {
	fmt.Println(res.JSON200.Version.Id)
}

func (self *CLIController) CreateKeystoneFederationAuthToken(ctx context.Context, res *CreateKeystoneFederationAuthTokenResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) CreateKeystoneToken(ctx context.Context, res *CreateKeystoneTokenResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneDomains(ctx context.Context, res *GetKeystoneDomainsResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) CreateKeystoneDomain(ctx context.Context, res *CreateKeystoneDomainResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) DeleteKeystoneDomainByID(ctx context.Context, res *DeleteKeystoneDomainByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneDomainByID(ctx context.Context, res *GetKeystoneDomainByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) UpdateKeystoneDomainByID(ctx context.Context, res *UpdateKeystoneDomainByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) UnassignKeystoneRoleFromUserDomain(ctx context.Context, res *UnassignKeystoneRoleFromUserDomainResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) AssignKeystoneRoleToUserDomain(ctx context.Context, res *AssignKeystoneRoleToUserDomainResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneGroups(ctx context.Context, res *GetKeystoneGroupsResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneGroupByID(ctx context.Context, res *GetKeystoneGroupByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneProjects(ctx context.Context, res *GetKeystoneProjectsResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) CreateKeystoneProject(ctx context.Context, res *CreateKeystoneProjectResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) DeleteKeystoneProjectByID(ctx context.Context, res *DeleteKeystoneProjectByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneProjectByID(ctx context.Context, res *GetKeystoneProjectByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) UpdateKeystoneProjectByID(ctx context.Context, res *UpdateKeystoneProjectByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) UnassignKeystoneRoleFromGroupProject(ctx context.Context, res *UnassignKeystoneRoleFromGroupProjectResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) AssignKeystoneRoleToGroupProject(ctx context.Context, res *AssignKeystoneRoleToGroupProjectResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) UnassignKeystoneRoleFromUserProject(ctx context.Context, res *UnassignKeystoneRoleFromUserProjectResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) AssignKeystoneRoleToUserProject(ctx context.Context, res *AssignKeystoneRoleToUserProjectResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneRoleAssignments(ctx context.Context, res *GetKeystoneRoleAssignmentsResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneRoles(ctx context.Context, res *GetKeystoneRolesResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) CreateKeystoneRole(ctx context.Context, res *CreateKeystoneRoleResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) DeleteKeystoneRoleByID(ctx context.Context, res *DeleteKeystoneRoleByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneRoleByID(ctx context.Context, res *GetKeystoneRoleByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) UpdateKeystoneRoleByID(ctx context.Context, res *UpdateKeystoneRoleByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneUsers(ctx context.Context, res *GetKeystoneUsersResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) CreateKeystoneUser(ctx context.Context, res *CreateKeystoneUserResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) DeleteKeystoneUserByID(ctx context.Context, res *DeleteKeystoneUserByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneUserByID(ctx context.Context, res *GetKeystoneUserByIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetKeystoneUserProjectsByUserID(ctx context.Context, res *GetKeystoneUserProjectsByUserIDResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) CreateKeystoneApplicationCredential(ctx context.Context, res *CreateKeystoneApplicationCredentialResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetPubkeys(ctx context.Context, res *GetPubkeysResponse, err error) {
	fmt.Println(res, err)
}

func (self *CLIController) GetWebUser(ctx context.Context, res *GetWebUserResponse, err error) {
	fmt.Println(res, err)
}
