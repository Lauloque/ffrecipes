package kitchen

type Cook struct {
	Pantry *Pantry
}

func (c Cook) Prepare(d Recipe) error {
	return d.Cook(c.Pantry)
}
