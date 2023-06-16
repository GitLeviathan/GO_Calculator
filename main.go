package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "bufio"
)

func main() {
	enter := strings.Split(Scan(), " ")
    
    flag_a, flag_b := 1, 1
    a, b := 0, 0
    act := ""
    
    if len(enter) != 3 {
        fmt.Println("Нужно ввести через пробел первое чилсо, действие, второе чилсо")
        flag_a, flag_b = 0, 0
    } else {
        tmp, err_a := strconv.Atoi(enter[0])
    	if err_a != nil {
    	    flag_a = 2
    	    tmp = Roman_to_arabian(enter[0])
    	}
    	a = tmp
    	tmp, err_b := strconv.Atoi(enter[2])
    	if err_b != nil {
    	    flag_b = 2
    		tmp = Roman_to_arabian(enter[2])
    	}
    	b = tmp
    	act = enter[1]
    }
    
    Calculation(a, b, act, flag_a, flag_b)
}

func Scan() string {
  in := bufio.NewScanner(os.Stdin)
  in.Scan()
  if err := in.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
  }
  return in.Text()
}

func Calculation (a int, b int, act string, flag_a int, flag_b int) {
    result := 0
    if flag_a == flag_b && (a > 0 && a <= 10) && (b > 0 && b <= 10) {
        switch act {
            case "+":
                result = a+b
            case "-":
                if a <= b && flag_a == 2 {
                    fmt.Println("Римляне не использовали отрицательные значения и ноль")
                } else {
                result = a-b
                }
            case "*":
                result = a*b
            case "/":
                result = a/b
            default:
                fmt.Println("Использовать можно только операции сложения, вычитания, умножения и деления")
            }
        } else {
            if flag_a != flag_b {
                fmt.Println("Цифры должны быть целыми и, либо арабскими, либо римскими, смешивать - нельзя")
            }
            if (a < 0 || a > 10) || (b < 0 || b > 10) {
                fmt.Println("На вход принимаются только числа от 1 до 10 включительно")
            }
    }
    
    if flag_a == 1 {
        fmt.Println(result)
    } else {
        fmt.Println(Arabian_to_roman(result))
    }
}

func Roman_to_arabian(tmp string) int {
    rom := strings.Split(tmp, "")

    var roman = map[string]int {
	    "I": 1, "V": 5, "X": 10, "L": 50, "C": 100,
    }

    parc := make([]int, len(rom))
    for i := 0; i < len(rom); i++ {
        parc[i] = roman[rom[i]]
    }
    
    arab := 0
    mem := parc[len(parc)-1]
    for i := len(rom) - 1; i >= 0; i-- {
        if parc[i] >= mem {
            arab += parc[i]
        } else {
            arab -= parc[i]
        }
    }
    
	return arab
}

func Arabian_to_roman(tmp int) string {
    arabian := []int {100, 90, 50, 40, 10, 9, 5, 4, 1}
    var arabian_char = map[int]string {
	    100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
    }
    
    result := ""
    for i := 0; tmp != 0; {
        if tmp >= arabian[i] {
            tmp -= arabian[i]
            result += arabian_char[arabian[i]]
        } else {
            i++
        }
    }

    return result
}
