// +build windows

/*
Copyright (c) 2019 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package plugin

import (
	"os/exec"
	"syscall"
)

func pluginCmd(cmd string) *exec.Cmd {
	pluginCmd := exec.Command(cmd)
	pluginCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return pluginCmd
}
