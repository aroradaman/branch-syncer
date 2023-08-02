package repository

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"k8s.io/klog/v2"
	"k8s.io/utils/exec"
)

func getTempDir() string {
	return filepath.Join(os.TempDir(), uuid.New().String())
}

// runCommand runs commands
func runCommand(command string) (string, string, error) {
	stdOut := bytes.Buffer{}
	stdErr := bytes.Buffer{}

	runner := exec.New()

	klog.InfoS("running command", "command", command)

	cmd := runner.Command("/bin/bash", "-c", command)
	cmd.SetStdout(&stdOut)
	cmd.SetStderr(&stdErr)

	if err := cmd.Run(); err != nil {
		return "", "", err
	}
	return stdOut.String(), stdErr.String(), nil
}
