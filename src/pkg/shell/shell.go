package shell

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Run(name string, stdin io.Reader, stdout, stderr io.Writer, arg ...string) error {
	exitCode, err := RunWithExitCode(name, stdin, stdout, stderr, arg...)
	if exitCode != 0 && err == nil {
		err = fmt.Errorf("failed to invoke %s(1)", name)
	}

	if err != nil {
		return err
	}

	return nil
}

func RunWithExitCode(name string, stdin io.Reader, stdout, stderr io.Writer, arg ...string) (int, error) {
	if _, err := exec.LookPath(name); err != nil {
		return 1, fmt.Errorf("%s(1) not found", name)
	}

	if stderr == nil {
		stderr = os.Stderr
	}

	cmd := exec.Command(name, arg...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			exitCode := exitErr.ExitCode()
			return exitCode, nil
		}

		return 1, fmt.Errorf("failed to invoke %s(1)", name)
	}

	return 0, nil
}
