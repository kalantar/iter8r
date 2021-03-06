package run

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

const (
	Text   = "text"
	Custom = "custom"
	K8S    = "k8s"
)

type Options struct {
	Streams              genericclioptions.IOStreams
	ConfigFlags          *genericclioptions.ConfigFlags
	ResourceBuilderFlags *genericclioptions.ResourceBuilderFlags
	namespace            string
	client               *kubernetes.Clientset

	local      bool
	experiment string

	outputFormat string
	overrides    []string
}

func newOptions(streams genericclioptions.IOStreams) *Options {
	rbFlags := &genericclioptions.ResourceBuilderFlags{}
	rbFlags.WithAllNamespaces(false)

	return &Options{
		Streams:              streams,
		ConfigFlags:          genericclioptions.NewConfigFlags(true),
		ResourceBuilderFlags: rbFlags,
	}
}
