package emma

import (
	"errors"
	"regexp"
	"strings"
)

// Single declaration
type decl struct {
	snippet  string
	property string
	value    string
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
			snippet:  strings.TrimSpace(sl[1]),
			property: strings.TrimSpace(sl[2]),
			value:    s,
		}
		ret = append(ret, dec)
	}

	return ret, nil
}
