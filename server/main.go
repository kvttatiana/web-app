package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Создать массив структур, Создать роут create, который принимает POST запросы
//По этому роуту мы должны брать данные, парсить их в структуру и класть в массив
//При удачных обстоятельствах фронту отправляем "Данные получены", если возникла ошибка, то "Данным пизда
// добавить роут /getusers с методом гет, который отправляет фронту массив структур
//роут /update с методом пут, для обновления передаём айди и новые данные. По айди ищем эл массива и обновляем данные
//Поменять структуру (добавить поле username и password)

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	ID       uint   `json:"id"`
}

var counter = 1
var array []User

func greet(w http.ResponseWriter, r *http.Request) {
	currentUser := User{}
	err := json.NewDecoder(r.Body).Decode(&currentUser)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Данным пизда"))
	} else {
		currentUser.ID = uint(counter)
		counter += 1
		fmt.Println(err)
		w.Write([]byte("Данные получены"))
		fmt.Println(currentUser)
		array = append(array, currentUser)

	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	jsonString, err := json.Marshal(array)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Не удалось отправить данные"))
	} else {
		w.Write(jsonString)
	}

} //роут /update с методом пут, для обновления передаём айди и новые данные. По айди ищем эл массива и обновляем данные
// роут /delete с методом делит
func upd(w http.ResponseWriter, r *http.Request) {
	currentUser := User{}
	err := json.NewDecoder(r.Body).Decode(&currentUser)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Пизда данным"))
		return
	}
	var cur = currentUser.ID
	if cur-1 >= uint(len(array)) {
		w.Write([]byte("Пользователя с таким ID не существует"))
		return
	}
	array[cur-1].Name = currentUser.Name
	array[cur-1].Surname = currentUser.Surname
	w.Write([]byte("Пользователь обновлён"))

}

func del(w http.ResponseWriter, r *http.Request) {
	currentUser := User{}
	err := json.NewDecoder(r.Body).Decode(&currentUser)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Пизда данным"))
		return
	}
	var cur = currentUser.ID
	if cur-1 >= uint(len(array)) {
		w.Write([]byte("Пользователя с таким ID не существует"))
		return
	}
	array = remove(array, currentUser.ID-1)
	w.Write([]byte("Пользователь удалён"))
}
func remove(array []User, s uint) []User {
	return append(array[:s], array[s+1:]...)
}

func register(w http.ResponseWriter, r *http.Request) {
	currentUser := User{}
	err := json.NewDecoder(r.Body).Decode(&currentUser)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("пока"))
		return
	}

	for _, value := range array {
		if value.Username == currentUser.Username {
			w.Write([]byte("Кыш"))
			return
		}
	}
	currentUser.ID = uint(counter)
	counter += 1
	array = append(array, currentUser)

}
func auth(w http.ResponseWriter, r *http.Request) {
	currentUser := User{}
	err := json.NewDecoder(r.Body).Decode(&currentUser)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Иди нахуй"))
		return
	}
	for _, value := range array {
		if value.Username == currentUser.Username && value.Password == currentUser.Password {
			w.Write([]byte("Авторизация прошла успешно"))
			return
		}
	}
	w.Write([]byte("Иди нахуй"))
}
func add(w http.ResponseWriter, r *http.Request) {
	currentUser := User{}
	err := json.NewDecoder(r.Body).Decode(&currentUser)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Пизда"))

	} else {
		for _, value := range array {
			if value.Name == currentUser.Name {
				w.Write([]byte("Кыш"))
				return
			}
		}
		for _, value := range array {
			if value.Surname == currentUser.Surname {
				w.Write([]byte("Кыш"))
				return
			}
		}
		w.Write([]byte("Информация добавлена"))
	}
}

//Добавить роут на регистрацию POST /register (проверить, что username уникальный)
//Добавить роут на авторизацию POST /auth
//Добавить роут на внесение доп. информации о пользователе, а именно добавить имя и фамилию PATCH /add
//(по желанию) Добавить хеширование для пароля

func main() {

	http.HandleFunc("POST /create", greet)
	http.HandleFunc("GET /getusers", getUsers)
	http.HandleFunc("PUT /update", upd)
	http.HandleFunc("DELETE /delete", del)
	http.HandleFunc("POST /register", register)
	http.HandleFunc("POST /auth", auth)
	http.HandleFunc("PATCH /add", add)
	http.ListenAndServe(":3000", nil)

}
