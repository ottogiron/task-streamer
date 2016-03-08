package streamer

import "github.com/ottogiron/taskstreamer"

type rabbitStreamer struct {
}

//NewRabbitStreamer Creates a new rabbit tasks streamer
func NewRabbitStreamer() taskstreamer.StreamAdapter {
	return &rabbitStreamer{}
}

func (s *rabbitStreamer) Next() (*taskstreamer.Task, error) {
	return nil, nil
}

func (s *rabbitStreamer) Open() error {
	return nil
}

func (s *rabbitStreamer) Close() error {
	return nil
}

func (s *rabbitStreamer) HasNext() bool {
	return false
}
