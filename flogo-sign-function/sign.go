package datetimex

import (
	"crypto/sha256"
	"crypto/md5"
	"time"
	"fmt"
)

type Sha256 struct {}
type Md5 struct {}
type GetTimestamp struct {}

func init() {
	_ = function.Register(&Sha256{})
	_ = function.Register(&Md5{})
}

func (s *Sha256) Name() (src string) {
	
	h := sha256.New()
	h.Write([]byte(src))
    res := hex.EncodeToString(h.Sum(nil))
    return res
}

func (s *Md5) Name() (src string) {
	
	h := md5.New()
	h.Write([]byte(src))
    res := hex.EncodeToString(h.Sum(nil))
    return res
}

func (s *GetTimestamp) Name() {
	
	timestamp = time.Now().Unix()
    return res
}