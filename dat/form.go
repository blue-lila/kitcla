package dat

type FormBag struct {
	Invalidations []*Invalidation
}

type Invalidation struct {
	Field   string
	Message string
}
