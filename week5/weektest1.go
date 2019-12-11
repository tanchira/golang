package week5_6

import "strconv"

//FizzWord รับพารามิเตอร์ 2 ตัว เมื่อนำตัวแรกมาหารด้วยตัวที่สองลงตัวจะคืนค่าเป็น "Fizz" หากไม่ลงตัวจะคืนค่าเป็นพารามิตเอร์ตัวแรกแต่มีไทป์เป็น string

func FizzWord(number int, mod int) string {

	if number%mod == 0 {
		return "fizz"
	}
	return strconv.Itoa(number)
}
