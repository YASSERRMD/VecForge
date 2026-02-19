package agent

type Chain struct {
	steps []Step
}

type Step interface {
	Execute(input interface{}) (interface{}, error)
	Name() string
}

func NewChain() *Chain {
	return &Chain{steps: make([]Step, 0)}
}

func (c *Chain) Add(step Step) *Chain {
	c.steps = append(c.steps, step)
	return c
}

func (c *Chain) Execute(input interface{}) (interface{}, error) {
	current := input
	for _, step := range c.steps {
		result, err := step.Execute(current)
		if err != nil {
			return nil, err
		}
		current = result
	}
	return current, nil
}

type QueryRewriteStep struct {
	rewriter *QueryRewriter
}

func NewQueryRewriteStep(r *QueryRewriter) *QueryRewriteStep {
	return &QueryRewriteStep{rewriter: r}
}

func (s *QueryRewriteStep) Execute(input interface{}) (interface{}, error) {
	query, ok := input.(string)
	if !ok {
		return nil, ErrInvalidInput
	}
	return s.rewriter.Rewrite(query), nil
}

func (s *QueryRewriteStep) Name() string { return "query_rewrite" }

var ErrInvalidInput = &ChainError{msg: "invalid input"}

type ChainError struct{ msg string }

func (e *ChainError) Error() string { return e.msg }
