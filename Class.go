package main

import "fmt"

type Human struct{
	name string
	sex string
}

func (this *Human) Eat() {				//结构体和方法绑定实现类
	fmt.Println("eat！！！")
	
}

func (this *Human) Sleep() {
	fmt.Println("sleep！！！")
}


type SuperMan struct{
	Human	//父类的继承
	age int
}

func (this *SuperMan) Run() {		//子类的新方法
	fmt.Println("run！！！")
}

func main() {
	a := Human{"Cjk","male"}
	a.Eat()
	a.Sleep()
	b := SuperMan{Human{"Jack","man"},22}
	b.Eat()
	b.Run()
}