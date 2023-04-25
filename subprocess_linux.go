//go:build !windows

// Copyright 2023 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Subprocess go package used to Kills process and all child processes it
// spawned in Linux or Windows.
package subprocess

import "syscall"

// Kill the process and all its children in Linux
func killProcessTree(pid int) (err error) {
	pgid, err := syscall.Getpgid(pid)
	if err != nil {
		return
	}
	syscall.Kill(-pgid, syscall.SIGKILL)
	return
}
