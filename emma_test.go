package emma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmma_ToCSS(t *testing.T) {
	d := decl{"pos-s", "position", "static"}
	actual := ToCSS(d)
	expected := ".u-pos-s { position: static; }\n"
	assert.Equal(t, actual, expected)

	d = decl{"ff-t", "font-family", `"Times New Roman", Times, Baskerville, Georgia, serif`}
	actual = ToCSS(d)
	expected = ".u-ff-t { font-family: \"Times New Roman\", Times, Baskerville, Georgia, serif; }\n"
	assert.Equal(t, actual, expected)
}

func TestEmma_contains_True(t *testing.T) {
	d := decl{"pos-s", "position", "static"}
	actual := contains(d, []string{"s-s"})
	assert.True(t, actual)

	actual = contains(d, []string{"s", "s", "s"})
	assert.True(t, actual)

	actual = contains(d, []string{"static", "position", "pos-s"})
	assert.True(t, actual)
}

func TestEmma_contains_False(t *testing.T) {
	d := decl{"pos-s", "position", "static"}
	actual := contains(d, []string{"pos-a"})
	assert.False(t, actual)

	actual = contains(d, []string{"s-s", "pos-a"})
	assert.False(t, actual)

	actual = contains(d, []string{"s", "s", "z"})
	assert.False(t, actual)
}

func TestEmma_containsDecl_True(t *testing.T) {
	d := decl{"pos-s", "position", "static"}
	actual := containsDecl(d, "s-s")
	assert.True(t, actual)

	actual = containsDecl(d, "ti")
	assert.True(t, actual)

	actual = containsDecl(d, "")
	assert.True(t, actual)
}

func TestEmma_containsDecl_False(t *testing.T) {
	d := decl{"pos-s", "position", "static"}
	actual := containsDecl(d, "pos-a")
	assert.False(t, actual)

	actual = containsDecl(d, "spo")
	assert.False(t, actual)

	actual = containsDecl(d, "ss")
	assert.False(t, actual)
}

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
