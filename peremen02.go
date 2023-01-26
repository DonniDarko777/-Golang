package main

import "fmt"

//глобальная переменная которая идет до перавой функции

var a, b, c int

//локальная переменная
func main() {
	//a, b, c = 1, 2, 3 // если присвоить значения то они будут как в первой локальной так и во второй
	a, b, c := 1, 2, 3 // несколько переменных в ряд с присвоение значений
	//a, b = b, a        // переназначили значение переменной
	a, _, c = 1, 6, 3 // можно пропускать значение которое меняется _

	fmt.Println(a, b, c)
	print()
}

//локальная переменная
func print() {
	a, b, c := 9, 5, 7
	fmt.Println(a, b, c)
	jojo() //что бы перешел в следующую лок переменную нужно написать название функции слудующего столба

}

//локальная переменная
func jojo() {
	a, b, c := 4, 6, 8
	fmt.Println(a, b, c)
	// если написать print в конце этой локальной переменной то будет писать вечно эти цифры
}