package zipconv_test

import (
	"fmt"

	"github.com/chikamim/zipconv"
)

func ExampleAddressToZip() {
	z, _ := zipconv.New("data.gob")
	code := z.AddressToZip("北海道札幌市中央区大通東")

	fmt.Println(code)
	// Output:
	// 0600041
}

func ExampleZipToAddress() {
	z, _ := zipconv.New("data.gob")
	code := z.ZipToAddress("0600041")

	fmt.Println(code)
	// Output:
	// 北海道札幌市中央区大通東
}
