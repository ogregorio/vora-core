package io

import (
	"strconv"
	"math"

	files "./files"
)

func Filestest() {

	for i:=0; i < math.MaxInt64; i++ {
		files.Write(TESTE_PATH, strconv(i))
		
}