package nullable

import (
	"encoding/json"
)

type String string

func (n *String) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, (*string)(n))
}
