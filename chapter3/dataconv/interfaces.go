package dataconv

import "fmt"

func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("It's a string!")
	case int:
		fmt.Println("It's an int!")
	case rune:
		fmt.Println("It's an rune")
	case float64:
		fmt.Println("It's an float64!")
	default:
		fmt.Println("not sure what is it...")
	}
}

func Iterfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)
	CheckType(1.23)
	CheckType(rune(1))

	var i interface{}

	i = "test"
	if val, ok := i.(string); ok {
		fmt.Println("val is", val)
	}

	i = "test"
	if _, ok := i.(int); !ok {
		fmt.Println("uh oh! glad we handled this")
	}

}
