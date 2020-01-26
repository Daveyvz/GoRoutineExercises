//Package acdc asks if you're prepared to calculate.
package acdc

//Sum adds an unlimited amount of int values together and returns the total.
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
