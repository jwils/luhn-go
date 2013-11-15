package luhn

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

var doubleSumDelta = [10]int{0, 1, 2, 3, 4, -4, -3, -2, -1, 0}

func Valid(digits string) bool {
	return calculateLuhnValue(digits) == 0
}

func calculateLuhnValue(digits string) int {
	sum := 0
	source := strings.Split(digits, "")
	for i := len(source) - 1; i >= 0; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)
		sum += n - ((len(source)-i)%2-1)*doubleSumDelta[n]
	}
	return (sum*9)%10
}

func randomString(length int) string {
 return fmt.Sprintf("%0" + strconv.Itoa(length) + "d", rand.Intn(int(math.Pow(10, float64(length)))))
}
