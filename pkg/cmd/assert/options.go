package run

import (
	"time"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

const (
	Completed    = "completed"
	NoFailure    = "nofailure"
	SLOs         = "slos"
	SlosByPrefix = "slosby"
)

type Options struct {
	Streams              genericclioptions.IOStreams
	ConfigFlags          *genericclioptions.ConfigFlags
	ResourceBuilderFlags *genericclioptions.ResourceBuilderFlags
	namespace            string
	client               *kubernetes.Clientset

	local      bool
	experiment string

	conds   []string
	timeout time.Duration
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
