package stream

import "github.com/ottogiron/taskstreamer"

type rabbitStream struct {
}

//NewRabbitStream Creates a new rabbit tasks streamer
func NewRabbitStream() taskstreamer.Stream {
	return &rabbitStream{}
}

func (s *rabbitStream) Next() (*taskstreamer.Task, error) {
	return nil, nil
}

func (s *rabbitStream) Open() error {
	return nil
}

func (s *rabbitStream) Close() error {
	return nil
}

func (s *rabbitStream) HasNext() bool {
	return false
}
