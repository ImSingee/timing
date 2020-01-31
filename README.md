# Timing

Go package for the [Timing](https://timingapp.com/) API.

## About

This package enables your Go application to send requests to the Timing service through the Timing REST API. 

Note that Timing has many APIs available, but currently this package only supports:

+ Projects
    + [ ] Return the complete project hierarchy.
    + [ ] Return a list containing all projects.
    + [ ] Create a new project.
    + [ ] Display the specified project.
    + [ ] Update the specified project.
    + [ ] Delete the specified project and all of its children.
 + Tasks
    + [ ] Start a new task.
    + [ ] Stop the currently running task.
    + [x] Return a list of tasks.
    + [ ] Create a new task.
    + [ ] Display the specified task.
    + [ ] Update the specified task.
    + [ ] Delete the specified task.

## Using the Package

Registering a Timing account and obtaining an API Key for using this library are not covered in this README. 

[Timing API Keys](https://web.timingapp.com/integrations/tokens) are required.

```go
package main

import (
    "github.com/ImSingee/timing"
    "log"
)

func init() {
    timing.Init("put API Key here")
}


func main() {
    r, e := timing.TimeEntries(timing.TimeEntriesRequest{
        IsRunning: true,
    })
    
    if e != nil {
        log.Print(e)
    }
    
    log.Print(r)
}
```

