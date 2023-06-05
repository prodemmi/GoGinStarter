package event

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"reflect"
)

type Event interface {
	setID()
	getID() string
}

type event struct {
	id string
}

func (e *event) setID() {
	e.id = generateID(e)
}

func (e *event) getID() string {
	return e.id
}

func generateID(event *event) string {
	// generate a unique identifier for the event object
	eventValue := reflect.ValueOf(event)
	eventValue = reflect.Indirect(eventValue)
	eventHash := sha256.New()
	eventHash.Write([]byte(eventValue.Type().String()))
	for i := 0; i < eventValue.NumField(); i++ {
		field := eventValue.Field(i)
		fieldValueBytes, err := json.Marshal(field.Interface())
		if err != nil {
			panic(err)
		}
		eventHash.Write(fieldValueBytes)
	}
	return hex.EncodeToString(eventHash.Sum(nil))
}
