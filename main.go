package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

// Структура для хранения данных о напитке
type Drink struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	ImageURL    string   `json:"imageURL"`
	Temperature string   `json:"temperature"`
	Composition []string `json:"composition"`
	Preparation string   `json:"preparation"`
	Price       []string `json:"price"`
}

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/submit", submitHandler)

	// Запуск сервера на порту 8080
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получение данных из формы
		id := r.FormValue("id")
		name := r.FormValue("name")
		imageURL := r.FormValue("imageURL")
		temperature := r.FormValue("temperature")
		composition := r.Form["composition[]"]
		fmt.Println(composition)
		preparation := r.FormValue("preparation")
		price := r.Form["price[]"]

		// Преобразование id в int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Создание объекта Drink
		newDrink := Drink{
			ID:          idInt,
			Name:        name,
			ImageURL:    imageURL,
			Temperature: temperature,
			Composition: composition,
			Preparation: preparation,
			Price:       price,
		}

		// Чтение существующих данных из файла JSON
		var drinks []Drink
		file, err := os.ReadFile("drink.json")
		if err != nil && !os.IsNotExist(err) {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		if len(file) != 0 {
			err = json.Unmarshal(file, &drinks)
			if err != nil {
				http.Error(w, "Error unmarshalling JSON", http.StatusInternalServerError)
				return
			}
		}

		// Добавление нового напитка в список
		drinks = append(drinks, newDrink)

		// Запись обновленного списка обратно в файл JSON
		file, err = json.MarshalIndent(drinks, "", " ")
		if err != nil {
			http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
			return
		}

		err = os.WriteFile("drink.json", file, 0644)
		if err != nil {
			http.Error(w, "Error writing file", http.StatusInternalServerError)
			return
		}

		// Ответ пользователю
		w.Write([]byte("Drink data saved successfully!"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
