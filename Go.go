package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	switch a {
	case 61:
		{
			fmt.Printf("Brasilia\n")
		}
	case 71:
		{
			fmt.Printf("Salvador\n")
		}
	case 11:
		{
			fmt.Printf("Sao Paulo\n")
		}
	case 21:
		{
			fmt.Printf("Rio de Janeiro\n")
		}
	case 32:
		{
			fmt.Printf("Juiz de Fora\n")
		}
	case 19:
		{
			fmt.Printf("Campinas\n")
		}
	case 27:
		{
			fmt.Printf("Vitoria\n")
		}
	case 31:
		{
			fmt.Printf("Belo Horizonte\n")
		}
	default:
		{
			fmt.Printf("DDD nao cadastrado\n")
		}
	}
}
