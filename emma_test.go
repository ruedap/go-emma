package emma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmma_parse(t *testing.T) {
	src := `
    ( pos-s       , position               , static ),
    ( pos-a       , position               , absolute ),
`
	actual, err := parse(src)
	assert.Nil(t, err)

	expected := []decl{
		{
			snippet:  "pos-s",
			property: "position",
			value:    "static",
		},
		{
			snippet:  "pos-a",
			property: "position",
			value:    "absolute",
		},
	}
	assert.Equal(t, actual, expected)
}

func TestEmma_parse_Blank(t *testing.T) {
	actual, err := parse("")
	assert.NotNil(t, err)

	expected := []decl{}
	assert.Equal(t, actual, expected)
}
