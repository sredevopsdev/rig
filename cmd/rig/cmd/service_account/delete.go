package service_account

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/rigdev/rig-go-api/api/v1/service_account"
	"github.com/rigdev/rig-go-sdk"
	"github.com/rigdev/rig/cmd/common"
	"github.com/rigdev/rig/pkg/uuid"
	"github.com/spf13/cobra"
)

func ServiceAccountDelete(ctx context.Context, cmd *cobra.Command, args []string, nc rig.Client) error {
	var id string
	var err error

	if len(args) > 0 {
		id = args[0]
	}

	if id == "" {
		id, err = common.PromptInput("ID:", common.ValidateNonEmptyOpt)
		if err != nil {
			return err
		}
	}

	_, err = uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = nc.ServiceAccount().Delete(ctx, &connect.Request[service_account.DeleteRequest]{
		Msg: &service_account.DeleteRequest{
			ServiceAccountId: id,
		},
	})
	if err != nil {
		return err
	}

	cmd.Println("Credential deleted")
	return nil
}
