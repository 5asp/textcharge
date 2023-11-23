package security_test

import (
	"testing"

	"github.com/aheadIV/textcharge/user-service/security"
)

func TestCreatePassword(t *testing.T) {
	password := security.CreatePassword("adnub")
	t.Log(password)
	check := security.CheckPassword(password, "adnub")
	t.Log(check)
}
