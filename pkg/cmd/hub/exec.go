package run

import (
	"errors"

	"github.com/kalantar/iter8r/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// complete sets all information needed for processing the command
func (o *Options) complete(cmd *cobra.Command, args []string) (err error) {
	viper.BindEnv("ITER8HUB")
	viper.SetDefault("ITER8HUB", "github.com/iter8-tools/iter8.git/mkdocs/docs/iter8hub/")
	return nil
}

// validate ensures that all required arguments and flag values are provided
func (o *Options) validate(cmd *cobra.Command, args []string) (err error) {
	if o.experiment == "" {
		return errors.New("an experiment must be specified")
	}
	return nil
}

// run runs the command
func (o *Options) run(cmd *cobra.Command, args []string) (err error) {
	_, err = utils.Execute(utils.Iter8Hub, "-e", o.experiment)
	return err
}
