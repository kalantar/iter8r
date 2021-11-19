package run

import (
	"fmt"

	"github.com/spf13/cobra"
)

var example = `
# Run a local experiment locally
%[1]s run

# Run an experiment in a Kubernetes cluster
%[1]s gen -o k8s | kubectl apply -f -
`

func NewCmd() *cobra.Command {
	o := newOptions()

	cmd := &cobra.Command{
		Use:          "run",
		Short:        "Run an experiment locally",
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

	return cmd
}
