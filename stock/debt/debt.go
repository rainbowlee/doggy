package debt

import (
	//"flag"
	"fmt"
)

func Test()  {
	fmt.Print("test")

	e := NewDebt("hengrui", 5.47, "hengrui1", 5.02)
	e.ToString()
	e.Estimate(5.31, 0.729)
}