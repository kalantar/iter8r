package utils

import (
	"context"
	"errors"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func GetClient(cf *genericclioptions.ConfigFlags) (*kubernetes.Clientset, error) {
	restConfig, err := cf.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

func GetExperiment(client *kubernetes.Clientset, ns string, nm string) (experiment *corev1.Secret, err error) {
	ctx := context.Background()

	// A name is provided; get this experiment, if it exists
	if len(nm) != 0 {
		experiment, err = client.CoreV1().Secrets(ns).Get(ctx, nm, metav1.GetOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				return nil, fmt.Errorf("experiment \"%s\" not found", nm)
			}
		}
		// verify that the job corresponds to an experiment
		if experiment != nil && !isExperiment(*experiment) {
			return nil, fmt.Errorf("experiment \"%s\" not found", nm)
		}

		return experiment, err
	}

	// There is no explict experiment name provided.
	// Get a list of all experiments.
	// Then select the one with the most recent create time.
	experiments, err := GetExperiments(client, ns)
	if err != nil {
		return experiment, err
	}

	// no experiments
	if len(experiments) == 0 {
		return experiment, errors.New("no experiments found")
	}

	for _, job := range experiments {
		if experiment == nil {
			experiment = &job
			continue
		}
		if job.ObjectMeta.CreationTimestamp.Time.After(experiment.ObjectMeta.CreationTimestamp.Time) {
			experiment = &job
		}
	}
	return experiment, nil
}

func GetExperiments(client *kubernetes.Clientset, ns string) (experiments []corev1.Secret, err error) {
	secrets, err := client.CoreV1().Secrets(ns).List(
		context.Background(), metav1.ListOptions{
			LabelSelector: "iter8/type=experiment",
		})
	if err != nil {
		return experiments, err
	}

	return jobListToExperimentJobList(*secrets), err
}

func jobListToExperimentJobList(secrets corev1.SecretList) (result []corev1.Secret) {
	for _, secret := range secrets.Items {
		if isExperiment(secret) {
			result = append(result, secret)
		}
	}
	return result
}

func isExperiment(e corev1.Secret) bool {
	return !strings.HasSuffix(e.GetName(), "-result")
}
