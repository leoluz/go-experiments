package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

type SomeInterface interface {
	SomeMethod()
}

type Impl struct{}

func (i *Impl) SomeMethod() {
	fmt.Println("hi")
}

func main() {
	maxlen := 0
	for _, v := range names {
		if lenth := len(v); lenth > maxlen {
			maxlen = lenth
		}
	}
	groups := make([][]string, maxlen)

	for _, name := range names {
		idx := len(name) - 1
		groups[idx] = append(groups[idx], name)
	}

	fmt.Println(groups)

	var sliceTest []SomeInterface
	i := &Impl{}
	sliceTest = append(sliceTest, i)
	//sliceTest = append(sliceTest, nil)
	for _, s := range sliceTest {
		s.SomeMethod()
	}

}
