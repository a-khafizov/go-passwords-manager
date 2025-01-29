package main

import (
	"app/password/account"
	"fmt"
)

func main() {
Menu:
	for {
		fmt.Println("__Менеджер паролей__")
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант: ")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scanln(&variant)
	return variant
}

func createAccount() {
	login := promtData("Введите логин:")
	password := promtData("Введите пароль:")
	url := promtData("Введите url:")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func findAccount() {

}

func deleteAccount() {

}

func promtData(promt string) string {
	fmt.Print(promt + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
