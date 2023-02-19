package lib

import "encoding/json"

func ToJSON(i interface{}) []byte {
	r, err := json.Marshal(i)
	HandleErr(err)
	return r
}

func FromJSON(bytes []byte, i interface{}) {
	err := json.Unmarshal(bytes, i)
	HandleErr(err)
}
