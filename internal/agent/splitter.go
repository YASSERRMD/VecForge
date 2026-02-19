package agent

import "strings"

type Splitter interface {
	Split(doc Document) []Document
}

type TextSplitter struct {
	chunkSize int
	overlap  int
}

func NewTextSplitter(size, overlap int) *TextSplitter {
	return &TextSplitter{chunkSize: size, overlap: overlap}
}

func (s *TextSplitter) Split(doc Document) []Document {
	text := doc.Content
	if len(text) <= s.chunkSize {
		return []Document{doc}
	}
	
	var chunks []Document
	for i := 0; i < len(text); i += s.chunkSize - s.overlap {
		end := i + s.chunkSize
		if end > len(text) {
			end = len(text)
		}
		
		chunk := doc
		chunk.Content = text[i:end]
		chunks = append(chunks, chunk)
		
		if end >= len(text) {
			break
		}
	}
	return chunks
}

type SentenceSplitter struct{}

func NewSentenceSplitter() *SentenceSplitter { return &SentenceSplitter{} }

func (s *SentenceSplitter) Split(doc Document) []Document {
	sentences := strings.Split(doc.Content, ".")
	var docs []Document
	for _, sent := range sentences {
		if strings.TrimSpace(sent) != "" {
			docs = append(docs, Document{Content: sent, Meta: doc.Meta})
		}
	}
	return docs
}
