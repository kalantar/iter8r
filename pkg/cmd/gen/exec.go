package run

import (
	"errors"
	"regexp"

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
	// Currently support only a single override; name=
	if len(o.overrides) > 1 {
		return errors.New("only 1 override is supported")
	}

	// Currently support only name=
	for _, override := range o.overrides {
		re := regexp.MustCompile("name=")
		if nil == re.Find([]byte(override)) {
			return errors.New("only override of name is supported")
		}
	}

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

	_, err = utils.Execute(utils.Iter8Gen, "-o", o.outputFormat)
	return err
}
