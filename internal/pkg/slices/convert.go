package slices

func Convert[In any, Out any](f func(In) Out, in []In) []Out {
	out := make([]Out, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}
