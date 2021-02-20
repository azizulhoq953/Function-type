package main

import "fmt"

func update(a*int,t*string){

	*a=*a+5
	*t=*t+"func"
	return


}

func main(){

	number:=10
	name:="Azizul"
	update(&number,&name)
	fmt.Println(name,number)

}