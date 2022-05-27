package main

import "fmt"

func Test() {

	const (
		c1 = iota //c1=0
		c2 = iota //c2=1
		c3 = iota //c3=2
	)

	fmt.Println("c1 = ", c1, " c2 = ", c2, " c3 = ", c3)

	const (
		a  = 1 << iota //a = 1
		b  = 1 << iota //b = 2
		c  = 1 << iota //c = 4
		d  = 4 << iota
		e2 = 120 >> iota //iota = 4  e2 = 120 / 2的4次方
		e  = 2 << iota

		f = iota >> 2
		// f = (iota * 10) >> 1
	)
	fmt.Println("a = ", a, " b = ", b, " c = ", c, "d=", d, " e=", e, f)

	const (
		v1 = iota //v1 = 0
		v2        //v2 = 1
		v3        //v3 = 2
	)
	fmt.Println("v1 = ", v1, " v2 = ", v2, " v3 = ", v3)

	const (
		x = 1 << iota //x = 1
		y             //y = 2
		z             //z = 4
	)
	fmt.Println("x = ", x, " y = ", y, " z = ", z)
	const (
		Ldate = 1 << iota
		Ltime
		Lmicroseconds
		Llongfile
		Lshortfile
		LUTC
		Lmsgprefix
		LstdFlags = Ldate | Ltime
	)
	fmt.Println("Ldate", Ldate, "Ltime", Ltime, "Lmicroseconds", Lmicroseconds, "Llongfile", Llongfile, "Lshortfile", Lshortfile, "LUTC", LUTC, "Lmsgprefix", Lmsgprefix, "LstdFlags", LstdFlags)

}

func main() {
	Test()

	const (
		a2 = 1 << iota   //iota = 0  a2 = 1 x 2的0次方
		b2 = 3 << iota   //iota = 1  b2 = 3 x 2的1次方
		c2               //iota = 2  c2 = 3 x 2的2次方
		d2               //iota = 3  d2 = 3 x 2的3次方
		e2 = 120 >> iota //iota = 4  e2 = 120 / 2的4次方
		f2               //iota = 5  f2 = 120 / 2的5次方
	)

	fmt.Println(a2, b2, c2, d2, e2, f2)
	//1 6 12 24 7 3

}
