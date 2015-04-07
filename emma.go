package emma

import (
	"errors"
	"regexp"
)

// Single declaration
type decl struct {
	snippet  string
	property string
	value    string
}

func parse(src string) ([]decl, error) {
	re := regexp.MustCompile(`\s+\((.+)\,(.+)\,(.+)\)\,.*`)
	res := re.FindAllStringSubmatch(src, -1)
	var ret []decl

	if len(res) < 1 {
		return []decl{}, errors.New("failed to parse source file")
	}

	for _, sl := range res {
		if len(sl) != 3 {
			continue
		}
	}

	return ret, nil
}

const Src string = `
    ( pos-s       , position               , static ),
    ( pos-a       , position               , absolute ),
    ( pos-r       , position               , relative ),
    ( pos-f       , position               , fixed ),
`
