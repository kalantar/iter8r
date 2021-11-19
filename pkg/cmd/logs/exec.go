package logs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/kalantar/iter8r/pkg/utils"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	experiment, err := utils.GetExperiment(o.client, o.namespace, o.experiment)
	if err != nil {
		return err
	}

	pods, err := o.client.CoreV1().Pods(o.namespace).List(
		context.Background(), metav1.ListOptions{
			LabelSelector: fmt.Sprintf("job-name=%s", experiment.GetName()),
		})
	if err != nil {
		return err
	}

	pod := pods.Items[0]

	// cf. https://stackoverflow.com/a/53870271/3482067
	podLogOpts := corev1.PodLogOptions{}
	req := o.client.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		return errors.New("error in opening stream")
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return errors.New("error in copy information from podLogs to buf")
	}
	str := buf.String()

	fmt.Println(str)

	return err
}
