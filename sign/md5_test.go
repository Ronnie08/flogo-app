package sign

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

var s = &fnMd5{}

func init() {
	function.ResolveAliases()
}

func TestMd5(t *testing.T) {
	final1, _ := s.Eval("12345623902390adadliiojjlloihhggu6@@@#@#fdafl;adjad0dufad9fdfadf")
	assert.NotNil(t, final1)
	fmt.Print(final1)
}
