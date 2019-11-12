package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}


func SendMessage struct {
	ID        int     `valid:",optional"`
	Priority  string  `valid:"in(low|normal|high)"`
	Recipient string  `schema:"to" valid:"email"`
	Subject   string  `valid:"msgSubject"`
	Inner     string  `schema:"-" valid:"-"`
	flag      int
}



func handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("request " + r.URL.String() + "\n\n"))

	msg := &SendMessage{}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknowKeys(true)
	err := decoder.Decode(msg, r.URL.Query())

	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal", 500)
		return
	}
	w.Write([]byte(fmt.Sprintf("Msg: %#v\n\n", msg)))

	_, err = govalidator.ValidateStruct(msg)

	if err != nil {

		if allErrs, ok := err.(govalidator.Errors); ok {

			for _, fld := range allErrs.Errors() {
				data := []byte(fmt.Sprintf("field: %#v\n\n", fld))
				w.Write(data)
			}
		}
		w.Write([]byte(fmt.Sprintf("error: %s\n\n", err)))
	} else {
		w.Write([]byte(fmt.Sprintf("msg is correct\n\n")))
	}
}



func init() {

	govalidator.CustomTypeTagMap.Set("msgSubject",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)

			if !ok {
				return false
			}
			if len(subject) == 0 || len(subject) > 10 {
				return false
			}
			return true
		}))
}