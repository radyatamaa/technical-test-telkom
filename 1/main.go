package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	// LIST PECAHAN
	listPecahanRupiah := []float64{100000,50000,20000,10000,5000,2000,1000,500,200,100}
	//===============================

	//INPUT
	var input float64 = 145000
	// =======================================

	//OUTPUT
	output := map[string]int{}
	for i := 0; i < len(listPecahanRupiah); i++ {
		if input < 100 {
			input = 100
		}
		jml := math.Floor(input / listPecahanRupiah[i])
		input = input - (listPecahanRupiah[i] * jml)

		if jml != 0 {
			output[fmt.Sprint("RP. ", listPecahanRupiah[i])] = int(jml)
		}
		if input == 0 {
			break
		}
	}

	outputString ,_:= json.MarshalIndent(output, "", "    ")
	fmt.Println(string(outputString))
	//===================================
}
