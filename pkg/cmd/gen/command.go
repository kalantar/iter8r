package run

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

var example = `
# generate text output of a local experiment using the built-in Go template
%[1]s gen --local

# generate text output of the most recent experiment using the built-in Go template
%[1]s gen

# generate text output of a specific experiment using the built-in Go template
%[1]s gen -e experiment

# generate output from experiment using a custom Go template specified in the iter8.tpl file
%[1]s gen -o custom

# generate manifest for a remote experiment
%[1]s gen -o k8s
`

func NewCmd(factory cmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := newOptions(streams)

	cmd := &cobra.Command{
		Use:          "gen",
		Short:        "Format experiment spec and its result using a go template",
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

	cmd.Flags().StringVarP(&o.outputFormat, "outputFormat", "o", Text, fmt.Sprintf("%v | %v | %v ", Text, Custom, K8S))
	cmd.Flags().StringSliceVar(&o.overrides, "set", nil, "overrides of the form name=value")

	return cmd
}
