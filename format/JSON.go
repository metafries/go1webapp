package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type CrewMember struct {
		ID                int      `json:"id,omitempty"`
		Name              string   `json:"name"`
		SecurityClearance int      `json:"clearancelevel"`
		AccessCodes       []string `json:"accesscodes"`
	}

	cm4 := CrewMember{1, "MetaFries", 10, []string{"GIR", "TOT"}}
	cm3 := CrewMember{Name: "Cashviar", SecurityClearance: 20, AccessCodes: []string{"NEW", "TOT"}}

	b4, err := json.Marshal(&cm4)
	PrintError(err)
	b3, err := json.Marshal(&cm3)
	PrintError(err)

	fmt.Println(string(b4))
	fmt.Println(string(b3))

}

func PrintError(err error) {
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
}
