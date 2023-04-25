//go:build windows

// Copyright 2023 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Subprocess go package used to Kills process and all child processes it
// spawned in Linux or Windows.
package subprocess

import (
	"syscall"
	"unsafe"
)

// KillProcessTree Kills the process and all its children in Windows
func KillProcessTree(pid int) (err error) {

	// Open a handle to the process with PROCESS_TERMINATE access
	handle, err := syscall.OpenProcess(syscall.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		return
	}
	defer syscall.CloseHandle(handle)

	// Get the list of child process IDs
	pids, err := getProcessChildren(pid)
	if err != nil {
		return
	}

	// Kill the child processes first
	for _, childPid := range pids {
		if err = killProcessTree(childPid); err != nil {
			return
		}
	}

	// Kill the process
	if err = syscall.TerminateProcess(handle, 0); err != nil {
		return
	}

	return
}

// getProcessChildren gets the list of child process IDs
func getProcessChildren(pid int) (pids []int, err error) {

	// Create a snapshot of the process list
	snapshot, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return
	}
	defer syscall.CloseHandle(snapshot)

	// Get the first process in the list
	var procEntry syscall.ProcessEntry32
	procEntry.Size = uint32(unsafe.Sizeof(procEntry))
	err = syscall.Process32First(snapshot, &procEntry)
	if err != nil {
		return
	}

	// Find the parent process and its children
	for {
		if procEntry.ProcessID == uint32(pid) {
			// Found the parent process, add its children to the list
			for {
				err := syscall.Process32Next(snapshot, &procEntry)
				if err != nil {
					break
				}
				if procEntry.ParentProcessID == uint32(pid) {
					pids = append(pids, int(procEntry.ProcessID))
				}
			}
			break
		}
		err = syscall.Process32Next(snapshot, &procEntry)
		if err != nil {
			return
		}
	}

	return
}
