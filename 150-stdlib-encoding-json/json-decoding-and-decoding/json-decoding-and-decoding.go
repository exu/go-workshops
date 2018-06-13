package main

import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page        int      `json:"page"`
	Fruits      []string `json:"healthy_fruits"`
	Color       string   `json:"color,omitempty"`
	NotExported int      `json:"-"`
}

func main() {
	p := fmt.Println

	// ----------------------------------------
	// Encoding
	// ----------------------------------------

	h1("Simple types")

	// basic types encoding
	boolJSON, _ := json.Marshal(true)
	p("JSON:bool:", string(boolJSON))
	intJSON, _ := json.Marshal(1)
	p("JSON:int:", string(intJSON))
	floatJSON, _ := json.Marshal(2.34)
	p("JSON:float:", string(floatJSON))
	stringJSON, _ := json.Marshal("gopher")
	p("JSON:string:", string(stringJSON))

	h1("Slices and maps")

	// slice
	sliceJSON, _ := json.Marshal([]string{"apple", "peach", "pear"})
	p("JSON:slice:", string(sliceJSON))
	// map
	mapJSON, _ := json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
	p(string(mapJSON))

	h1("Default Structs")

	// encoding struct w/o tags (annotation)
	response1 := Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	respnse1JSON, _ := json.Marshal(response1)
	p(string(respnse1JSON))

	h1("Annotated structs")

	// We can use tags to configure fields encoded to json
	// look at Response2 struct `json:""` tags added on the
	// right of type definition
	response2 := &Response2{
		Page:        1,
		Fruits:      []string{"apple", "peach", "pear"},
		NotExported: 100,
		Color:       "MAGENTA",
	}
	response2JSON, _ := json.Marshal(response2)
	p("Response2:  std ", string(response2JSON))

	response2 = &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	response2JSON, _ = json.Marshal(response2)
	p("Response2: omit ", string(response2JSON))

	// ----------------------------------------
	// JSON decoding
	// ----------------------------------------

	h1("Decoding")

	// example json bytes (e.g. from request body)
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// we need some data structure where data will be decoded
	// if you are not sure or can't estimate those schemas you can
	// use simpliest `map[string]interface{}` to allow decoding of
	// any json file
	var dat map[string]interface{}

	// decoding with error checking
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	p("DECODED: dat:", dat)

	// if we want to use values from map above (interface{} part)
	// we need to cast values to given type
	num := dat["num"].(float64)
	p("DECODED: dat.num", num)

	// access to embedded data requires series of casts (uuuu how ugly!)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	p("DECODED: ", str1)

	// So how to handle complicated cases?
	// decoding to struct FTW!
	str := `{"page": 1, "healthy_fruits": ["apple", "peach"]}`
	res := Response2{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	p("DECODED: response2: ", res)
	p("DECODED: response2: pierwszy fruit:", res.Fruits[0])

	h1("Stream decoded")

	// Above all examples was done by encoding to array of bytes
	// but as I said before streamiing in go is quite powerful,
	// we can use streams for json too.
	// we can create encoder which recieves io.Writer as parameter
	// in example below there is os.Stdout writer but it could be
	// e.g. http.Response writer which write data directly to HTTP response
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	p("Writing to os.stdOut")
	enc.Encode(d)
}

// h1 helper function
func h1(message string) {
	fmt.Println()
	fmt.Println("---------- ", message, " ----------")
}
