package extractor

import (
	"github.com/const-tmp/rotation"
	"regexp"
)

type Regex struct {
	vg rotation.ValueGetter
	re *regexp.Regexp
}

func NewRegex(vg rotation.ValueGetter, re *regexp.Regexp) Regex {
	return Regex{vg: vg, re: re}
}

func (r Regex) GetValue() string {
	return r.re.FindStringSubmatch(r.vg.GetValue())[1]
}
