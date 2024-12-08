package main

import "fmt"

type AnimalIF interface { //接口是对类的抽象，接口的底层是一个指针，它指向了接口的方法
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的一个类
type Cat struct {
	color string
}

func (this *Cat) Sleep() { //重写了接口的方法，将接口中的方式具体实现了
	fmt.Println("Cat is Sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的一个类
type Dog struct {
	color string
}

func (this *Dog) Sleep() { //不同的类实现同一个接口的方法，体现了接口的多态
	fmt.Println("Dog is Sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func main() {
	var animal AnimalIF //声明了一个animal变量，他的类型是AnimalIF
	animal = &Cat{"Yellow"}

	animal.Sleep()

	animal = &Dog{"Black"}
	animal.Sleep()
}

/*
Go中多态的三要素：
1.有一个父类（接口），父类中的方法只是声明了，并没有被实现。
2.有子类，子类继承了父类，并实现了父类（接口）中的所有方法。
3.父类类型的变量（他是一个指针）指向了子类的对象。

多态的具体体现：父类（接口）类型的变量指向不同的子类对象，那么这个接口类型的变量在调用方法时就调用子类中具体实现的方法
*/
