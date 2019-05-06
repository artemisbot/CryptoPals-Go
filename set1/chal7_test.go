package set1

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestChal7(t *testing.T) {

	data, err := ioutil.ReadFile("data/7.txt")
	//fmt.Println(string(data[:len(data)-1]))
	if err != nil {
		panic(err)
	}
	enc, err := base64.StdEncoding.DecodeString(string(data))
	plaintext := Chal7(enc, []byte("YELLOW SUBMARINE"))
	fmt.Println(string(plaintext))
}