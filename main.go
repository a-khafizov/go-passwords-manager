package main

import (
	"app/password/account"
	"app/password/encrypter"
	"app/password/files"
	"app/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": FindAccountsByUrl,
	"3": FindAccountsByLogin,
	"4": deleteAccount,
}

func main() {
	fmt.Println("__Менеджер паролей__")
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Ошибка чтения env")
	}

	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
Menu:
	for {
		variant := promtData(
			"1. Создать аккаунт",
			"2. Найти аккаунт по url",
			"3. Найти аккаунт по login",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите вариант",
		)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или LOGIN")
		return
	}
	vault.AddAccount(*myAccount)
}

func FindAccountsByUrl(vault *account.VaultWithDb) {
	url := promtData("Введите url для поиска: ")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputReult(&accounts)
}

func FindAccountsByLogin(vault *account.VaultWithDb) {
	login := promtData("Введите login для поиска: ")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputReult(&accounts)
}

func outputReult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
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

func promtData[T any](promt ...T) string {
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
