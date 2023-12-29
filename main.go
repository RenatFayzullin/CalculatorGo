package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type KeyValue struct {
	Key   string
	Value string
}

var rimStr = []KeyValue{
	{Key: "1", Value: "I"},
	{Key: "2", Value: "II"},
	{Key: "3", Value: "III"},
	{Key: "4", Value: "IV"},
	{Key: "5", Value: "V"},
	{Key: "6", Value: "VI"},
	{Key: "7", Value: "VII"},
	{Key: "8", Value: "VIII"},
	{Key: "9", Value: "IX"},
	{Key: "10", Value: "X"},
}

var resultArab = true

func main() {
	showStart()
	inputText := getInputText()
	arrString, err := checkText(inputText)
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	result := operation(arrString)
	showFinish(result)
}

func showStart() {
	fmt.Println("Введите операцию:")
}

func getInputText() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputText := scanner.Text()
	return inputText
}

func checkText(text string) ([]string, error) {
	temp := strings.ReplaceAll(text, " ", "")
	splitString, err := splitByMathOperator(temp)

	if err != nil {
		return nil, err
	}

	return splitString, error(nil)

}

func splitByMathOperator(input string) ([]string, error) {
	var result []string
	current := ""
	countOperator := 0

	for _, char := range input {
		if unicode.IsSpace(char) || isMathOperator(char) {
			if current != "" {
				result = append(result, current)
				current = ""
			}
			if isMathOperator(char) && countOperator < 1 {
				result = append(result, string(char))
				countOperator++
			} else {
				return nil, errors.New("В строке должен только 1 оператор")
			}
		} else {
			current += string(char)
		}
	}
	if current != "" {
		result = append(result, current)
	}

	if len(result) != 3 {
		return nil, errors.New("Неверный формат")
	}

	errArab := checkArabStr(result)

	if errArab {

		result, errRim := checkRimStr(result)

		if errRim {
			return nil, errors.New("Операнд не явлется числом в римском или арабском виде")
		}
		resultArab = false
		return result, nil

	} else {
		return result, nil
	}

}

func isMathOperator(char rune) bool {
	return char == '+' || char == '-' || char == '/' || char == '*'
}

func parseInt(str string) error {
	_, err := strconv.Atoi(str)
	return err
}

func parseRim(str string) (string, error) {
	err := errors.New("Error")
	str = strings.ToUpper(str)
	romanPattern := regexp.MustCompile(`^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	if romanPattern.MatchString(str) {
		arabicNumber, err := romanToArabic(str)
		return arabicNumber, err
	}
	return "0", err
}

func operation(array []string) string {
	result, err := operationArab(array)

	if err != nil {
		fmt.Println("Ошибка", err)
		return "0"
	}

	if resultArab {
		return result
	} else {
		resultRim, err := arabicToRoman(result)
		if err != nil {
			fmt.Println("Ошибка", err)
		}
		return resultRim
	}

}

func operationArab(array []string) (string, error) {
	op1, _ := strconv.Atoi(array[0])
	op2, _ := strconv.Atoi(array[2])

	operator := array[1]

	switch operator {
	case "+":
		return strconv.Itoa(op1 + op2), nil
	case "-":
		return strconv.Itoa(op1 - op2), nil
	case "*":
		return strconv.Itoa(op1 * op2), nil
	case "/":
		if op2 == 0 {
			return "0", errors.New("Делать на ноль нельзя")
		}
		if op1%op2 != 0 {
			return strconv.Itoa(op1/op2) + " с остатком : " + strconv.Itoa(op1%op2), nil
		}
		return strconv.Itoa(op1 / op2), nil
	default:
		return "0", nil
	}
}

func checkArabStr(arrayStr []string) bool {
	err1 := parseInt(arrayStr[0])
	err2 := parseInt(arrayStr[2])

	if err1 != nil || err2 != nil {
		return true
	}
	return false
}

func checkRimStr(arrayStr []string) ([]string, bool) {
	err3 := error(nil)
	err4 := error(nil)
	arrayStr[0], err3 = parseRim(arrayStr[0])
	arrayStr[2], err4 = parseRim(arrayStr[2])

	if err3 != nil || err4 != nil {
		return nil, true
	}
	return arrayStr, false
}

func showFinish(str string) {
	fmt.Println("Результат: ", str)
}

func romanToArabic(s string) (string, error) {
	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result int
	previousValue := 0

	for _, char := range strings.Split(s, "") {
		value, found := romanNumerals[char]
		if !found {
			return "", fmt.Errorf("некорректный символ римского числа: %s", char)
		}

		result += value
		if previousValue < value {
			result -= 2 * previousValue
		}

		previousValue = value
	}

	return fmt.Sprintf("%d", result), nil
}

func arabicToRoman(s string) (string, error) {
	err := errors.New("Аргумент должен быть в диапазоне от 1 до 3999")
	num, err := strconv.Atoi(s)
	if num <= 0 || num > 3999 {
		return "", err
	}

	romanNumerals := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	result := ""
	for _, value := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for num >= value {
			result += romanNumerals[value]
			num -= value
		}
	}

	return result, nil
}
