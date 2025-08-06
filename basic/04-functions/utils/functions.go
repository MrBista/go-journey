package utils

import (
	"errors"
	"fmt"
	"go-journey/basic/helper"
)

func CallFunctionLearnSection() {
	helper.CalledFunction("CallFunctionLearnSection")

	sayHello()

	greating("Bisboy")

	valAdding := add(1, 2)

	val, err := calculate(1, 2, "/")

	fmt.Println("value=", val, "error", err, "valAdding", valAdding)

	valDevide, errDevide := devide(2, 3)

	fmt.Println("vvalDevideal", valDevide, "errDevide", errDevide)

	name, age, status := gePersonInfo()

	fmt.Println("name", name, "age", age, "status", status)

	// ========================================
	// 5. PENGGUNAAN FUNCTION AS A VALUE
	// ========================================
	// type value, mirip interface yang merupakan template aja
	var mathTypeOperarettion MathOperation = devide

	res, err := mathTypeOperarettion(2, 3)

	fmt.Println("resultnya adalah ", res, err)

	addFunc := getOperation("+")

	resAdd, errAdd := addFunc(2, 3)

	substractionFunc := getOperation("-")

	resSubs, errSubs := substractionFunc(10, 2) // curryig function

	fmt.Println("Hasil dari pengurangan 10 - 2", resSubs, errSubs)

	fmt.Println("resAdd", resAdd, "errAdd", errAdd)
}

// ========================================
// 1. DEKLARASI FUNGSI DASAR
// ========================================

// Fungsi tanpa parameter dan return value
func sayHello() {

	fmt.Println("Hello world function")
}

// fungsi dengan paramter
func greating(name string) {
	fmt.Println("Hallo ", name, " selamat datang pada pembelajaran ini")
}

// fungsi dengan return value
func add(a, b int) int {
	return a + b
}

// fungsi multiple parameter dengan return value dan error
func calculate(a, b int, operation string) (int, error) {

	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil

	default:
		return 0, errors.New("please input operation")

	}
}

// ========================================
// 2. MULTIPLE RETURN VALUES
// ========================================

// fungsi dengan multiple return values
func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

// contoh lain multiple return value

func gePersonInfo() (string, int, bool) {
	return "John Doe", 25, false
}

// ========================================
// 3. NAMED RETURN VALUES
// ========================================

// Named return values - variabel return sudah dideklarasikan
func rectangleArea(length, width float64) (area float64) {

	area = length * width
	return
}

// Named return values dengan multiple values
func circleProperties(radius float64) (area, circumference float64) {
	area = 3.14159 * radius * radius
	circumference = 2 * 3.14159 * radius
	return
}

// kombinasi named dan explicit return
func processNumber(x, y int) (sum, product int) {

	sum = x + y
	product = x * y

	return sum, product
}

// ========================================
// 4. VARIADIC FUNCTIONS
// ========================================

// variadic function - menerima jumlah parameter yang bervariasi
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total
}

// variadic dengan parameter lain
func greetMultiple(greeting string, names ...string) {
	for i, item := range names {
		fmt.Printf("item %d: %v\n", i+1, item)
	}
}

// Variadic dengan tipe interface{} untuk menerima berbagai tipe
func printAll(items ...interface{}) {
	for i, item := range items {
		fmt.Printf("Item %d: %v\n", i+1, item)
	}
}

// ========================================
// 5. FUNCTION AS A VALUE
// ========================================

// mirip interface di java

type MathOperation func(int, int) (int, error)

func addMath(a, b int) (int, error) {

	return a + b, nil
}

func divedMath(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}

	return a / b, nil
}

func substractionMath(a, b int) (int, error) {
	return a - b, nil
}

func multipleMath(a, b int) (int, error) {
	return a * b, nil
}

func getOperation(op string) MathOperation {
	switch op {
	case "+":
		return addMath
	case "-":
		return substractionMath
	case "/":
		return divedMath
	case "*":
		return multipleMath
	default:
		return nil
	}
}
