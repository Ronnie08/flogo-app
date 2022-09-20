package datatimex

import (
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnTimestamp{})
}

type fnTimestamp struct {
}

func (fnTimestamp) Name() string {
	return "timestamp"
}

func (fnTimestamp) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt}, true
}

func (fnTimestamp) Eval(params ...interface{}) (interface{}, error) {
	return time.Now().Unix(), nil
}
