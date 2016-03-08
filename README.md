# Task Streamer

Streams tasks to different type of runners or processors.


Example

```go
package main

import (
    "github.com/ottogiron/taskstreamer"
    "github.com/ottogiron/scripttaskrunner"
  )

streamAdapter := NewRabbitStreamerAdapter(&taskstreamer.RabbitConfig{
    URL: 'amqp://guest:guest@localhost:5672/'
  })


scriptTaskRunner := scripttaskrunner.New("node", "script.js", "arg1", "arg2")

taskStreamer := NewTaskStreamer(streamer)
taskStreamer.register(scriptTaskRunner)
taskStreamer.start()

taskStreamer.onTaskDone(func (payload taskstreamer.Payload) {
    //something
  })

taskStreamer.onError(func (err error){
    //something
  })
```
