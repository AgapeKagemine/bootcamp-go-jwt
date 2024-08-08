package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Check Password with Hash
func Check(h, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	fmt.Println(err)
	return err == nil
}
