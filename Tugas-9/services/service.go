package services

import (
	"fmt"
	"strings"
)

func LuasPersegi(sisi int, isSamaSisi bool) interface{} {
	if sisi == 0 && isSamaSisi {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if isSamaSisi {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, sisi*sisi)
	} else if isSamaSisi == false && sisi > 0 {
		return sisi
	}

	return nil
}

func Sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func FormatNumbers(nums []int) string {
	formattedNumbers := make([]string, len(nums))

	for i, num := range nums {
		formattedNumbers[i] = fmt.Sprint(num)
	}

	return strings.Join(formattedNumbers, " + ")
}
