package auth

import (
	"fmt"
	"testing"
)

func TestComparePasswords(t *testing.T) {

}

func TestHashAndSalt(t *testing.T) {
	hashAndSalt, _ := HashAndSalt("123456")
	fmt.Printf(hashAndSalt)
}
