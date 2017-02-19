// Package initialize provides a way to organize initializers. With this library,
// the `init` function can be simplified to one line of code at
// the cost of making each initializer function a little more complex.
//
//	type initializer struct{}
//
//	func (v initializer) Initialize01Log(){
//		...
//	}
//
//	func (v initializer) Initialize02DB(){
//		...
//	}
//
//	func (v initializer) Initialize03Redis(){
//		...
//	}
//
//	...
//	...
//
//	func init() {
//		initialize.AllFrom(initializer{})
//	}
package initialize

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
)

// AllFrom collects and runs initializers defined in an object.
func AllFrom(v interface{}) {
	if v == nil {
		return
	}

	methods := make(map[int]int)
	re := regexp.MustCompile("Initialize(\\d+)\\w+")
	val := reflect.ValueOf(v)
	typ := val.Type()

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		matches := re.FindStringSubmatch(method.Name)
		if len(matches) == 2 {
			idx, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(fmt.Sprintf("Error in initializers: %s", err))
			}
			methods[idx] = i
		}
	}

	var keys []int
	for k := range methods {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		val.Method(methods[k]).Call(nil)
	}
}
