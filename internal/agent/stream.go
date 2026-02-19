package agent

import "time"

type StreamHandler func(chunk string) error

type Stream struct {
	ch chan string
}

func NewStream() *Stream {
	return &Stream{ch: make(chan string, 10)}
}

func (s *Stream) Send(chunk string) {
	select {
	case s.ch <- chunk:
	default:
	}
}

func (s *Stream) Close() {
	close(s.ch)
}

func (s *Stream) Handler(fn StreamHandler) {
	for chunk := range s.ch {
		if err := fn(chunk); err != nil {
			return
		}
	}
}

type StreamingLLM struct {
	llm LLM
}

func NewStreamingLLM(llm LLM) *StreamingLLM {
	return &StreamingLLM{llm: llm}
}

func (s *StreamingLLM) Generate(prompt string, fn StreamHandler) error {
	result, err := s.llm.Generate(prompt)
	if err != nil {
		return err
	}
	
	for i := 0; i < len(result); i += 5 {
		end := i + 5
		if end > len(result) {
			end = len(result)
		}
		fn(result[i:end])
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}
