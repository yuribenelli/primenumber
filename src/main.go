package main

import "fmt"

func main() {
	//prendo input utente
	var input int
	var isPrime bool

	fmt.Println("Inserisci numero: ")
	fmt.Scan(&input)
	if input <= 2 {
		fmt.Println(input)
	}
	//creo slice con capacity max = input utente
	sli := make([]int, 0, input)
	//ciclo esterno (test per input e input - j)
	inputSave := input
	for j := input; j > 1; j-- {
		for i := input - 1; i >= 2; i-- { // test interno per divisori inferiori di input
			x := input % i
			if x == 0 { //se divisibile allora input non è un numero primo
				isPrime = false
				break
			} else if x != 0 {
				isPrime = true
			}
		}
		if isPrime {
			sli = append(sli, input)
		}
		input--
	}
	fmt.Println(sli)
	fmt.Printf("Ci sono %v numeri primi compresi fra 0 e %v \n", len(sli), inputSave)
}
