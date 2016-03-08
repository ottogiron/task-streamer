package taskstreamer

//TaskStreamer streams data from different sources
type TaskStreamer interface {
	Register(processor *TaskProcessor)
	Start() error
}

//Stream for differnt types of streams
type Stream interface {
	Open() error
	Close() error
	Next() (*Task, error)
	HasNext() bool
}

//TaskProcessor defines the interface for a task processor
type TaskProcessor interface {
	Process(task Task) error
}

//Payload for a task
type Payload map[string]interface{}

//Task to be processed
type Task struct {
	Payload Payload
}

//NewTaskStreamer returns a new instance of taskStreamer
func NewTaskStreamer() TaskStreamer {
	processors := []*TaskProcessor{}
	return &taskStreamer{processors}
}

type taskStreamer struct {
	processors []*TaskProcessor
}

func (ts *taskStreamer) Start() error {
	return nil
}

func (ts *taskStreamer) Register(taskProcessor *TaskProcessor) {
	ts.processors = append(ts.processors, taskProcessor)
}
