package run

import (
	"github.com/kalantar/iter8r/pkg/utils"
	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

// complete sets all information needed for processing the command
func (o *Options) complete(factory cmdutil.Factory, cmd *cobra.Command, args []string) (err error) {
	o.namespace, _, err = factory.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}

	o.client, err = utils.GetClient(o.ConfigFlags)
	if err != nil {
		return err
	}

	return err
}

// validate ensures that all required arguments and flag values are provided
func (o *Options) validate(cmd *cobra.Command, args []string) (err error) {
	return nil
}

// run runs the command
func (o *Options) run(cmd *cobra.Command, args []string) (err error) {
	if !o.local {
		err = utils.FetchResultsAsFile(o.client, o.namespace, o.experiment)
		if err != nil {
			return err
		}
	}

	_, err = utils.Execute(utils.Iter8Assert, o.getArgs()...)
	return err
}

func (o *Options) getArgs() []string {
	// args := []string{utils.Iter8Assert}
	args := []string{}
	for _, v := range o.conds {
		args = append(args, "-c")
		args = append(args, v)
	}
	args = append(args, "-t")
	args = append(args, o.timeout.String())

	return args
}
