package sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnMd5{})
}

type fnMd5 struct {
}

func (fnMd5) Name() string {
	return "md5"
}

func (fnMd5) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (fnMd5) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("md5 function first parameter [%+v] must be []byte", params[0])
	}
	h := md5.New()
	h.Write([]byte(s1))
	res := hex.EncodeToString(h.Sum(nil))
	return res, err
}
