package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Username string

const MinUsernameLength = 3
const MaxUsernameLength = 50

func (u Username) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(u))
}

func (u *Username) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	strLen := len(str)
	if strLen < MinUsernameLength {
		return fmt.Errorf("Username must be at least %d characters long", MinUsernameLength)
	} else if strLen > MaxUsernameLength {
		return fmt.Errorf("Username must be at most %d characters long", MaxUsernameLength)
	}

	*u = Username(str)
	return nil
}

func (u Username) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *Username) Scan(src interface{}) error {
	if src == nil {
		*u = Username("")
		return nil
	}

	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("Username cannot convert %T to string", src)
	}

	*u = Username(str)
	return nil
}
