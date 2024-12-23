package vo

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/mail"
)

type Email string

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
