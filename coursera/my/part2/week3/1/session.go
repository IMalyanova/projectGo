package main

import (
	"math/rand"
	"sync"
)

type Session struct {
	Login 		string
	Useragent	string
}

type SessionID struct {
	ID string
}

const sessKeyLen = 10

//выкидываем глобальные переменные
//var (
//	sessions = map[SessionID]*Session{}
//	mu		= &sync.RWMutex{}
//)

type SessionManagerI interface {

	Create(*Session) (*SessionID, error)
	Check(*SessionID) *Session
	Delete(*SessionID)
}

type SessionManager struct {
	mu 		*sync.Mutex
	session map[SessionID]*Session{}
}

func NewSessManager(){
	return &SessionManager{
		mu:      &sync.Mutex{},
		session: map[SessionID]*Session{}, //хранилище сессий
	}
}

func (sm *SessionManager) Create(in *Session) (*SessionID, error) {
	sm.mu.Lock()
	id := SessionID{RandStringRunes(sessKeyLen)}
	sm.mu.Unlock()
	sessions[id] = in
	return &id, nil
}

func (sm *SessionManager) Check(in *SessionID) *Session {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if sess, ok := sessions[*in]; ok {
		return sess
	}
	return nil
}

func (sm *SessionManager) Delete(in *SessionID) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sessions, *in)
}

var letterRunes = []rune("msffnsirsndfigldrndbgfmlcm")

func RandStringRunes(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}