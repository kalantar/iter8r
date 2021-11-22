package hub

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// "gotest.tools/assert"

func Test_ExecuteHubCommand_NoExperiment(t *testing.T) {
	cmd := NewCmd()

	// stderrBuffer := bytes.NewBufferString("")
	// stdoutBuffer := bytes.NewBufferString("")
	// cmd.SetOut(stdoutBuffer)
	// cmd.SetErr(stderrBuffer)

	err := cmd.Execute()
	assert.EqualError(t, err, ExperimentRequiredError)

	// stderr, err := ioutil.ReadAll(stderrBuffer)
	// assert.NilError(t, err)

	// assert.Equal(t, fmt.Sprintf("Error: %s\n", ExperimentRequiredError), string(stderr))
}

func Test_ExecuteHubCommand_InvalidExperimentDefaultHub(t *testing.T) {
	cmd := NewCmd()

	stderrBuffer := bytes.NewBufferString("")
	stdoutBuffer := bytes.NewBufferString("")
	cmd.SetOut(stdoutBuffer)
	cmd.SetErr(stderrBuffer)

	logBuffer := bytes.NewBufferString("")
	log.SetOutput(logBuffer)
	// defer func () {
	// 	log.SetOuptut*(os.Stderr)
	// }()

	cmd.SetArgs([]string{"-e", "invalid-experiment"})
	err := cmd.Execute()
	assert.EqualError(t, err, "exit status 1")

	t.Logf("log: '%s'\n", logBuffer.String())
	t.Logf("stdout: '%s'\n", stdoutBuffer.String())
	t.Logf("stderr: '%s'\n", stderrBuffer.String())
	// stdout, err := ioutil.ReadAll(stderrBuffer)
	// assert.NoError(t, err)

	// assert.Contains(t, string(stdout), "unable to get: github.com/iter8-tools/iter8.git/mkdocs/docs/hub/invalid-experiment")
	// assert.Equal(t, fmt.Sprintf("Error: %s\n", ExperimentRequiredError), string(stderr))
}

func Test_ExecuteHubCommand_ValidExperiment(t *testing.T) {
	cmd := NewCmd()
	os.Setenv("ITER8HUB", "")
	cmd.SetArgs([]string{"-e", "valid-experiment"})
	err := cmd.Execute()
	assert.NoError(t, err)

}

// func Test_ExecuteHubCommand_InValidExperimentInvalidHub(t *testing.T) {
// 	cmd := NewCmd()
// 	cmd.SetArgs([]string{"-e", "valid-experiment"})
// 	err := cmd.Execute()
// 	assert.EqualError(t, err, ExperimentRequiredError)
// }
