package sign

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnSha256{})
}

type fnSha256 struct {
}

func (fnSha256) Name() string {
	return "sha256"
}

func (fnSha256) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt}, true
}

func (fnSha256) Eval(params ...interface{}) (interface{}, error) {
	s1, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("sha256 function first parameter [%+v] must be []byte", params[0])
	}
	h := sha256.New()
	h.Write([]byte(s1))
	res := hex.EncodeToString(h.Sum(nil))
	return res, err
}
