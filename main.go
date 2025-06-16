package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n int8

func main() {
	rand.Seed(time.Now().UnixNano())
	lowLetters := "abcdefghijklmnopqrstuvwxyz"
	upLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChars := "!@#$%^&*()-_=+,.?/:;{}[]~"
	digits := "0123456789"
	var all, parol string
	for i := 0; i <= 10; i++ {
		all += string(lowLetters[rand.Intn(len(lowLetters))])
	}
	for i := 0; i <= 10; i++ {
		all += string(upLetters[rand.Intn(len(upLetters))])
	}
	for i := 0; i <= 10; i++ {
		all += string(specialChars[rand.Intn(len(specialChars))])
	}
	for i := 0; i <= 10; i++ {
		all += string(digits[rand.Intn(len(digits))])
	}
	proverka()
	for i := 0; i <= int(n); i++ {
		if i == 0 {
			parol += string(all[rand.Intn(19)])
		} else {
			parol += string(all[rand.Intn(len(all))])
		}
	}
	fmt.Printf("Пароль: %s", parol)
}
func proverka() {
	for {
		fmt.Print("Выберите необходимую длину пароля: ")
		_, err := fmt.Scan(&n)
		if err == nil {
			if n < 8 {
				fmt.Println("Введена не корректная длина пароля.")
				var newn string
				fmt.Scanln(&newn)
			} else {
				break
			}
		} else {
			fmt.Println("Веденно неверное значение.")
			var newn string
			fmt.Scanln(&newn)
		}
	}
}
