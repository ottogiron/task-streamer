# Task Streamer

Streams tasks to different type of processors.


Example

```go
package main

import (
    "github.com/ottogiron/taskstreamer"
    "github.com/ottogiron/scripttaskrunner"
  )

rabbitStream := NewRabbitStream(&taskstreamer.RabbitConfig{
    URL: 'amqp://guest:guest@localhost:5672/'
  })


scriptProcessor := processor.NewScriptProcessor("node", "script.js", "arg1", "arg2")

taskStreamer := NewTaskStreamer(rabbitStream)
taskStreamer.Register(scriptProcessor)
taskStreamer.Start()

taskStreamer.OnTaskDone(func (payload taskstreamer.Payload) {
    //do something
  })  

taskStreamer.OnError(func (err error){
    //do something
  })

taskStreamer.OnDone(func () {
  //do something
})
```
