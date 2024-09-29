package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, b, c string
	var res, num, num2 int

	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	input := myscanner.Text()

	var data []string
	data = strings.Split(input, " ")
	if len(data) > 3 {
		panic("ошибка ввода")
	}
	a = data[0]
	b = data[1]
	c = data[2]

	numbers := map[string]string{
		"I":        "1",
		"II":       "2",
		"III":      "3",
		"IV":       "4",
		"V":        "5",
		"VI":       "6",
		"VII":      "7",
		"VIII":     "8",
		"IX":       "9",
		"X":        "10",
		"XI":       "11",
		"XII":      "12",
		"XIII":     "13",
		"XIV":      "14",
		"XV":       "15",
		"XVI":      "16",
		"XVII":     "17",
		"XVIII":    "18",
		"XIX":      "19",
		"XX":       "20",
		"XXI":      "21",
		"XXII":     "22",
		"XXIII":    "23",
		"XXIV":     "24",
		"XXV":      "25",
		"XXVI":     "26",
		"XXVII":    "27",
		"XXVIII":   "28",
		"XXIX":     "29",
		"XXX":      "30",
		"XXXI":     "31",
		"XXXII":    "32",
		"XXXIII":   "33",
		"XXXIV":    "34",
		"XXXV":     "35",
		"XXXVI":    "36",
		"XXXVII":   "37",
		"XXXVIII":  "38",
		"XXXIX":    "39",
		"XL":       "40",
		"XLI":      "41",
		"XLII":     "42",
		"XLIII":    "43",
		"XLIV":     "44",
		"XLV":      "45",
		"XLVI":     "46",
		"XLVII":    "47",
		"XLVIII":   "48",
		"XLIX":     "49",
		"L":        "50",
		"LI":       "51",
		"LII":      "52",
		"LIII":     "53",
		"LIV":      "54",
		"LV":       "55",
		"LVI":      "56",
		"LVII":     "57",
		"LVIII":    "58",
		"LIX":      "59",
		"LX":       "60",
		"LXI":      "61",
		"LXII":     "62",
		"LXIII":    "63",
		"LXIV":     "64",
		"LXV":      "65",
		"LXVI":     "66",
		"LXVII":    "67",
		"LXVIII":   "68",
		"LXIX":     "69",
		"LXX":      "70",
		"LXXI":     "71",
		"LXXII":    "72",
		"LXXIII":   "73",
		"LXXIV":    "74",
		"LXXV":     "75",
		"LXXVI":    "76",
		"LXXVII":   "77",
		"LXXVIII":  "78",
		"LXXIX":    "79",
		"LXXX":     "80",
		"LXXXI":    "81",
		"LXXXII":   "82",
		"LXXXIII":  "83",
		"LXXXIV":   "84",
		"LXXXV":    "85",
		"LXXXVI":   "86",
		"LXXXVII":  "87",
		"LXXXVIII": "88",
		"LXXXIX":   "89",
		"XC":       "90",
		"XCI":      "91",
		"XCII":     "92",
		"XCIII":    "93",
		"XCIV":     "94",
		"XCV":      "95",
		"XCVI":     "96",
		"XCVII":    "97",
		"XCVIII":   "98",
		"XCIX":     "99",
		"C":        "100",
	}
	arabicNumbers := map[string]bool{
		"1":  true,
		"2":  true,
		"3":  true,
		"4":  true,
		"5":  true,
		"6":  true,
		"7":  true,
		"8":  true,
		"9":  true,
		"10": true,
	}
	if (arabicNumbers[a] && !arabicNumbers[c]) || (!arabicNumbers[a] && arabicNumbers[c]) {
		panic("Ошибка ввода")
	}

	slice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	for j := 0; j <= 9; j++ {
		if slice[j] == a {
			num, _ = strconv.Atoi(a)
			num2, _ = strconv.Atoi(c)
			switch {
			case b == "+":
				res = num + num2
			case b == "-":
				res = num - num2
			case b == "*":
				res = num * num2
			case b == "/":
				res = num / num2
			}
			fmt.Print(res)
			break
		}

	}
	if !arabicNumbers[a] && !arabicNumbers[c] {
		for key, val := range numbers {
			if key == a {
				a = val
				break
			}
		}
		for key, val := range numbers {
			if key == c {
				c = val
			}
		}

		num, _ = strconv.Atoi(a)
		num2, _ = strconv.Atoi(c)

		if num < 1 || num > 10 || num2 < 1 || num2 > 10 {
			panic("Ошибка ввода")
		}

		switch b {
		case "+":
			res = num + num2
		case "-":
			res = num - num2
		case "*":
			res = num * num2
		case "/":
			if num2 == 0 {
				panic("Ошибка ввода")
			}
			res = num / num2
		default:
			panic("Ошибка ввода")
		}
		if res < 1 {
			panic("Ошибка ввода")
		}

		if !arabicNumbers[a] && !arabicNumbers[c] && res < 1 {
			panic("Ошибка ввода")
		}
		if arabicNumbers[a] && arabicNumbers[c] && res < 0 {
			fmt.Print(res)
		} else {
			str := strconv.Itoa(res)
			for key, val := range numbers {
				if val == str {
					str = key
					fmt.Print(str)
					break
				}
			}
		}
	}
}
