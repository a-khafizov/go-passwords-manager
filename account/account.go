package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOIPASDFGHJKLZXCVBNM1234567890-*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc Account) Output() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.Password = string(res)
}

// 1. если логина нет - ошибка
// 2. если нет пароля - генерируем
func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Login:     login,
		Password:  password,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
