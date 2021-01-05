package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

var laberinto [15][15]string
var x, y, cont, pasos, k int
var optimo [2]int

func main() {
	/*Limpia la pantalla*/
	c := exec.Command("clear")
	c = exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Print("\tProyecto Final de Programacion\n\n")
	fmt.Print("\t  ┏┓     ┏┓           ┏┓ \n")
	fmt.Print("\t  ┃┃     ┃┃          ┏┛┗┓\n")
	fmt.Print("\t  ┃┃  ┏━━┫┗━┳━━┳━┳┳━ ┗┓┏╋━━┓\n")
	fmt.Print("\t  ┃┃ ┏┫┏┓┃┏┓┃┃━┫┏╋┫┏┓┓┃┃┃┏┓┃\n")
	fmt.Print("\t  ┃┗━┛┃┏┓┃┗┛┃┃━┫┃┃┃┃┃┃┃┗┫┗┛┃\n")
	fmt.Print("\t  ┗━━━┻┛┗┻━━┻━━┻┛┗┻┛┗┛┗━┻━━┛\n\n")

	//Condiciona a presionar "enter para continuar y limpia la pantall"
	fmt.Print("\tPresiona 'Enter' para comenzar...\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	c = exec.Command("clear")
	c.Stdout = os.Stdout

	color.HiGreen("\nInicio\n")
	estructura()
	diagrama()
	Encuentra_E()
	camina(laberinto, x, y)
}

//Se llena todo el arreglo con el simbolo "*"
func estructura() {
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			laberinto[i][j] = "*"
		}
	}
}

//Aqui se define la entrada,salida y el camino del laberinto
func diagrama() {
	laberinto[1][1] = "E"
	laberinto[12][13] = "S"
	laberinto[1][2] = " "
	laberinto[1][3] = " "
	laberinto[1][4] = " "
	laberinto[1][6] = " "
	laberinto[1][7] = " "
	laberinto[1][8] = " "
	laberinto[1][10] = " "
	laberinto[1][11] = " "
	laberinto[1][12] = " "

	laberinto[2][4] = " "
	laberinto[2][6] = " "
	laberinto[2][8] = " "
	laberinto[2][10] = " "
	laberinto[2][12] = " "
	laberinto[3][1] = " "
	laberinto[3][2] = " "
	laberinto[3][4] = " "
	laberinto[3][6] = " "
	laberinto[3][8] = " "
	laberinto[3][10] = " "
	laberinto[3][12] = " "

	laberinto[4][1] = " "
	laberinto[4][4] = " "
	laberinto[4][6] = " "
	laberinto[4][8] = " "
	laberinto[4][10] = " "
	laberinto[4][12] = " "

	laberinto[5][1] = " "
	laberinto[5][4] = " "
	laberinto[5][6] = " "
	laberinto[5][8] = " "
	laberinto[5][10] = " "
	laberinto[5][12] = " "

	laberinto[6][1] = " "
	laberinto[6][2] = " "
	laberinto[6][3] = " "
	laberinto[6][4] = " "
	laberinto[6][6] = " "
	laberinto[6][8] = " "
	laberinto[6][10] = " "
	laberinto[6][12] = " "

	laberinto[7][2] = " "
	laberinto[7][6] = " "
	laberinto[7][8] = " "
	laberinto[7][10] = " "
	laberinto[7][12] = " "

	laberinto[8][2] = " "
	laberinto[8][3] = " "
	laberinto[8][4] = " "
	laberinto[8][5] = " "
	laberinto[8][6] = " "
	laberinto[8][8] = " "
	laberinto[8][10] = " "
	laberinto[8][12] = " "

	laberinto[9][5] = " "
	laberinto[9][8] = " "
	laberinto[9][10] = " "
	laberinto[9][12] = " "

	laberinto[10][5] = " "
	laberinto[10][8] = " "
	laberinto[10][10] = " "
	laberinto[10][12] = " "
	laberinto[10][13] = " "

	laberinto[11][5] = " "
	laberinto[11][8] = " "
	laberinto[11][9] = " "
	laberinto[11][10] = " "
	laberinto[11][13] = " "

	laberinto[12][5] = " "
	laberinto[12][11] = " "
	laberinto[12][12] = " "

	laberinto[13][5] = " "
	laberinto[13][6] = " "
	laberinto[13][7] = " "
	laberinto[13][8] = " "
	laberinto[13][9] = " "
	laberinto[13][10] = " "
	laberinto[13][11] = " "

}

//Imprime el laberinto ya estructurado
func impresion(laberinto [15][15]string) {
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			fmt.Printf(" %s ", laberinto[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Print("\n\n")
}

//Se va a buscar la entrada a partir de la letra "E"
//Si lo encuentra se guarda la posicion de "E" en dos variables utilizadas como comodin
func Encuentra_E() {

	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			if laberinto[i][j] == "E" {
				x = i
				y = j
			}
		}
	}
}

//Mediante condiciones se buscara la salida
func camina(laberinto [15][15]string, x, y int) {

	//Si en la posicion se encuentra un espacio lo sustituye por un punto
	//Asi se marcara el camino
	if laberinto[x][y] == " " {
		laberinto[x][y] = "."
		pasos++
	}

	//Valida si se encontro la salida e imprime la solucion
	if laberinto[x][y] == "S" {
		color.HiBlue("\nSolucion %d\n", cont+1)
		for i := 0; i < 15; i++ {
			for j := 0; j < 15; j++ {
				fmt.Printf("%s ", laberinto[i][j])
			}
			fmt.Printf("\n")
		}
		//Imprime el numero de pasos por solucion posible
		//Se utiliza un arreglo para guardar los pasos de cada solucion
		fmt.Printf("\tNumero de pasos: %d\n\n\n", pasos)
		optimo[k] = pasos
		cont++
		k++
		//Se evalua el camino con menos pasos para llegar a la salida
		if optimo[0] < optimo[1] {
			fmt.Printf("Camino optimo: %d pasos\n", optimo[0])
			fmt.Printf("Se encontraron %d caminos de salida \n", cont)
		}

		return
	}
	//Imprime los pasos dados
	impresion(laberinto)

	//Aqui se consigue que se logre visualizar el recorrido del camino
	duration := time.Second
	time.Sleep(duration)

	//Evalua la condicion y si es verdadera camina hacia la izquierda
	if (laberinto[x-1][y] == " " || laberinto[x-1][y] == "S") && laberinto[x-1][y] != "*" {
		camina(laberinto, x-1, y)
	}
	//Evalua la condicion y si es verdadera camina hacia hacia arriba
	if (laberinto[x][y-1] == " " || laberinto[x][y-1] == "S") && laberinto[x][y-1] != "*" {
		camina(laberinto, x, y-1)
	}
	//Evalua la condicion y si es verdadera camina hacia la derecha
	if (laberinto[x+1][y] == " " || laberinto[x+1][y] == "S") && laberinto[x+1][y] != "*" {
		camina(laberinto, x+1, y)
	}
	//Evalua la condicion y si es verdadera camina hacia abajo
	if (laberinto[x][y+1] == " " || laberinto[x][y+1] == "S") && laberinto[x][y+1] != "*" {
		camina(laberinto, x, y+1)
	}
	//Evalua la condicion y si es verdadera resta los pasos dados para no contabilizarlos
	if laberinto[x][y] == "*" || laberinto[x][y] == "." {
		pasos--
		return
	}

	return
}
