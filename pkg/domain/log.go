package audit

import (
	"errors"
	"time"
)

const (
	ENTITY_USER = "USER"
	ENTITY_BOOK = "BOOK"

	ACTION_CREATE   = "CREATE"
	ACTION_UPDATE   = "UPDATE"
	ACTION_GET      = "GET"
	ACTION_DELETE   = "DELETE"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
)

// Bind const strings to LogRequest's variables
var (
	entities = map[string]LogRequest_Entities{
		ENTITY_USER: LogRequest_USER,
		ENTITY_BOOK: LogRequest_BOOK,
	}

	actions = map[string]LogRequest_Actions{
		ACTION_REGISTER: LogRequest_REGISTER,
		ACTION_LOGIN:    LogRequest_LOGIN,
		ACTION_CREATE:   LogRequest_CREATE,
		ACTION_UPDATE:   LogRequest_UPDATE,
		ACTION_GET:      LogRequest_GET,
		ACTION_DELETE:   LogRequest_DELETE,
	}
)

type LogItem struct {
	Action    string    `bson:"action"`
	Entity    string    `bson:"entity"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}

func ToPbEntity(entity string) (LogRequest_Entities, error) {
	val, ex := entities[entity]
	if !ex {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}

func ToPbAction(action string) (LogRequest_Actions, error) {
	val, ex := actions[action]
	if !ex {
		return 0, errors.New("invalid action")
	}

	return val, nil
}
