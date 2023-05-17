package extractor

import (
	"github.com/const-tmp/rotation"
	"regexp"
)

type Env struct {
	re Regex
}

func NewEnv(vg rotation.ValueGetter, varName string) Env {
	return Env{re: NewRegex(vg, regexp.MustCompile(varName+`=(.*)`))}
}

func (e Env) GetValue() string {
	return e.re.GetValue()
}
