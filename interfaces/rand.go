package main

import (
	"fmt"
	"math/rand"
	"time"
)

type customRand struct {
	*rand.Rand
	count int
}

func NewCustomRand(i int64) *customRand {
	return &customRand{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (cr *customRand) RangeRand(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max-min) + min
}

func (cr *customRand) GetCount() int {
	return cr.count
}

/* If the outer type implements a method that has the same name as a method implemented by the inner type,
   then the outer method will be given priority. */
func (cr *customRand) Intn(n int) int {
	cr.count++
	return cr.Rand.Intn(n) + 1
}

func main() {
	cr := NewCustomRand(time.Now().UnixNano())
	fmt.Println(cr.RangeRand(5, 30))
	fmt.Println(cr.Intn(10))
	fmt.Println(cr.GetCount())
	/* Outer type stays as a different type than the inner type,
	   meaning that a value of the inner type can not be assigned to the outer type or vice versa.
	   var r *rand.Rand = cr
	   => Error: Cannot use cr (type *customRand) as type *rand.Rand in assignment.
	   cr = &rand.Rand
	   => Error: Cannot use literal (type *rand.Rand) as type *customRand in assignment. */
}
