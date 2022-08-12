package main

import (
	"fmt"
	"os"
)

var code1 = []byte{1,5,1,7,10,0}  //add
var code2 = []byte{1,5,1,7,11,0}  //sub
var code3 = []byte{1,7,1,5,12,0}  //mul
var code4 = []byte{1,3,1,12,13,0} //quo
var code5 = []byte{1,5,1,7,14,0}  //rem
var code6 = []byte{1,7,1,5,15,0}  //and
var code7 = []byte{1,7,1,5,16,0}  //or
var code8 = []byte{1,7,1,5,17,0}  //xor

var pathbase="D:/go/src/golang.org/x/tools/go/ssa/interp/testdata/"
var bytecodes = []string{
	pathbase+"test1add.go.btc",
	pathbase+"test1sub.go.btc",
	pathbase+"test1mul.go.btc",
	pathbase+"test1quo.go.btc",
	pathbase+"test1rem.go.btc",
	pathbase+"test1and.go.btc",
	pathbase+"test1or.go.btc",
	pathbase+"test1xor.go.btc",
}


func main() {
	for _, bytecode := range bytecodes {
		interp(bytecode)
	}

}

func interp(bytecode string){
	var filePath=bytecode
	data := f_read(filePath)
	ret, _ := Run(data)
	fmt.Println(data)
	fmt.Println(ret[len(ret)-1])
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func f_write(f *os.File, code []byte) {
	_, err := f.Write(code)
	check(err)
}

func f_read(file_path string) []byte {
	dat, err := os.ReadFile(file_path)
	check(err)
	return dat
}