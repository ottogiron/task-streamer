package taskstreamer

//TaskStreamer streams data from different sources
type TaskStreamer interface {
	Start() error
}

//StreamAdapter for differnt types of streams
type StreamAdapter interface {
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
	return &taskStreamer{}
}

type taskStreamer struct {
	taskProcessor TaskProcessor
}

func (ts *taskStreamer) Start() error {
	return nil
}
