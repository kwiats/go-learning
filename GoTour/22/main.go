//Ćwiczenie: Readers
//Zaimplementuj typ Reader który wysyła nieskończony strumień zawierający tylko znak ASCII 'A'.

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Dodaj metodę Read([]byte) (int, error) do MyReader.

func (my MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
