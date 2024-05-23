package format

import "fmt"

func Result(expr string, res float64) string {
	return fmt.Sprintf("CALCULATION SUCCESS: %s = %.2f", expr, res)

}
