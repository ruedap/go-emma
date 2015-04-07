package emma

import (
	"reflect"
	"testing"
)

func TestEmma_parse_BlankSrc(t *testing.T) {
	actual, err := parse("")

	if err == nil {
		t.Fatal("expected error but not")
	}

	expected := []decl{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}
