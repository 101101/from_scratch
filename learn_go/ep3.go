package main

// go run gotut.go

import ("fmt")

/* ------- example 1 -------  
func add(x float64,y float64) float64 {
	return x+y
}

func main() {
	var num1 float64 = 5.6
	var num2 float64 = 9.5

	fmt.Println(add(num1,num2))
}
*/

/* ------- reduced syntax ------- 
func add(x,y float64) float64 {
	return x+y
}

func main() {
	var num1,num2 float64 = 5.6, 9.5

	fmt.Println(add(num1,num2))
}
*/


/* ------- shorthand ------- */
func add(x,y float64) float64 {
	return x+y
}

func main() {
	num1,num2 := 5.6, 9.5

	fmt.Println(add(num1,num2))
}




/* ------- string ------- */
func multiple(a,b string) (string,string) {
	return a,b
}

func main() {
	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1,w2))
}



/* ------- convert type ------- */
func multiple(a,b string) (string,string) {
	return a,b
}

func main() {
	var a int = 62
	var b float64 = float64(a)

	x := a 	// will be type int
}