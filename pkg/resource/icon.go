package resource

import "strings"

type Icon struct {
	Src  Src    `json:"src,omitempty"`
	Name string `json:"name,omitempty"`
}

type Src string

var StorageUrl string

func (str Src) MarshalText() ([]byte, error) {
	s := string(str)

	if s != "" && !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
		s = StorageUrl + s
	}

	return []byte(s), nil
}
