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
	for _, zip := range c.Zipcodes {
		if strings.HasPrefix(zip.Address(), address) {
			return zip.Code
		}
	}
	return ""
}

func (c *Converter) ZipToAddress(code string) string {
	for _, zip := range c.Zipcodes {
		if zip.Code == code {
			return zip.Address()
		}
	}
	return ""
}
