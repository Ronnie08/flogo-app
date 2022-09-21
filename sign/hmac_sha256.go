package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnHmacSha256{})
}

type fnHmacSha256 struct {
}

func (fnHmacSha256) Name() string {
	return "hmac_sha256"
}

func (fnHmacSha256) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (fnHmacSha256) Eval(params ...interface{}) (interface{}, error) {
	msg, err := coerce.ToString(params[0])
	key, err1 := coerce.ToBytes(params[1])
	if err != nil {
		return nil, fmt.Errorf("hmac_sha256 function first parameter [%+v] must be []byte", params[0])
	}
	if err1 != nil {
		return nil, fmt.Errorf("hmac_sha256 function first parameter [%+v] must be []byte", params[1])
	}
	h := hmac.New(sha256.New, key)
	h.Write([]byte(msg))
	res := hex.EncodeToString(h.Sum(nil))
	return res, err
}
