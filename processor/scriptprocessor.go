package processor

import "github.com/ottogiron/taskstreamer"

type scriptProcessor struct {
}

//NewScriptProcessor returns a new instance of a script processor
func NewScriptProcessor(command string, args ...string) taskstreamer.TaskProcessor {
	return &scriptProcessor{}
}

func (sp *scriptProcessor) Process(task taskstreamer.Task) error {
	return nil
}
