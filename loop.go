package main

import (
	"fmt"
)

func main() {
	var a,e string
	var b,c,d,f int
	a = "y"
	f = 1

	fmt.Println("Selamat mencoba aplikasi aritmatika")

    for a != "n" {
	fmt.Print("Apakah anda akan melakukan perhitungan ( y / n) = ")
	fmt.Scan(&a)

		if a == "y" {
			fmt.Println("______________________________________")
			fmt.Print("Input Operand 1 : ")
			fmt.Scan(&b)
			fmt.Print("Input Operator aritmatika : ")
			fmt.Scan(&e)
			fmt.Print("Input Operand 2 : ")
			fmt.Scan(&c)
			
			 if e == "+" {
				d = b+c
			} else if e == "-" {
				d = b-c
			} else if e == "*" {
				d = b*c
			} else if e == "/" {
				d = b/c
			} else if e == "^" {
				d  = b
				for ; f < c; {
					d = d*b
					f++
				}
			} else if e == "%" {
				d = b%c
			} else {
				fmt.Print("Operator tidak terdefinisi, mulai dari awal lagi ya. ")
			}

				fmt.Print("Hasil ")
				fmt.Print(b)
				fmt.Print(e)
				fmt.Print(c)
				fmt.Print(" = ")
				fmt.Println(d)
				fmt.Println("______________________________________")

		}else if a != "n"{
			fmt.Println("Input hanya -y- untuk YES atau -n- untuk NO")
			
		}
		
	}
fmt.Println("Bye, Terimakasih")
    
//VER 1.0.0
    }