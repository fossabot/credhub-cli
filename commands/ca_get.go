package commands

import (
	"fmt"

	"github.com/pivotal-cf/cm-cli/actions"
	"github.com/pivotal-cf/cm-cli/client"
	"github.com/pivotal-cf/cm-cli/config"
	"github.com/pivotal-cf/cm-cli/repositories"
)

type CaGetCommand struct {
	CaIdentifier string `short:"n" required:"yes" long:"name" description:"Name of the CA to retrieve"`
}

func (cmd CaGetCommand) Execute([]string) error {
	caRepository := repositories.NewCaRepository(client.NewHttpClient())
	config := config.ReadConfig()
	action := actions.NewAction(caRepository, config)

	ca, err := action.DoAction(
		client.NewGetCaRequest(
			config.ApiURL,
			cmd.CaIdentifier), cmd.CaIdentifier)

	if err != nil {
		return err
	}

	fmt.Println(ca)

	return nil
}