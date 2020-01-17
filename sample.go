package redefinitions

type Sample struct {
	Value int64
}

type SamplePair struct {
	Sample1 *Sample
	Sample2 *Sample
}

func (s *SamplePair) Equal(o *SamplePair) bool {
	return s.Sample1.Value == o.Sample1.Value && s.Sample2.Value == o.Sample2.Value
}

func (s *SamplePair) Swap() *SamplePair {
	return &SamplePair{
		Sample1: s.Sample2,
		Sample2: s.Sample1,
	}
}
