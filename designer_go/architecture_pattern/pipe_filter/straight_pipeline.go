package pipefilter

type StraightPipeline struct {
	Name    string
	Filters []Filter
}

func NewStraightPipeline(name string, filter ...Filter) *StraightPipeline {
	return &StraightPipeline{
		name,
		filter,
	}
}

func (f *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
