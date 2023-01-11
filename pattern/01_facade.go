package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера
на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Применяется когда нужно обращаться со сложной подсистемой, реализует простой интерфейс для доступа к
определенной функциональности подсистемы.
Плюсы:
-Изолирует клиента от сложной подсистемы
-Снижает сложность программы
-Можно вывести код зависимый от внешней системы в единое место
Минусы:
-Может стать "Божественным объектом" привязанными ко всем классам программы
*/

type Client struct {
	name string
	card *Card
}

type Card struct {
	balance  float64
	password int
}

type Bank struct {
	cards []Card
}

type ATMFacade struct {
	bank   *Bank
	client *Client
}

func newClient() *Client {
	return &Client{
		card: &Card{
			balance:  0,
			password: 123,
		},
		name: "Ilya",
	}
}
func (client *Client) Balance() float64 {
	return client.card.balance
}

func (client *Client) AddCash(cash float64) {
	client.card.balance += cash
}

func newATMFacade() *ATMFacade {
	return &ATMFacade{
		client: newClient(),
	}
}

func (atm *ATMFacade) WorkATMFacade() error {
	fmt.Printf("Choose what to use\n1 - Balance\n2 - Add Money\n3 - Take Money\n4 - Exit\n")
	var mode int
	for {
		fmt.Scanln(&mode)
		switch mode {
		case 1:
			fmt.Printf("Your balance is: %.2f\n", atm.client.Balance())
		case 2:
			fmt.Printf("Place money in the hole\n")
			var cash float64
			fmt.Scanln(&cash)
			atm.client.AddCash(cash)
		case 3:
			fmt.Printf("How many cash do you want to take?\n")
			var cash float64
			fmt.Scanln(&cash)
			for atm.client.Balance() < cash {
				fmt.Printf("You don't have enough money.\nYou have only %.2f on your balance\n", atm.client.Balance())
				fmt.Scanln(&cash)
			}
			atm.client.card.balance -= cash
		case 4:
			break
		default:
			return errors.New(fmt.Sprintf("You have chosen wrong mode\n"))
		}
		break
	}

	return nil
}
