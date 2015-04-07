package emma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmma_Find(t *testing.T) {
	src := `
    ( pos-s       , position               , static ),
    ( pos-a       , position               , absolute ),
    ( pos-r       , position               , relative ),
    ( pos-f       , position               , fixed ),
`
	terms := []string{"position"}
	actual := Find(src, terms)
	expected := []decl{
		{"pos-s", "position", "static"},
		{"pos-a", "position", "absolute"},
		{"pos-r", "position", "relative"},
		{"pos-f", "position", "fixed"},
	}
	assert.Equal(t, actual, expected)
}

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

func TestEmma_ToJSON(t *testing.T) {
	ds := []decl{
		{"pos-s", "position", "static"},
		{"pos-a", "position", "absolute"},
	}

	actual, err := ToJSON(ds)
	assert.Nil(t, err)

	expected := "[{\"Snippet\":\"pos-s\",\"Property\":\"position\",\"Value\":\"static\"},{\"Snippet\":\"pos-a\",\"Property\":\"position\",\"Value\":\"absolute\"}]"
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
			Snippet:  "pos-s",
			Property: "position",
			Value:    "static",
		},
		{
			Snippet:  "pos-a",
			Property: "position",
			Value:    "absolute",
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
			Snippet:  "ti--9999",
			Property: "text-indent",
			Value:    "-9999px",
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
			Snippet:  "ff-t",
			Property: "font-family",
			Value:    `"Times New Roman", Times, Baskerville, Georgia, serif`,
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
