package main

import "fmt"

func main() {

	parsedResponse := map[string]interface{}{
		"size_in_kb": 12000,
		"custom_headers": []string{
			"X-Booo: your mama",
			"X-Foo: ",
		},
		"html": map[string]interface{}{
			"head": map[string]interface{}{
				"title": "Some parsed html DOM",
			},
			"body": map[string]interface{}{
				"div": "Main page body",
				"form": map[string]interface{}{
					"select": map[string]interface{}{
						"values":  []int{1, 2, 3, 4, 5},
						"values2": []string{"A"},
					},
				},
			},
		},
		"comment1": "Some comment",
	}

	fmt.Println(parsedResponse)
}
