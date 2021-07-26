package main

import (
	"fmt"
	calc "modularity-demo/calculator"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	color.Red(strconv.Itoa(calc.Add(10, 20)))
	color.Green(strconv.Itoa(calc.Subtract(10, 20)))
	color.Yellow(fmt.Sprintf("Invocation count = %d\n", calc.GetInvocationCount()))
}
