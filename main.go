package main

import (
	"app/password/account"
	"app/password/files"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		fmt.Println("__Менеджер паролей__")
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func createAccount(vault *account.VaultWithDb) {
	login := promtData("Введите логин:")
	password := promtData("Введите пароль:")
	url := promtData("Введите url:")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		color.Red("Неверный формат URL или LOGIN")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promtData("Введите url для поиска: ")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData("Введите url для удаления: ")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удаленно")
	} else {
		color.Red("Не найдено")
	}
}

func promtData(promt string) string {
	fmt.Print(promt + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
