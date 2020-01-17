package card

import (
	"fmt"
)

var store [53]card

type card struct {
	store  string
	symbol string
}

func allcard(want int) string {

	store[0] = card{"A", "โพธิ์แดง"}
	store[1] = card{"A", "โพธิ์ดำ"}
	store[2] = card{"A", "ข้าวหลามตัด"}
	store[3] = card{"A", "ดอกจิก"}
	store[4] = card{"2", "โพธิ์แดง"}
	store[5] = card{"2", "โพธิ์ดำ"}
	store[6] = card{"2", "ข้าวหลามตัด"}
	store[7] = card{"2", "ดอกจิก"}
	store[8] = card{"3", "โพธิ์แดง"}
	store[9] = card{"3", "โพธิ์ดำ"}
	store[10] = card{"3", "ข้าวหลามตัด"}
	store[11] = card{"3", "ดอกจิก"}
	store[12] = card{"4", "โพธิ์แดง"}
	store[13] = card{"4", "โพธิ์ดำ"}
	store[14] = card{"4", "ข้าวหลามตัด"}
	store[15] = card{"4", "ดอกจิก"}
	store[16] = card{"5", "โพธิ์แดง"}
	store[17] = card{"5", "โพธิ์ดำ"}
	store[18] = card{"5", "ข้าวหลามตัด"}
	store[19] = card{"5", "ดอกจิก"}
	store[20] = card{"6", "โพธิ์แดง"}
	store[21] = card{"6", "โพธิ์ดำ"}
	store[22] = card{"6", "ข้าวหลามตัด"}
	store[23] = card{"6", "ดอกจิก"}
	store[24] = card{"7", "โพธิ์แดง"}
	store[25] = card{"7", "โพธิ์ดำ"}
	store[26] = card{"7", "ข้าวหลามตัด"}
	store[27] = card{"7", "ดอกจิก"}
	store[28] = card{"8", "โพธิ์แดง"}
	store[29] = card{"8", "โพธิ์ดำ"}
	store[30] = card{"8", "ข้าวหลามตัด"}
	store[31] = card{"8", "ดอกจิก"}
	store[32] = card{"9", "โพธิ์แดง"}
	store[33] = card{"9", "โพธิ์ดำ"}
	store[34] = card{"9", "ข้าวหลามตัด"}
	store[35] = card{"9", "ดอกจิก"}
	store[36] = card{"10", "โพธิ์แดง"}
	store[37] = card{"10", "โพธิ์ดำ"}
	store[38] = card{"10", "ข้าวหลามตัด"}
	store[39] = card{"10", "ดอกจิก"}
	store[40] = card{"J", "โพธิ์แดง"}
	store[41] = card{"J", "โพธิ์ดำ"}
	store[42] = card{"J", "ข้าวหลามตัด"}
	store[43] = card{"J", "ดอกจิก"}
	store[44] = card{"Q", "โพธิ์แดง"}
	store[45] = card{"Q", "โพธิ์ดำ"}
	store[46] = card{"Q", "ข้าวหลามตัด"}
	store[47] = card{"Q", "ดอกจิก"}
	store[48] = card{"K", "โพธิ์แดง"}
	store[49] = card{"K", "โพธิ์ดำ"}
	store[50] = card{"K", "ข้าวหลามตัด"}
	store[51] = card{"K", "ดอกจิก"}
	return fmt.Sprintf("%v", store[(want%52)-1])

}
