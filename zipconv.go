package zipconv

import (
	"bufio"
	"encoding/gob"
	"os"
	"strings"
)

type Converter struct {
	Zipcodes []Zip
}

func New(path string) (converter Converter, err error) {
	f, err := os.Open(path)
	if err != nil {
		return converter, err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	codes := []Zip{}
	err = gob.NewDecoder(buf).Decode(&codes)
	if err != nil {
		return converter, err
	}
	converter.Zipcodes = codes
	return converter, nil
}

func (c *Converter) AddressToZip(address string) string {
	for _, code := range c.Zipcodes {
		if strings.HasPrefix(code.Address(), address) {
			return code.ZipCode
		}
	}
	return ""
}

func (c *Converter) ZipToAddress(zip string) string {
	for _, code := range c.Zipcodes {
		if code.ZipCode == zip {
			return code.Address()
		}
	}
	return ""
}
