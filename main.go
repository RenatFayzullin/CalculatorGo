package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	for _, kv := range rimStr {
		if kv.Value == str {
			return kv.Key, nil
		}
	}
	return "0", err
}

func operation(array []string) string {
	//result, err := "0", error(nil)
	result, err := operationArab(array)

	if err != nil {
		fmt.Println("Ошибка", err)
		return "0"
	}

	if resultArab {
		return result
	} else {
		resultRim, err := parseArabToRim(result)
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

func parseArabToRim(str string) (string, error) {

	value, _ := strconv.Atoi(str)
	if 0 > value {
		return "0", errors.New("В римском нет отрицательных чисел")
	}
	for _, kv := range rimStr {
		if kv.Key == str {
			return kv.Value, error(nil)
		}
	}

	return str + " : Программа выводит числа в римском формате только до X(10)", error(nil)
}

func showFinish(str string) {
	fmt.Println("Результат: ", str)
}
