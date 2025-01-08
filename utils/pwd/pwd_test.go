package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	hash := HashPwd("123456")
	fmt.Println(hash)
}
func TestCheckPwd(t *testing.T) {
	ok := CheckPwd("$2a$04$roI0MMiJ0kTrL0r8HiaklOX.ydfIkc5fdwIE0hV.B9gFHi5pGqavS", "123456")
	fmt.Println(ok)
}
