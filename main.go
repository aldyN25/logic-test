package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Example:
// maxRedigit(123) --> 321
// maxRedigit(231) --> 321
// maxRedigit(321) --> 321
// maxRedigit(-1) --> null
// maxRedigit(0) --> null
// maxRedigit(99) --> null
// maxRedigit(1000) --> null

// Given an array/list [] of integers , Construct a product array
// Of same size Such That prod[i] is equal to The Product of all
// the elements of Arr[] except Arr[i].

// alternateCase("abc") => "ABC"
// alternateCase("ABC") => "abc"
// alternateCase("Hello World") => "hELLO wORLD"

// Create solution function that accept 1 parameter that be will multiple number 3 and 5
// while each result of that multiplication is still lower than parameter inputed
// solution (10) *=> 23 = 3 + 5 + 6 + 9*
// *solution (20)* => 78 = 3 + 5 + 6 + 9 + 10 + 12 + 15 + 18
func main() {
	fmt.Println(*BiggestNumber(123))
	ProductArray([]int{1, 2, 3, 4})
	AlternateCase("HelLo WorlD")
	FizzBuzzcabangLaen(20)
	NearestFibonacci(18)
	ReverseString("I am A Great human")
}
func findMiddleAlphabet() {

}
func ReverseString(s string) {
	var result string
	str := strings.Split(s, " ")
	for index, v := range str {
		c := strings.Split(v, "")

		for i := len(c) - 1; i >= 0; i-- {
			if isUpper := c[len(c)-i-1]; isUpper == strings.ToUpper((isUpper)) {
				u := strings.ToUpper(c[i])
				result += u
			} else {
				result += strings.ToLower(c[i])
			}
		}
		if index != len(str)-1 {
			result += " "
		}
	}
	fmt.Println(result)

}

func NearestFibonacci(num int) {

	if num == 0 {
		return
	}

	prev := 0
	next := 1

	current := prev + next

	for current <= num {
		prev = next

		next = current

		current = prev + next
	}
	fmt.Println(prev)
	fmt.Println(next)
	fmt.Println(current)
	fmt.Println(num)
	res := 0
	if current-num < num-next {
		res = current
	} else {
		res = next
	}
	fmt.Println(res)
}

func FizzBuzzcabangLaen(num int) {
	result := 0

	for i := 3; i < num; i++ {
		if i%3 == 0 && i%5 == 0 {
			result += i
			continue
		}
		if i%3 == 0 {
			result += i
		}
		if i%5 == 0 {
			result += i
		}
	}
	fmt.Println(result)
}

func AlternateCase(d string) {
	str := ""
	temp := strings.Split(d, "")
	for i := 0; i < len(d); i++ {
		if strings.ToUpper(temp[i]) == temp[i] {
			str += strings.ToLower(temp[i])
		} else {
			str += strings.ToUpper(temp[i])
		}
	}
	fmt.Println(str)
}

func ProductArray(data []int) {
	var result []int
	for i := range data {
		product := 1
		for j, k := range data {
			if i == j {
				continue
			}
			product *= k
		}
		result = append(result, product)
	}
	fmt.Println(result)
}
func BiggestNumber(num int) *string {
	if num <= 100 || num >= 1000 {
		return nil
	}

	s := strconv.Itoa(num)
	str := strings.Split(s, "")
	sort.Strings(str)
	newStr := strings.Join(str, "")
	r := []rune(newStr)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	newR := string(r)
	return &newR
}
