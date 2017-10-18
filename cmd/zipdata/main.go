package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/chikamim/zipconv"
)

func main() {
	csvPath := "KEN_ALL.utf8.CSV"
	outPath := "zip.gob"
	codes, err := parseCSV(csvPath)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = gob.NewEncoder(buf).Encode(codes)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(outPath, buf.Bytes(), os.ModePerm)
}

func parseCSV(path string) (codes []zipconv.Zip, err error) {
	f, err := os.Open(path)
	if err != nil {
		return codes, err
	}

	r := csv.NewReader(f)
	rep := regexp.MustCompile(`（.+`)
	for {
		r, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return codes, err
		}
		// 01101,"064  ","0640941","ﾎｯｶｲﾄﾞｳ","ｻｯﾎﾟﾛｼﾁｭｳｵｳｸ","ｱｻﾋｶﾞｵｶ","北海道","札幌市中央区","旭ケ丘",0,0,1,0,0,0
		z := zipconv.Zip{}
		z.ZipCode = r[2]
		z.Prefecture = r[6]
		z.City = r[7]
		z.Town = rep.ReplaceAllString(r[8], "")
		if strings.Contains(z.Town, "以下に掲載がない場合") || strings.Contains(z.Town, "除く") || strings.Contains(z.Town, "以外") {
			z.Town = ""
		}
		codes = append(codes, z)
	}
	return codes, nil
}
