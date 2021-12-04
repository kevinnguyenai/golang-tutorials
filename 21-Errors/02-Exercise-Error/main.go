package main 

import (

	"fmt"
	"math"
	"bytes"
	"strconv"
)

type ErrNegativeSqrt float64


func (e ErrNegativeSqrt) Error() string {
	var result bytes.Buffer
	result.WriteString("Cannot Sqrt Negative Number: ")
	result.WriteString(strconv.FormatFloat(float64(e),'E',-1,64))
	return result.String()
}

func Sqrt(x float64) (float64, error) {
	if(x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x*x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}