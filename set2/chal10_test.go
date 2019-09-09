package set2

import (
	"CryptoPals/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal10(t *testing.T) {
	pt := []byte("Hello World!\x04\x04\x04\x04")
	assert.Equal(t, pt, util.AESCBCDecrypt(
		util.AESCBCEncrypt(pt, []byte("YELLOW SUBMARINE"), []byte("\x00")),
		[]byte("YELLOW SUBMARINE"), []byte("\x00")))
	fmt.Println(string(Chal10("data/10.txt")))
}
