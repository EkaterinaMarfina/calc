package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c, str string
	var res, num, num2 int
	fmt.Scan(&a, &b, &c)

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
		"1":   true,
		"2":   true,
		"3":   true,
		"4":   true,
		"5":   true,
		"6":   true,
		"7":   true,
		"8":   true,
		"9":   true,
		"10":  true,
		"11":  true,
		"12":  true,
		"13":  true,
		"14":  true,
		"15":  true,
		"16":  true,
		"17":  true,
		"18":  true,
		"19":  true,
		"20":  true,
		"21":  true,
		"22":  true,
		"23":  true,
		"24":  true,
		"25":  true,
		"26":  true,
		"27":  true,
		"28":  true,
		"29":  true,
		"30":  true,
		"31":  true,
		"32":  true,
		"33":  true,
		"34":  true,
		"35":  true,
		"36":  true,
		"37":  true,
		"38":  true,
		"39":  true,
		"40":  true,
		"41":  true,
		"42":  true,
		"43":  true,
		"44":  true,
		"45":  true,
		"46":  true,
		"47":  true,
		"48":  true,
		"49":  true,
		"50":  true,
		"51":  true,
		"52":  true,
		"53":  true,
		"54":  true,
		"55":  true,
		"56":  true,
		"57":  true,
		"58":  true,
		"59":  true,
		"60":  true,
		"61":  true,
		"62":  true,
		"63":  true,
		"64":  true,
		"65":  true,
		"66":  true,
		"67":  true,
		"68":  true,
		"69":  true,
		"70":  true,
		"71":  true,
		"72":  true,
		"73":  true,
		"74":  true,
		"75":  true,
		"76":  true,
		"77":  true,
		"78":  true,
		"79":  true,
		"80":  true,
		"81":  true,
		"82":  true,
		"83":  true,
		"84":  true,
		"85":  true,
		"86":  true,
		"87":  true,
		"88":  true,
		"89":  true,
		"90":  true,
		"91":  true,
		"92":  true,
		"93":  true,
		"94":  true,
		"95":  true,
		"96":  true,
		"97":  true,
		"98":  true,
		"99":  true,
		"100": true,
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

	for key, val := range numbers {
		if key == a {
			a = val
			break
		}
	}
	for key, val := range numbers {
		if key == c {
			c = val
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
			str = strconv.Itoa(res)
			break
		}

	}

	for key, val := range numbers {
		if val == str {
			str = key
			fmt.Print(str)
			break
		}
	}

}
