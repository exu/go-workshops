package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	// tests if string match given pattern
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// if we want use given expression more than once it'll be good to compile it
	r, _ := regexp.Compile("p([a-z]+)ch")

	// on compiled object we can call methods
	fmt.Println("MatchString:", r.MatchString("peach"))

	// find matched string
	fmt.Println("FindString:", r.FindString("peach punch"))

	// find string indexed based on given regexp
	fmt.Println("FindStringIndex:", r.FindStringIndex("peach punch"))

	// first match and submatch found
	fmt.Println("FindStringSubmatch:", r.FindStringSubmatch("peach punch"))

	// submatch index
	fmt.Println("FindStringSubmatchIndex:", r.FindStringSubmatchIndex("peach punch"))

	// `All` submatches
	fmt.Println("FindAllString:", r.FindAllString("peach punch pinch", -1))

	fmt.Println("FindAllStringSubmatchIndex:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println("FindAllStringSubmatch:", r.FindAllStringSubmatch("peach punch", -1))

	// you can limit matches
	fmt.Println("FindAllString:", r.FindAllString("peach punch pinch", 2))

	// All methods above have "String" in name
	// If we want to operate on `[]bytes` we can omit "String" in function name
	fmt.Println("Match:", r.Match([]byte("peach")))

	// MustCompile will panic when expression will be invalid
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("Compiled object:", r)

	// replacing string
	fmt.Println("ReplaceAllString:", r.ReplaceAllString("a peach", "<fruit>"))

	// preg_replace_callback from PHP world :)
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("ReplaceAllFunc", string(out))
}
