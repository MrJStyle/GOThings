package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	age  int
}

func (u *user) notify() {
	fmt.Println(u.name, u.age)
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Println(a.name, a.age, a.level)
}

func doSomething(n notifier) {
	n.notify()
}

func main() {
	a := admin{
		user:  user{"boy", 10},
		level: "top",
	}

	a.user.notify()

	// 此处若admin没有实现notify方法， 会提升user的notify方法来调用
	a.notify()

	doSomething(&a)
}
