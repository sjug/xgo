package queenattack

import "fmt"

func CanQueenAttack(w, b string) (attack bool, err error) {
	if err = isValid(w); err != nil {
		return
	}
	if err = isValid(b); err != nil {
		return
	}
	dc := w[0] - b[0]
	dr := w[1] - b[1]
	switch {
	case dc == 0 && dr == 0:
		err = fmt.Errorf("Invalid Argument: Same Square")
	case dc == 0 || dr == 0 || dc == dr || dc == -dr:
		attack = true
	}
	return
}

func isValid(s string) (err error) {
	if len(s) != 2 || s[0] < 'a' || s[0] > 'h' || s[1] > '8' {
		err = fmt.Errorf("Invalid argument: %s", s)
	}
	return
}
