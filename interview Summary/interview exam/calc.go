package main

import ("fmt")

func main(){
 var n1 int
 var n2 int
fmt.Println("Input n1")
fmt.Scanf("%d",&n1)
fmt.Println("Input n2")
fmt.Scanf("%d",&n2)
sum:=n1%n2

fmt.Println("n1 % n2 = ",sum)
}

