// Ćwiczenie: rot13Reader
// Często spotykanym schematem postępowania jest tworzenie interfejsu typu io.Reader wewnątrz którego znajduje się kolejny io.Reader, przy czym zewnętrzy interfejs posiada inne metody obsługujące strumień danych.
//
// Na przykład, funkcja gzip.NewReader bierze io.Reader (strumień skompresowanych danych) i zwraca *gzip.Reader który również implementuje io.Reader (strumień zdekompresowanych danych).
//
// Zaimplementuj rot13Reader który implementuje io.Reader oraz czyta z tego io.Reader, modyfikując otrzymany strumień po przez użycie szyfru podstawieniowego rot13 do wszystkich liter alfabetu.
//
// Typ rot13Reader jest już dla ciebie zdefiniowany. Zrób z niego interfejs typu io.Reader po przez zaimplementowanie dla niego metody Read.
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case b >= 'A' && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	case b >= 'a' && b <= 'z':
		return 'a' + (b-'a'+13)%26
	default:
		return b
	}
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	el, err := r.r.Read(b)
	for i := 0; i < el; i++ {
		b[i] = rot13(b[i])
	}
	return el, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
