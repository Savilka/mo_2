package main

import (
	"encoding/json"
	"os"
)

type Input struct {
	X0     []float32 `json:"x0"`
	Dx     float32   `json:"dx"`
	Eps    float32   `json:"eps"`
	Lambda float32   `json:"lambda"`
}

func searchInterval(x []float32, delta, lambda float32) []float32 {

}

func calculateFunction(x []float32, number int) float32 {
	switch number {
	case 1:
		return 100*(x[1]-x[0])*(x[1]-x[0]) + (1-x[0])*(1-x[0])
	case 2:

	case 3:

	}
	return 0
}

func search(dx float32, x0 []float32) ([]float32, int) {
	var flag int
	var x []float32
	copy(x, x0)
	for i := 0; i < len(x0); i++ {
		var cx []float32
		copy(cx, x)
		cx[i] += dx
		if calculateFunction(cx, 1) <= calculateFunction(x, 1) {
			x[i] += dx
			flag = 1
		} else {
			cx[i] -= 2 * dx
			if calculateFunction(cx, 1) <= calculateFunction(x, 1) {
				x[i] += dx
				flag = 1
			}
		}
	}
	if flag == 0 {
		return x0, flag
	}
	return x, flag
}

func main() {
	inputRaw, err := os.ReadFile("input.json")
	if err != nil {
		return
	}
	var input Input
	err = json.Unmarshal(inputRaw, &input)
	if err != nil {
		return
	}

	for input.Dx >= input.Eps {
		x, flag := search(input.Dx, input.X0)
		if flag == 0 {
			input.Dx /= 10
			continue
		}

	}

}
