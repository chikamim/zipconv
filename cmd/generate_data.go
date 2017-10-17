package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"io"
	"io/ioutil"
	"os"

	"github.com/chikamim/zipconv"
)

func main() {
	csvPath := os.Args[1]
	outPath := "data.gob"
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
		//if r[8] != "以下に掲載がない場合" {
		z.Town = r[8]
		codes = append(codes, z)
	}
	return codes, nil
}
