package utils

import (
	"encoding/json"
)

func JsonDecodeString(String string) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	json.Unmarshal([]byte(String), &jsonMap)
	return jsonMap
}

func JsonDecodeByte(bytes []byte) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	json.Unmarshal(bytes, &jsonMap)
	return jsonMap
}
func JsonEncodeMapToByte(stringMap map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(stringMap)
	if err != nil {
		return nil
	}
	return jsonBytes
}

func JsonEncode(data interface{}) string {
	s, e := json.Marshal(data)
	if e != nil {
		return ""
	}
	return string(s)
}

func JsonDecode(data string, inter interface{}) error {
	return json.Unmarshal([]byte(data), inter)
}
func JsonDecodeWithBytes(data []byte, inter interface{}) error {
	return json.Unmarshal(data, inter)
}

func StructToMap(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(data)
	_ = json.Unmarshal(j, &m)
	return m
}
