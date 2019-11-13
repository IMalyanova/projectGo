package main

import (
	"Golang/golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"database/sql"
	"html/template"
	"net/http"
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

}




func main() {
	
}
