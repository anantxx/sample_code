/************* SetFinalizer ************************
func SetFinalizer(obj interface{}, finalizer interface{})
Finalizer is basically a function which will be called when your object will lost all references
and will be found by GC. In Go you can add finalizers to your objects with runtime.SetFinalizer function.
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

type Foo struct {
	name string
	num  int
}

func finalizer(f *Foo) {
	fmt.Println("a finalizer has run for ", f.name, f.num)
}

var counter int

func MakeFoo(name string) (a_foo *Foo) {
	a_foo = &Foo{name, counter}
	counter++
	runtime.SetFinalizer(a_foo, finalizer)
	return
}

func Bar() {
	f1 := MakeFoo("one")
	f2 := MakeFoo("two")

	fmt.Println("f1 is: ", f1.name)
	fmt.Println("f2 is: ", f2.name)
}

func main() {
	for i := 0; i < 3; i++ {
		Bar()
		time.Sleep(time.Second)
		// Run Garbage Collection
		runtime.GC()
	}
	fmt.Println("done.")
}
