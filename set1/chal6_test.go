package set1

import (
	"fmt"
	"github.com/mattnotmitt/CryptoPals-Go/util"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestChal6(t *testing.T) {
	hd := util.HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	expectedHd := 37
	assert.Equal(t, expectedHd, hd)

	data, err := ioutil.ReadFile("data/6.txt")
	util.Check(err)
	enc := util.FromBase64(string(data))
	dec, _ := chal6(enc, 1)
	fmt.Println(string(dec))
}
