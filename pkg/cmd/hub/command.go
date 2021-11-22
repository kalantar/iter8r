package hub

import (
	"fmt"

	"github.com/spf13/cobra"
)

var example = `
# Download an experiment sample from the Iter8 hub
%[1]s hub -e experiment

# Download an experiment from another hub
ITER8HUB=url %[1]s hub -e experiment
`

func NewCmd() *cobra.Command {
	o := newOptions()

	cmd := &cobra.Command{
		Use:          "hub",
		Short:        "Download an experiment",
		Example:      fmt.Sprintf(example, "iter8"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.complete(c, args); err != nil {
				return err
			}
			if err := o.validate(c, args); err != nil {
				return err
			}
			if err := o.run(c, args); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&o.experiment, "experiment", "e", "", "experiment; if not specified, the most recently created one is used")

	return cmd
}
