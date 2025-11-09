# Task Scheduler

> Task scheduler (Cron, batch jobs)

## Installation
```bash
go get github.com/modsynth/task-scheduler
```

## Usage
```go
package main

import (
    "github.com/modsynth/task-scheduler"
    "log"
)

func main() {
    s := scheduler.New()
    
    // Run every day at midnight
    s.AddJob("0 0 * * *", func() {
        log.Println("Daily task executed")
    })
    
    // Run every 5 minutes
    s.AddJob("*/5 * * * *", func() {
        log.Println("5-minute task executed")
    })
    
    s.Start()
    defer s.Stop()
    
    select {} // Keep running
}
```

## Cron Format
```
* * * * *
│ │ │ │ │
│ │ │ │ └─ Day of week (0-6)
│ │ │ └─── Month (1-12)
│ │ └───── Day of month (1-31)
│ └─────── Hour (0-23)
└───────── Minute (0-59)
```

## Version
v0.1.0

## License
MIT
