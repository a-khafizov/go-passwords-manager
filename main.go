package main

import (
	"app/password/account"
	"app/password/files"
	"app/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		fmt.Println("__Менеджер паролей__")
		variant := promtData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант: ",
		})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promtData([]string{"Введите логин:"})
	password := promtData([]string{"Введите пароль:"})
	url := promtData([]string{"Введите url:"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или LOGIN")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promtData([]string{"Введите url для поиска: "})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData([]string{"Введите url для удаления: "})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удаленно")
	} else {
		output.PrintError("Не найдено")
	}
}

func promtData[T any](promt []T) string {
	for i, line := range promt {
		if i == len(promt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
