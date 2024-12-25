package vo

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"net/mail"
)

type Email string

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(e))
}

func (e *Email) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(str); err != nil {
		return errors.New("Invalid email format")
	}

	*e = Email(str)
	return nil
}

func (e Email) Value() (driver.Value, error) {
	return string(e), nil
}

func (e *Email) Scan(src interface{}) error {
	if src == nil {
		*e = Email("")
		return nil
	}

	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("Email cannot convert %T to string", src)
	}

	*e = Email(str)
	return nil
}
