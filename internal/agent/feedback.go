package agent

import "time"

type Feedback struct {
	Query    string
	ResultID string
	Rating   int
	Ts       time.Time
}

type FeedbackCollector struct {
	feedbacks []Feedback
}

func NewFeedbackCollector() *FeedbackCollector {
	return &FeedbackCollector{feedbacks: make([]Feedback, 0)}
}

func (fc *FeedbackCollector) Add(f Feedback) {
	f.Ts = time.Now()
	fc.feedbacks = append(fc.feedbacks, f)
}

func (fc *FeedbackCollector) GetAll() []Feedback {
	return fc.feedbacks
}

func (fc *FeedbackCollector) Positive() []Feedback {
	var pos []Feedback
	for _, f := range fc.feedbacks {
		if f.Rating > 3 {
			pos = append(pos, f)
		}
	}
	return pos
}

func (fc *FeedbackCollector) Negative() []Feedback {
	var neg []Feedback
	for _, f := range fc.feedbacks {
		if f.Rating <= 3 {
			neg = append(neg, f)
		}
	}
	return neg
}
