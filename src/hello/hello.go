// package main
// import "fmt"
// func main() {
//     fmt.Println("hello world")
//     fmt.Println("1 +  1 = ", 1+1)
//     fmt.Println(!true)
//     var e int
//     fmt.Println(e)
//     f := "short"
//     fmt.Println(f)
// }

// var x int = 5
// func f() {
//   fmt.Println(x)
// }
//   func main() {
//   f()
// }

// func f() (int, int) {
//   return 5, 6
// }

// func main() {
//   x, y := f()
// }

// func add(args ...int) int {
//   total := 0
//   for _, v := range args {
//     total += v
//   }
//   return total
// }
// func main() {
//   fmt.Println(add(1,2,3))
// }


// iota and pointers
package main

import "fmt"

var (
    i int = 9
    hello string = "Hello world"
    pi float32 = 3.14
    c complex64 = 3+5i
)

func main() {
    fmt.Println("Hexadecimal address of i is: ", &i)
    fmt.Println("Hexadecimal address of hello is: ", &hello)
    fmt.Println("Hexadecimal address of pi is: ", &pi)
    fmt.Println("Hexadecimal address of c is: ", &c)
}


























