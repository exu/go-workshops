package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	// testuje czy string zgadza się z narzuconym patternem
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// jezeli chcemy wykorzystywać częściej ten regexo kompilujemy go
	r, _ := regexp.Compile("p([a-z]+)ch")

	// na skompilowanym obiekcie możemy wykonywać operacje
	fmt.Println(r.MatchString("peach"))

	// znajduje match
	fmt.Println(r.FindString("peach punch"))

	// Znajduje indeksy na podstawie patternu
	fmt.Println(r.FindStringIndex("peach punch"))

	// informacje na temat pierwszego napotkanego matcha i submatcha
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// indeks submatch'a
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// wariant "All" odnosi się do wszystkich submatchy w regexp
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatch("peach punch", -1))

	// Limit ilości match'y
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Wszystkie powyższe funkcje miały "String" w nazwie np: MatchString.
	// Jeżeli mamy []byte to możemy ją przekazać do grupy
	// fukcji bez String w nazwie
	fmt.Println(r.Match([]byte("peach")))

	// MustCompile spanikuje jak nie uda się skompilować regexpa
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// Zamiana
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// preg_replace_callback :)
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
