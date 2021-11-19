package get

import (
	"fmt"

	assert "github.com/kalantar/iter8r/pkg/cmd/assert"
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
	experiments, err := utils.GetExperiments(o.client, o.namespace)
	if err != nil {
		return err
	}

	if len(experiments) == 0 {
		fmt.Println("no experiments found")
		return err
	}

	fmt.Printf("%-16s  %-9s  %-6s  %-14s\n", "NAME", "COMPLETED", "FAILED", "SLOS SATISFIED")
	for _, experiment := range experiments {
		_ = utils.FetchResultsAsFile(o.client, o.namespace, experiment.GetName())
		fmt.Printf("%-16s  %-9t  %-6t  %-14t\n", experiment.GetName(), isCompleted(), failed(), slos())
	}
	return nil
}

func isCompleted() bool {
	return evaluateCondition(assert.Completed)
}

func failed() bool {
	return !evaluateCondition(assert.NoFailure)
}

func slos() bool {
	return evaluateCondition(assert.SLOs)
}

func evaluateCondition(condition string) bool {
	_, err := utils.ExecuteSilent(utils.Iter8Assert, "-c", condition)
	return err == nil
}
