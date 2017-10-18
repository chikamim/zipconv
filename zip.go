package zipconv

type Zip struct {
	Code       string
	Prefecture string
	City       string
	Town       string
}

func (z *Zip) Address() string {
	return z.Prefecture + z.City + z.Town
}
