// required package 'main' for every golang project
package main

import (
	"fmt"
	"strconv"
)

type Boss struct {
	name        string
	age         int
	vision      string
	geniusLevel string
}

func insertArray(a []int, i int) []int {
	a = append(a, i)
	return a
}

func main() {

	type human struct {
		name  string
		skill string
	}

	init := make([]human, 0) // slice of size 2
	hs := []human{{"superman", "flying"}}

	for _, insert := range hs {
		init = append(init, human{name: insert.name, skill: insert.skill})
	}

	for key, val := range init {
		fmt.Printf("%v » %v", key, val)
	}

	firstArray := []int{}
	numberOfInsertions := 4
	// slice1 := make([]int, 2, 2) // » slice of type integer which have length = 2 and capacity = 2, declared using built-in function ¯make¯
	for i := 0; i < numberOfInsertions; i++ {
		firstArray = insertArray(firstArray, i)
	}
	for index, element := range firstArray {
		fmt.Printf("index[%d] » %v\n", index, element)
	}
	fmt.Printf("\n")
	slice := firstArray[1:]
	fmt.Printf("capacity of the ¯slice¯ is %v\n", cap(slice))
	for i := range slice {
		fmt.Printf("%+v\n", i)
	}

	grade := 79
	if !(grade >= 50) {
		fmt.Printf("failed")
	} else if grade >= 51 && grade <= 80 {
		fmt.Printf("ordinary pass")
	} else {
		fmt.Printf("distinguished pass")
	}

	number := 10
	fmt.Printf("%v, type: %T\n", number, number)

	// format int into string
	fmt.Printf("%v, type: %T\n", number, strconv.Itoa(number))

	// format int into float64
	fmt.Printf("%v, type: %T\n", float64(number), float64(number))

	var myvar1, myvar2, myvar3 = myfunc(2, 4)

	// Display the values
	fmt.Printf("Result is: %d", myvar1)
	fmt.Printf("\nResult is: %d", myvar2)
	fmt.Printf("\nResult is: %d\n", myvar3)

	var a, b, c = returnMultiple(10, 10)
	fmt.Printf("%+v, %+v, %+v\n", a, b, c)

	var me Boss
	me.name = "Kalby"
	me.age = 26
	me.vision = "Better life"
	me.geniusLevel = "more than you know"
	printStruct(me)

	var ab int
	fmt.Printf("%d, %v, %T\n\n", ab, ab, ab)

	var flag bool
	flag = true
	for flag == true {
		fmt.Printf("flag: %v\n", flag)
		flag = false
		if flag == false {
			fmt.Printf("flag has been set to %v", flag)
		}
	}

	numbers := []int{1, 2, 3, 4}
	for i := range numbers {
		fmt.Printf("range val: %v\n", i)
	}

	// go map
	goMap := make(map[string]string)
	goMap["key1"] = "value1"
	goMap["key2"] = "value2"

	goMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	// print map
	fmt.Printf("%v\n", goMap)
	fmt.Printf("%v\n", goMap2)

	// print value in the map
	fmt.Printf("%v\n", goMap["key1"])
	fmt.Printf("%v\n", goMap2["key1"])

	// loop through the map by accessing its keys and values
	for key, val := range goMap2 {
		fmt.Printf("key: %v\tvalue: %v\n", key, val)
	}

	// loop through the map by accessing only its values
	for _, val := range goMap2 {
		fmt.Printf("value: %v\n", val)
	}
	// go map

}

// multiple return values with the given input -
// function parameters
func myfunc(p, q int) (int, int, int) {
	return p - q, p * q, p + q
}

func returnMultiple(a int, b int) (int, int, bool) {
	flag := true
	if a-b == 0 {
		fmt.Printf("%+v", flag)
	} else {
		flag = false
	}
	return a + b, a - b, flag
}

func printStruct(me Boss) {

	fmt.Printf("name: %s\n", me.name)
	fmt.Printf("age: %d\n", me.age)
	fmt.Printf("vision: %s\n", me.vision)
	fmt.Printf("genius level: %s\n", me.geniusLevel)
}
