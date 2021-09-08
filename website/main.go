package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) { /* w - что-то можно написать на странице, r - что-то прочесть */
	bob := User{
		name:      "Bob",
		age:       33,
		money:     1000,
		avgGrades: 4.2,
		happiness: 0.9,
	}
	fmt.Fprintf(w, bob.getAllInfo())
	bob.setNewName("Alex")
}

func contactsPage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "contacts page")
}

func handleReq() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contactsPage) // нужно ставить / в конце, чтобы браузер переводил на нужную страницу
	http.ListenAndServe(":8080", nil)
}

type User struct { // это структура, что-то типо класса, наследуясь можно создавать объекты
	name string
	age uint16
	money int16
	avgGrades, happiness float64
}

func (u *User) getAllInfo() string {  /*u - юзер тип его объект, нужен для указания иныф в строке*/
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func main() {
	handleReq()
}

