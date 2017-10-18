package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/chikamim/zipconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage:")
		os.Exit(1)
	}
	in := os.Args[1]
	re := regexp.MustCompile("^[0-9]+$")
	z, err := zipconv.New("zip.gob")
	if err != nil {
		panic(err)
	}
	var ret string
	if re.MatchString(in) {
		ret = z.ZipToAddress(in)
	} else {
		ret = z.AddressToZip(in)
	}

	fmt.Println(ret)
}
