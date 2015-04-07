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

func TestEmma_parse_Comment(t *testing.T) {
	src := `
    ( ti--9999    , text-indent            , -9999px ),             // Emmet: ti-
`
	actual, err := parse(src)
	assert.Nil(t, err)

	expected := []decl{
		{
			snippet:  "ti--9999",
			property: "text-indent",
			value:    "-9999px",
		},
	}
	assert.Equal(t, actual, expected)
}

func TestEmma_parse_FontFamily(t *testing.T) {
	src := `
    ( ff-t        , font-family            , '"Times New Roman", Times, Baskerville, Georgia, serif' ),
`
	actual, err := parse(src)
	assert.Nil(t, err)

	expected := []decl{
		{
			snippet:  "ff-t",
			property: "font-family",
			value:    `"Times New Roman", Times, Baskerville, Georgia, serif`,
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