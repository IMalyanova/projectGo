package main

import (
	"Golang/golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
)

type Item struct {
	ID			int
	Title		string
	Description string
	Updated		sql.NullString
}

type Handler struct {
	DB		*sql.DB
	Tmpl	*template.Template
}



func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

	items := []*Item{}

	rows, err := h.DB.Qery("SELECT id, title, updated FROM items")

	if err != nil {
		return
	}

	for rows.Next() {
		post := &Item{}
		err = rows.Scan(&post.Id, &post.Title, &post.Updated)
	}

	if err != nil {
		return
	}

	for rows.Next() {
		post := &Item{}

		err = rows.Scan(&post.Id, &post.Title, &post.Updated)

		if err != nil {
			return
		}

		items = append(items, post)
	}
	//надо закрывать соединение, иначе будет течь
	rows.Close()

	err = h.Tmpl.ExecuteTemplate(w, "index.html", struct {
		Items []*Item
	}{
		Items:items,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func (h *Handler) AddForm(w http.ResponseWriter, r *http.Request) {

	err := h.Tmpl.ExecuteTemplate(w, "create.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



func (h *Hendler) Add(w http.ResponseWriter, r *http.Request) {
	// в целях упрощения примера пропущена валидация

	result, err := h.DB.Exec (
		"INSERT INTO items ('title', 'description') VALUES (?, ?)",
		r.FormValue("title"),
		r.FormValue("description"),
		)

	if err != nil {
		return
	}

	affected, err := result.RowAffected()
	if err != nil {
		return
	}

	lastID, err := result.LastInserId()
	if err != nil {
		return
	}if err != nil {
		return
	}

	fmt.Println("Insert - RowAffected", affected, "LastInsertId: ", lastID)
	http.Redirect(w, r, "/", http.StatusFound)
}



func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		return
	}

	post := &Item{}
	//QueryRow сам закрывает контент
	row := h.DB.QueryRow("SELECT id, title, updated, description FROM" +
		" items WHERE id = ?", id)

	err = row.Scan(&post.Id, &post.Title, &post.Updated, &post.Description)
	//__err_panic(err)

	err = h.Tmpl.ExecuteTemplate(w, "edit.html", post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		return
	}

	// в целях упрощения примера пропущена валидация

	result, err := h.DB.Exec(
		"UPDATE items SET `title` = ?,  `description` = ?" +
			", `updated` = ?" +
			"WHERE id = ?",
			r.FormValue("title"),
			r.FormValue("description"),
			"rvasily",
			id,
		)
	//__err_panic(err)

	affected, err := result.RowsAffected()
	//_err_panic(err)

	fmt.Println("Update - RowsAffected", affected)

	http.Redirect(w, r, "/", http.StatusFound)
}



func(h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := steconv.Atoi(vars[id])

	if err != nil {
		return
	}

	result, err := h.DB.Exec (
		"DELETE FROM items WHERE id = ?", id)

	//_err_panic(err)

	affected, err := result.RowsAffected()
	//_err_panic(err)

	fmt.Println("Delete - RowsAffected", affected)

	w.Header().Set("Content-type", "application/json")
	resp := `{"affected": ` + strconv.Itoa(int(affected)) + `}`
	w.Write([]byte(resp))
}




func main() {

	// основные настройки к базе
	dsn := "root@tcp(localhost:3306)/coursera?"

	// указываем кодировку
	dsn += "&charset=utf8"

	// отказываемся от prapared statements
	// параметры подствляются сразу
	dsn += "&interpolateParams=true"

	db, err := sql.Open("mysql", dsn)

	db.SetMaxOpenConns(10)

	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		panic(err)
	}

	handlers := &Handler {
		DB:   db,
		Tmpl: template.Must(template.ParseGlob("../crud_templates")),
	}

	// в целях упрощения примера пропущена автоизация и csrf
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.List).Methods("GET")
	r.HandleFunc("/items", handlers.List).Methods("GET")
	r.HandleFunc("/items/new", handlers.AddForm).Methods("GET")
	r.HandleFunc("/items/new", handlers.Add).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.Edit).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.Update).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.Delete).Methods("DELETE")

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)
}

// не используйте такой код в прошакшене
// ошибка должна всегда явно обрабатываться
func __err_panic(err error) {
	if err != nil {
		panic(err)
	}
}