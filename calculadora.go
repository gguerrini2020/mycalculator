package mycalculator

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type calc struct{}

func (calc)operate(entrada string, operador string)int {
	entradaLimpia := strings.Split(entrada , operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador , " El operador no esta soportado")	
		return 0
	}

}

func parsear(entrada string) int {
	operador , _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada()string  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return  scanner.Text()
}

