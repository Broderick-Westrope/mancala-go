package mancala_test

import (
	"reflect"
	"testing"
)

func checkEquals(t *testing.T, msg, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v, want %v", msg, got, want)
	}
}
