package service

func mapCollection[F any, T any](c []F, m func(F) *T) []*T {
	var r []*T
	for _, v := range c {
		r = append(r, m(v))
	}
	return r
}
