package main

import "fmt"
const(
prefix = "Hello, " 
spanish = "Spanish"
french = "French"
spanishPrefix = "Hola, "
frenchPrefix = "Bonjour, "
)


func hw(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (pre string) {
	switch language {
	case french:
		pre = frenchPrefix
	case spanish:
		pre = spanishPrefix
	default:
		pre = prefix
	}
	return
}
	

	
func main(){
	fmt.Println(hw("hello world",""))
}

