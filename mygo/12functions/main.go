package main

import "fmt"

var localslice = make([]int, 3)

type User struct {
	Name  string
	Email string
	age   int
}

func main() {
	superman()
	mul := multiplyme(4, 2)
	fmt.Println("mul value is", mul)

	sum, length, name := adder(1, 2, 3, 4, 55, 6)
	fmt.Println("sum value is ", sum)
	fmt.Println("length is ", length)
	fmt.Println("name is ", name)

	//testing pointer to int with func
	intval := 3
	intp := &intval
	fmt.Println("before passing intval to the func", intval)
	manipulateIntVal(intval)
	fmt.Println("after passing intval to the func", intval)

	fmt.Println("before passing intp to the func", *intp)
	manipulateIntP(intp)
	fmt.Println("after passing intp to the func", *intp)
	fmt.Println("after passing intval to the func", intval)

	fmt.Println("before change passing intp to the func", *intp)
	manipulateIntPChangePointer(intp)
	fmt.Println("after change passing intp to the func", *intp)
	fmt.Println("after change passing intval to the func", intval)
	//testing array with func
	var myarr [3]int
	for i := range myarr {
		myarr[i] = i
	}
	fmt.Println("before passing array to the func", myarr)
	manipulateArray(myarr)
	fmt.Println("after passing array to the func", myarr)

	//testing pointer to array with func
	arrp := &myarr
	fmt.Println("before passing pointer to array to the func", arrp)
	manipulatePointerToArray(arrp)
	fmt.Println("after passing pointer to array to the func", arrp)
	fmt.Println("after passing array to the func", myarr)

	//testing slices with func--because slice is basically a combination of pointer to array, length, and capacity
	var myslice []int = make([]int, 3)
	for i := range myslice {
		myslice[i] = i
	}
	fmt.Println("before passing slice to the func", myslice)
	manipulateSlice(myslice)
	fmt.Println("after passing slice to the func", myslice)

	//testing pointer to slice with func
	myslicep := &myslice
	fmt.Println("before passing pointer to slice to the func", myslicep)
	manipulatePointerToSlice(myslicep)
	fmt.Println("after passing pointer to slice to the func", myslicep)

	//testing struct with func
	vignesh := User{"Vignesh", "rvignesh12jun@gmail.com", 24}
	fmt.Println("before passing struct to the func", vignesh)
	manipulateStruct(vignesh)
	fmt.Println("after passing struct to the func", vignesh)

	//testing pointer to struct with func
	vigneshp := &vignesh
	fmt.Println("before passing pointer to struct to the func", vignesh)
	manipulateStructP(vigneshp)
	fmt.Println("after passing pointer to struct to the func", vignesh)

	fmt.Println("before passing change pointer to struct to the func", vignesh)
	manipulateStructPChangePointer(vigneshp)
	fmt.Println("after passing change pointer to struct to the func", vignesh)

}

func superman() {
	fmt.Println("I am superman")
}

func multiplyme(v1 int, v2 int) int {
	return v1 * v2
}

func adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum += values[i]
	}

	length := len(values)
	name := "this is a name"
	return sum, length, name
}

func manipulateArray(arr [3]int) {
	arr[2] = 10
}

func manipulatePointerToArray(arrp *[3]int) {
	arrp[2] = 10
}

func manipulateSlice(slice []int) {
	slice[2] = 10
}

//below operation not supported
// func manipulatePointerToSlice(slicep *[]int) {
// 	slicep[2] = 15
// }

func manipulatePointerToSlice(slicep *[]int) {
	// localslice := make([]int, 3)
	fmt.Println("localslice", localslice)
	fmt.Println("inside manipulatePointerToSlice before overwrite", slicep)
	slicep = &localslice
	fmt.Println("inside manipulatePointerToSlice after overwrite", slicep)

}

func manipulateIntVal(val int) {
	val = 10
}

func manipulateIntP(p *int) {
	*p = 20
}

func manipulateStruct(obj User) {
	obj.age = 100
}

func manipulateStructP(obj *User) {
	obj.age = 100
}

func manipulateStructPChangePointer(obj *User) {
	praveen := User{"Praveen", "rvignesh12jun@gmail.com", 24}
	obj = &praveen
	obj.age = 100
}

func manipulateIntPChangePointer(intp *int) {
	val := 100
	intp = &val
}
