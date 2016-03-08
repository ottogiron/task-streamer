package processor

import "github.com/ottogiron/taskstreamer"

type command struct {
}

//NewCommandProcessor returns a new instance of a script processor
func NewCommandProcessor(cmd string, args ...string) taskstreamer.TaskProcessor {
	return &command{}
}

func (sp *command) Process(task taskstreamer.Task) error {
	return nil
}
