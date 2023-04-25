# Subprocess

This go package used to Kills process and all child processes it spawned in Linux or Windows.

[![GoDoc](https://godoc.org/github.com/kirill-scherba/subprocess?status.svg)](https://godoc.org/github.com/kirill-scherba/subprocess/)
[![Go Report Card](https://goreportcard.com/badge/github.com/kirill-scherba/subprocess)](https://goreportcard.com/report/github.com/kirill-scherba/subprocess)

The main idea behind creating this package is that the Windows code part is
very large and it is not practical to copy this code to every project.

## Usage example

```go
package main

import "kirill-scherba/subprocess" 

func main() {

    // Create some process, f.e. run exec.Run() function which execute buch 
    // file with subprocesses
    // 
    // var cmd *exec.Cmd
    // var err error
    // cmd, err = exec.Run("executable_name", "parameter")
    // ...
    // Use the killProcessTree function when you want stop execution started 
    // process and all its child process.
    //

    // Kill the process and all its children
    err = subprocess.killProcessTree(cmd.Process.Pid)
    if err != nil {
        return
    }

    // ...

}
```

## Licence

[BSD](LICENSE)
