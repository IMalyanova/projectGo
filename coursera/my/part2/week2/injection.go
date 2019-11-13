package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

var (
	db *sql.DB
)

var loginFormTmpl = `
<html>
	<body>
	<form actoin="/login" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>
`



func main() {

	// основные настройки к базе
	dsn := "root@tcp(localhost:3306)/coursera?"

	//указываем кодировку
	dsn += "&charset=utf8"

	//отказываемся от prapared statements
	// параметры подставляются сразу
	dsn += "&inteerpolateParams=true"

	var err error
	//создаем структуру базы
	//но соедиенеие происхоит только при первом запросе
	db, err := sql.Open("mysql", dsn)
	PanicOnErr(err)

	err = db.Ping() // первое подключение к базе
	PanicOnErr(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(loginFormTmpl))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		var(
			id			int
			login, body string
		)

		inputLogin := r.FormValue("login")
		body += fmt.Sprintln("inputLogin:", inputLogin)

		//ПЛОХО! НЕ ДЕЛАТЬ ТАК
		//параметры не экранированы должным образом
		//мы подставляем запрос в параметр как есть
		query := fmt.Sprintf("SELECT id, login FROM users " +
			"WHERE login = '%s' LIMIT 1", inputLogin)

		body += fmt.Sprintln("Sprint query:", query)

		row := db.QueryRow(query)
		err := row.Scan(&id, &login)

		if err == sql.ErrNoRows {
			body += fmt.Sprintln("Sprint case: NOT FOUND")
		} else {
			PanicOnErr(err)
			body += fmt.Sprintln("Sprint case: id:", id, "login:", login)
		}

		//ПРАВИЛЬНО
		//Мы используем плейсхолдерыб там параметры будут эранированы должным образом
		row = db.QueryRow("SELECT id, login FROM users WHERE login = ? LIMIT 1", inputLogin)
		err = row.Scan(&id, &login)

		if err == sql.ErrNoRows {
			body += fmt.Sprintln("Placeholders case: NOT FOUND")
		} else {
			PanicOnErr(err)
			body += fmt.Sprintln("Placeholders id:", id, "login:", login)
		}

		w.Write([]byte(body))
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}



func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}






































