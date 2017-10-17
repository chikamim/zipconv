package zipconv

type Zip struct {
	ZipCode    string
	Prefecture string
	City       string
	Town       string
}

func (z *Zip) Address() string {
	return z.Prefecture + z.City + z.Town
}
