package vo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocql/gocql"
)

type UUID gocql.UUID
type UUIDs []UUID

func (uuid UUID) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	return gocql.Marshal(info, gocql.UUID(uuid))
}

func (uuid UUID) Bytes() []byte {
	return gocql.UUID(uuid).Bytes()
}

func (uuid UUID) Value() (driver.Value, error) {
	return uuid.Bytes(), nil
}

func (uuid *UUID) Scan(src interface{}) error {
	if src == nil {
		*uuid = UUID(gocql.UUID{})
		return nil
	}

	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("UserId cannot convert %T to bytes", src)
	}

	cqlUUID, err := gocql.ParseUUID(str)
	if err != nil {
		panic(err)
	}

	*uuid = UUID(cqlUUID)
	return nil
}

func (uuid UUID) MarshalJSON() ([]byte, error) {
	return gocql.UUID(uuid).MarshalJSON()
}

func (uuid *UUID) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	cqlUUID, err := gocql.ParseUUID(strings.TrimSpace(str))
	if err != nil {
		return fmt.Errorf("Invalid UUID '%s'", str)
	}

	*uuid = UUID(cqlUUID)
	return nil
}

func (uuids *UUIDs) UnmarshalJSON(data []byte) error {
	var strSlice []string
	if err := json.Unmarshal(data, &strSlice); err != nil {
		return err
	}

	var result []UUID
	for _, strUUID := range strSlice {
		uuid, err := gocql.ParseUUID(strUUID)
		if err != nil {
			return fmt.Errorf("Invalid UUID '%s'", strUUID)
		}
		userID := UUID(uuid)
		result = append(result, userID)
	}

	*uuids = result
	return nil
}

func (uuids UUIDs) ToSlice() []UUID {
	return uuids
}
