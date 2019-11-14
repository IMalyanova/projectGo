package main

import (
	"context"
	"coursera/microservices/gateway/session"
	"net/http"
)

var loginFormTmpl = []byte(`
<html>
	<body>
	<form action="/login" method="post">
	Login: <input type="text" name="login">
	Password: <input type="password" name="password">
	<input type="submit" value="Login">
	</form>
	</body>
</html>
`)

var (
	sessManager session.AuthCheckerClient
)


func checkSession(r *http.Request) (*session.Session, error) {
	cookieSessionID, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	sess, err := seeManager.Check(
		context.Background(),
		&session.SessionID{
			ID: cookieSessionID.Value,
		})
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func innerPage(w http.ResponseWriter, r *http.Request) {
	sess, err := checkSession(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if see == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if see == nil {
		w.Write(loginFormTmpl)
		return
	}
}