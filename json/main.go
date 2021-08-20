package main

import (
	"encoding/json"
	"fmt"
)

type ThatGirl struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	HotLevel HotLevel
}

type HotLevel struct {
	Level string `json:"level"`
}

func main() {
	hotlevel := HotLevel{Level: "Hell"}
	girl := ThatGirl{Name: "Kendra", Age: 25, HotLevel: hotlevel}

	byteArray, err := json.MarshalIndent(girl, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", string(byteArray))

}
