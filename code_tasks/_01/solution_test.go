package _01

import (
	"reflect"
	"testing"
)

var (
	human  = Human{"Jonathan Joestar", 12, "Sutendo Useru"}
	action = Action{human, Human{"Dio Brando", 13, "Za Warudo"}, "basic action", "name", 3}
)

func TestActionInheritsHuman(t *testing.T) {
	actionType := reflect.TypeOf(action)
	humanType := reflect.TypeOf(human)
	for i := 0; i < actionType.NumField(); i++ {
		if actionField := actionType.Field(i); actionField.Anonymous && actionField.Type == humanType {
			t.Skip()
		}
	}
	t.Fatalf("Action type doesn't inherit Human struct")
}
