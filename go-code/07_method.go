package main

import "fmt"

type Person1 struct {
	Name string
	Age int
}

func (q Person1) getName() string {
	return q.Name
}

func (q *Person1)setName(Name string)  {
	q.Name = Name
}

func main() {

	p := new(Person1)
	//p.Name = "keen"
	p.setName("keen")
	p.Age= 18
	fmt.Println(p.getName())  //keen

}
