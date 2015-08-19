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
	// ENKODOWANIE
	// ----------------------------------------

	h1("Simple types")

	// encoding podstawowych typów
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
	// mapa
	mapJSON, _ := json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
	p(string(mapJSON))

	h1("Default Structs")

	// mapowanie strukturki do jsona bez annotacji
	response1 := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	respnse1JSON, _ := json.Marshal(response1)
	p(string(respnse1JSON))

	h1("Annotated structs")

	// możemy użyć annotacji aby skonfgurować pola jsonowe w zwracanych strukturkach
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

	h1("Deooding")

	// ----------------------------------------
	// DEKODOWANIE
	// ----------------------------------------

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// musimy zapewnic zmienna do ktorej typu bedziemy w stanie
	// odkodować dane (static typing)

	var dat map[string]interface{}

	// dekodowanie i sprawdzenie czy wystąpiły błędy
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	p("DECODED: dat:", dat)

	// jeżeli chcemy używać typów z jsona zdekodowanego dynanicznie musimy
	// zcastować wartości
	num := dat["num"].(float64)
	p("DECODED: dat.num", num)

	// Dostęp do zagnieżdzonych danych wymaga serii castów (uuuu)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	p("DECODED: ", str1)

	// Jednak na ratunek przychodzi dekodowanie do struktur
	// whoa! type safety na dekodowaniu danych z JSONa
	str := `{"page": 1, "healthy_fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	p("DECODED: response2: ", res)
	p("DECODED: response2: pierwszy fruit:", res.Fruits[0])

	h1("Stream decoded")

	// Powyżej mieliśmy dekodowanie do stringów i bajtów jako pośrednika pomiędzy jsonem
	// a danymi. Jednak możemy również stream'ować JSON'a bezpośrednio do writer'ów
	// przykładowym writerem jest os.Stdout
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	p("Writing to os.stdOut")
	enc.Encode(d)
}

func h1(message string) {
	fmt.Println()
	fmt.Println("---------- ", message, " ----------")
}
