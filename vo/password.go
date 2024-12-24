package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Password string
type HashedPassword string

const PasswordMinLength = 8
const PasswordMaxLength = 64

func (p Password) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

func (p *Password) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	pwdLen := len(str)
	if pwdLen < PasswordMinLength {
		return fmt.Errorf("Password must be at least %d characters long", PasswordMinLength)
	} else if pwdLen > PasswordMaxLength {
		return fmt.Errorf("Password must be at most %d characters long", PasswordMaxLength)
	}

	*p = Password(str)
	return nil
}

func (p Password) Bytes() []byte {
	return []byte(string(p))
}

func (p HashedPassword) Value() (driver.Value, error) {
	return string(p), nil
}

func (p HashedPassword) Bytes() []byte {
	return []byte(string(p))
}
