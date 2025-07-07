package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type TimeWrapper struct {
	time.Time
}

// Implement database/sql.Scanner interface
func (t *TimeWrapper) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case []byte:
		return json.Unmarshal(v, &t.Time)
	case string:
		return t.UnmarshalGQL(v)
	case nil:
		t.Time = time.Time{}
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into TimeWrapper", value)
	}
}

// Implement database/sql/driver.Valuer interface
func (t TimeWrapper) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

// Implement graphql.Marshaler interface
func (t TimeWrapper) MarshalGQL(w io.Writer) {
	if t.IsZero() {
		w.Write([]byte("null"))
		return
	}
	w.Write([]byte(fmt.Sprintf(`"%s"`, t.Time.Format(time.RFC3339))))
}

// Implement graphql.Unmarshaler interface
func (t *TimeWrapper) UnmarshalGQL(v interface{}) error {
	if v == nil {
		t.Time = time.Time{}
		return nil
	}

	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("TimeWrapper must be a string")
	}

	parsed, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}