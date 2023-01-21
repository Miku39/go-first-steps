package main

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName, LastName string
}

type Employee struct {
	Account
	Credits float64
}

func (a *Account) ChangeName(newFirstName, newLastName string) {
	a.FirstName = newFirstName
	a.LastName = newLastName
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
}

func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits += amount
		return e.Credits, nil
	}
	return 0.0, errors.New("Invalid credit amount.")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("You can't remove more credits than the account has.")
	}
	return 0.0, errors.New("You can't remove negative numbers.")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func main() {
	miku, _ := CreateEmployee("Miku", "Sano", 500)

	// 名前の出力
	fmt.Println(miku)

	// 残高の確認
	fmt.Println(miku.CheckCredits())

	// クレジットの追加
	credits, err := miku.AddCredits(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	// クレジットの削除
	_, err = miku.RemoveCredits(2500)
	if err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	}

	// 名前の変更
	miku.ChangeName("Mmmm", "Ssss")

	// 名前の出力
	fmt.Println(miku)
}
