package jwts

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken(JwtPayload{
		Role:     1,
		UserId:   1,
		Username: "zhangsan",
	}, "zhangsan", 8)
	fmt.Println(token, err)
}
func TestParseToken(t *testing.T) {
	payload, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InpoYW5nc2FuIiwicm9sZSI6MSwiZXhwIjoxNzM2MzU3OTE4fQ.880koo8ia3l0ASBPiHFcwCllBFhF5hY8Athhx21Uhoo", "zhangsan", 8)
	fmt.Println(payload, err)
}
