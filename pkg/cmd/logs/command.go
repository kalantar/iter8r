package logs

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

var example = `
# Display logs from an experiment execution
%[1]s logs -e experiment

# Display logs the most recent experiment execution
%[1]s logs
`

func NewCmd(factory cmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := newOptions(streams)

	cmd := &cobra.Command{
		Use:          "logs",
		Short:        "Retrieve logs from an experiment execution",
		Example:      fmt.Sprintf(example, "iter8"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.complete(factory, c, args); err != nil {
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
