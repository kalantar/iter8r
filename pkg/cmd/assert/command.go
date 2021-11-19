package run

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

var example = `
# Assert properties of an experiment
%[1]s assert
`

func NewCmd(factory cmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := newOptions(streams)

	cmd := &cobra.Command{
		Use:          "assert",
		Short:        "Assert if an experiment satisfies specified conditions",
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
	cmd.Flags().BoolVarP(&o.local, "local", "l", false, "use locally executed experiment; any cluster options are ignored")

	cmd.Flags().StringSliceVarP(&o.conds, "condition", "c", nil, fmt.Sprintf("%v | %v | %v | %v=<version number>", Completed, NoFailure, SLOs, SlosByPrefix))
	cmd.MarkFlagRequired("condition")
	cmd.Flags().DurationVarP(&o.timeout, "timeout", "t", 0, "timeout duration (e.g., 5s)")

	return cmd
}
