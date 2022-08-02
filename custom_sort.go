package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type By func(u1, u2 *user) bool

func (by By) Sort(users []user) {
	us := &userSorter{
		users: users,
		by:    by,
	}
	sort.Sort(us)
}

type userSorter struct {
	users []user
	by    func(u1, u2 *user) bool
}

func (s *userSorter) Len() int {
	return len(s.users)
}

func (s *userSorter) Swap(i, j int) {
	s.users[i], s.users[j] = s.users[j], s.users[i]
}

func (s *userSorter) Less(i, j int) bool {
	return s.by(&s.users[i], &s.users[j])
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	age := func(u1, u2 *user) bool {
		return u1.Age < u2.Age
	}

	name := func(u1, u2 *user) bool {
		return u1.First < u2.First
	}

	By(age).Sort(users)
	fmt.Println(users)

	By(name).Sort(users)
	fmt.Println(users)
}
