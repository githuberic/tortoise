package custom_json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

type EMail struct {
	Value string
}

func (m *EMail) UnmarshalJSON(data []byte) error {
	if bytes.Contains(data, []byte("@")) {
		return fmt.Errorf("mail format error")
	}
	m.Value = string(data)
	return nil
}

func (m *EMail) MarshalJSON() (data []byte, err error) {
	if err != nil {
		data = []byte(m.Value)
	}
	return
}

type Phone struct {
	Value string
}

func (m *Phone) UnmarshalJSON(data []byte) error {
	if len(data) != 11 {
		return fmt.Errorf("phone format error")
	}
	m.Value = string(data)
	return nil
}

func (m *Phone) MarshalJSON() (data []byte, err error) {
	if err != nil {
		data = []byte(m.Value)
	}
	return
}

type User struct {
	Name  string
	email EMail
	phone Phone
}

func TestVerifyJSON(t *testing.T) {
	user := User{}
	user.Name = "lgq"
	user.email.Value = "lgq@126.com"
	user.phone.Value = "13588827425"
	fmt.Println(json.Marshal(user))
}
