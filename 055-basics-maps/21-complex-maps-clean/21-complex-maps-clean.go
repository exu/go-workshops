package main

import (
	"fmt"
)

type h map[string]interface{}

func main() {

	parsedResponse := h{

		"size_in_kb": 12000,
		"custom_headers": []string{
			"X-Booo: your mama",
			"X-Foo: ",
		},
		"html": h{
			"head": h{
				"title": "Some parsed html DOM",
			},
			"body": h{
				"div": "Main page body",
				"form": h{
					"select": h{
						"values": []int{1, 2, 3, 4, 5},
					},
				},
			},
		},
		"comment1": "Some comment",
	}

	fmt.Println(parsedResponse)
}
