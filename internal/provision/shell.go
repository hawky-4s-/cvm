package provision

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

var defaultShellTimeout = 30 * time.Second

func RunCmd(ctx context.Context, path string, args []string, debug bool) (out string, err error) {
	var cmd *exec.Cmd

	if ctx == nil {
		// construct default timeout context
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), defaultShellTimeout)
		defer cancel()

		cmd = exec.CommandContext(ctx, path, args...)
	} else {
		cmd = exec.CommandContext(ctx, path, args...)
	}

	var b []byte
	b, err = cmd.CombinedOutput()
	out = string(b)

	if debug {
		fmt.Println(strings.Join(cmd.Args[:], " "))

		if err != nil {
			fmt.Println("RunCMD ERROR")
			fmt.Println(out)
		}
	}

	return
}
