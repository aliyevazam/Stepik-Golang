package main

import "fmt"

type Person interface {
	shower()
	breakfast()
	work()
	experience()
}

type Person1 struct {
	name   string
	jobs   string
	age    int
	salary string
}

func (s Person1) shower() {
	fmt.Println(s.name, "Prinimaet dush")
}
func (s Person1) breakfast() {
	fmt.Println(s.name, "Delayet zavtrak")
}
func (s Person1) work() {
	fmt.Println(s.name, "Idyot na rabotu v Freelancing company on ", s.jobs)
}
func (s Person1) experience() {
	fmt.Println(s.name, "Polucahyet ", s.salary, " Deneg v mesyace yemu ", s.age)
}

func main() {
	a := &Person1{name: "Abdulaziz", jobs: "Developer", age: 19, salary: "500$"}
	b := &Person1{name: "Ali", jobs: "Project Manager", age: 24, salary: "7000$"}
	j := [...]Person{a, b}
	for _, j := range j {
		j.shower()
		j.breakfast()
		j.work()
		j.experience()
	}
}
