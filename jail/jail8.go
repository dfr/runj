package jail

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// CreateJail wraps the jail(8) command to create a jail
func CreateJail(ctx context.Context, confPath string) (Jail, error) {
	cmd := exec.CommandContext(ctx, "jail", "-icf", confPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		return nil, err
	}
	id, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return nil, err
	}
	return FromID(ID(id)), nil
}

// DestroyJail wraps the jail(8) command to destroy a jail
func DestroyJail(ctx context.Context, confPath, jail string) error {
	cmd := exec.CommandContext(ctx, "jail", "-f", confPath, "-r", jail)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
	}
	return err
}
