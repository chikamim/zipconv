package zipconv_test

import (
	"fmt"

	"github.com/chikamim/zipconv"
)

func ExampleAddressToZip() {
	z, _ := zipconv.New("zip.gob")
	fmt.Println(z.AddressToZip("北海道札幌市中央区大通東"))
	// Output:
	// 0600041
}

func ExampleZipToAddress() {
	z, _ := zipconv.New("zip.gob")
	fmt.Println(z.ZipToAddress("1000014"))
	fmt.Println(z.ZipToAddress("6511102"))
	fmt.Println(z.ZipToAddress("3300081"))
	// Output:
	// 東京都千代田区永田町
	// 兵庫県神戸市北区山田町下谷上
	// 埼玉県さいたま市中央区新都心
}
