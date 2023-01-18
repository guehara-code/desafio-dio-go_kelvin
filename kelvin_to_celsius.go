package main //pacore principal//

import "fmt"

const t_kelvin = 300.0

func main() {
	var t_celsius = t_kelvin - 273.0
	fmt.Printf("Temperatura em Celsius = %g\n", t_celsius)
}
