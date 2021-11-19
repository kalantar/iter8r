package run

import (
	"github.com/kalantar/iter8r/pkg/utils"

	"github.com/spf13/cobra"
)

// complete sets all information needed for processing the command
func (o *Options) complete(cmd *cobra.Command, args []string) (err error) {
	return nil
}

// validate ensures that all required arguments and flag values are provided
func (o *Options) validate(cmd *cobra.Command, args []string) (err error) {
	return nil
}

// run runs the command
func (o *Options) run(cmd *cobra.Command, args []string) (err error) {
	_, err = utils.Execute(utils.Iter8Run)
	return err
}
