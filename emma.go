package emma

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Single declaration
type decl struct {
	Snippet  string
	Property string
	Value    string
}

func Find(src string, terms []string) []decl {
	decls, err := parse(src)
	if err != nil {
		return []decl{}
	}

	var ret []decl
	for _, d := range decls {
		if contains(d, terms) {
			ret = append(ret, d)
		}
	}
	return ret
}

func ToCSS(decls []decl) string {
	var str string
	for _, d := range decls {
		str += fmt.Sprintf(".u-%s { %s: %s; }\n", d.Snippet, d.Property, d.Value)
	}

	return str
}

func ToJSON(decls []decl) (string, error) {
	if len(decls) == 0 {
		return "[]", nil
	}

	b, err := json.Marshal(decls)
	return string(b), err
}

func contains(d decl, terms []string) bool {
	for _, t := range terms {
		if !containsDecl(d, t) {
			return false
		}
	}

	return true
}

func containsDecl(d decl, term string) bool {
	if strings.Contains(d.Snippet, term) {
		return true
	}

	if strings.Contains(d.Property, term) {
		return true
	}

	if strings.Contains(d.Value, term) {
		return true
	}

	return false
}

func parse(src string) ([]decl, error) {
	re := regexp.MustCompile(`\s+\((.+?)\,(.+?)\,(.+)\)\,.*`)
	res := re.FindAllStringSubmatch(src, -1)
	var dec decl
	var ret []decl

	if len(res) < 1 {
		return []decl{}, errors.New("failed to parse source file")
	}

	for _, sl := range res {
		if len(sl) != 4 {
			continue
		}

		s := strings.TrimSpace(sl[3])
		switch {
		case s[0] == `'`[0] && s[len(s)-1] == `'`[0]:
			s = strings.Trim(s, `'`)
		case s[0] == `"`[0] && s[len(s)-1] == `"`[0]:
			s = strings.Trim(s, `"`)
		}

		dec = decl{
			Snippet:  strings.TrimSpace(sl[1]),
			Property: strings.TrimSpace(sl[2]),
			Value:    s,
		}
		ret = append(ret, dec)
	}

	return ret, nil
}
