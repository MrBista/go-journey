package utils

import "fmt"

func checkType(value interface{}) {
	if str, ok := value.(string); ok {
		fmt.Println("Value merupakan type string berupa ", str)
	} else if intVal, ok := value.(int); ok {
		fmt.Println("Value merupakan integer berupa: ", intVal)
	}
}

func TypeAssertionLearn() {
	fmt.Println("Type assertion learn")
	checkType("Buba")
}
