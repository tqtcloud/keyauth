package utils

import "testing"

func TestHash(t *testing.T) {
	v := Hash("abc")
	t.Log(v)
}

func TestHashPassword(t *testing.T) {
	v := HashPassword("abc")
	t.Log(v)
	ok := CheckPasswordHash("abc", "$2a$14$z1yyjtyDuohHBVPfhfSVEORmBYO5Tc4kQfJ1espW.2oiytzzgFaOy")
	t.Log(ok)
	ok = CheckPasswordHash("abc", "$2a$14$uK2YKOnmjLGV3lHpiXa6LONhvWd2ja9unZP0pEc7XmyeY2li5KwqC")
	t.Log(ok)
}
