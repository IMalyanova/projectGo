package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func main() {


	dsn := "root@tcp(localhost:3306)/coursera?"
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"

	db, err := gorm.Open("mysql", dsn)
	db.DB()
	db.DB().Ping()
	//_err_panic(err)

	handlers := &Handler{
		DB:   db,
		Tmpl: template.Must(template.ParseGlob("../gorm templates/*")),
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.List).Methods("GET")
	r.HandleFunc("/items", handlers.List).Methods("GET")
	r.HandleFunc("/items/new", handlers.AddForm).Methods("GET")
	r.HandleFunc("/items/new", handlers.Add).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.Edit).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.Update).Methods("DELETE")

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)
}



type Item struct {
	Id			int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Title		string
	Description string
	Updated		string `sql:"null"`
}



func (i *Item) TableName() string {
	return "items"
}



func (i *Item) BeforeSave() (err error) {
	fmt.Println("trigger on before save")
	return
}



type Handler struct {
	DB 		*gorm.DB
	Tmpl	*template.Template
}



func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

	items := []*Item{}

	db := h.DB.Find(&items)
	err := db.Error

	if err != nil {
		return
	}

	err = h.Tmpl.ExecuteTemplate(w, "index.html", struct {
		Items []*Item
	}{
		Items: items,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {

	newItem := &Item {
		Title:			r.FormValue("title"),
		Description:	r.FormValue("descriptions"),
	}

	db := h.DB.Create(&newItem)
	err := db.Error
	//_err_panic(err)

	affected := db.RowsAffected

	fmt.Println("Insert - RowsAffected", affected, "LastInsertId: ", newItem.Id)

	http.Redirect(w, r, "/", http.StatusFound)
}



func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	//_err_panic(err)

	post := &Item{}

	db := h.DB.Find(post, id)
	err = db.Error

	if err === gorm.ErrRecordNotFound {
		fmt.Println("Record not found", id)
	} else {
		//_err_panic(err)
	}

	err = h.Tmpl.ExecuteTemplate(w, "edit.html", post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	//_err_panic(err)

	post := &Item{}
	h.DB.Find(post, id)

	post.Title = r.FormValue("title")
	post.Description = r.FormValue("description")
	post.Updated = "rvasily"

	db := h.DB.Save(post)
	err = db.Error
	// _err_panic(err)
	affected := db.RowsAffected

	fmt.Println("Update - RowsAffected", affected)

	http.Redirect(w, r, "/", http.StatusFound)
}



func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	//_err_panic(err)

	db := h.DB.Delete(&Item{Id: id})
	err = db.Error
	//_err_panic(err)
	affected := db.RowsAffected

	fmt.Println("Delete - RowsAffected", affected)

	w.Write([]byte(resp))
}




