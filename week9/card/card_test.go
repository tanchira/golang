package card

import "testing"

func Testallcard(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{12, "3 ดอกจิก"},
		{1, "A โพธิ์แดง"},
		{22, "6 โพธิ์ดำ"},
		{33, "9 โพธิ์แดง"},
		{11, "3 ข้าวหลามตัด"},
		{44, "J ดอกจิก"},
		{13, "4 โพธิ์แดง"},
		{15, "4 ข้าวหลามตัด"},
		{17, "5 โพธิ์แดง"},
		{19, "5 ข้าวหลามตัด"},
		{21, "6 โพธิ์แดง"},
		{23, "6 ข้าวหลามตัด"},
		{25, "7 โพธิ์แดง"},
		{27, "7 ข้าวหลามตัด"},
		{29, "8 โพธิ์แดง"},
		{31, "8 ข้าวหลามตัด"},
		{33, "9 โพธิ์แดง"},
		{35, "9 ข้าวหลามตัด"},
		{39, "10 ข้าวหลามตัด"},
		{41, "J โพธิ์แดง"},
		{43, "J ข้าวหลามตัด"},
		{45, "Q โพธิ์แดง"},
		{47, "Q ข้าวหลามตัด"},
		{49, "K โพธิ์แดง"},
		{51, "K ข้าวหลามตัด"},
		{2, "A โพธิ์ดำ"},
		{4, "A ดอกจิก"},
		{6, "2 โพธิ์ดำ"},
		{8, "2 ดอกจิก"},
		{10, "3 โพธิ์ดำ"},
		{16, "4 ดอกจิก"},
		{18, "5 โพธิ์ดำ"},
		{20, "5 ดอกจิก"},
		{24, "6 ดอกจิก"},
		{26, "7 โพธิ์ดำ"},
		{28, "7 ดอกจิก"},
		{30, "8 โพธิ์ดำ"},
		{32, "8 ดอกจิก"},
		{34, "9 โพธิ์ดำ"},
		{36, "9 ดอกจิก"},
		{38, "10 โพธิ์ดำ"},
		{40, "10 ดอกจิก"},
		{42, "J โพธิ์ดำ"},
		{44, "J ดอกจิก"},
		{46, "Q โพธิ์ดำ"},
		{50, "K โพธิ์ดำ"},
		{48, "Q ดอกจิก"},
		{3, "A ข้าวหลามตัด"},
		{5, "2 โพธิ์แดง"},
		{7, "2 ข้าวหลามตัด"},
		{9, "3 โพธิ์แดง"},
		{14, "4 โพธิ์ดำ"},
	}
	for _, test := range testCases {

}
