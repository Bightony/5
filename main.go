package main
import (
    "fmt"
    "5/strutil" 
    "5/numutil"
    "github.com/google/uuid"
    "github.com/fatih/color"
)
func main() {
    s := "ghjYUInnb"
    fmt.Println("Задание 2")
    fmt.Println("Оригинал:", s)
    fmt.Println("Переворот:", strutil.Reverse(s))
    fmt.Println("Верхний регистр:", strutil.ToUpper(s))
    fmt.Println("Задание 3")
    var n int
    var f int
    fmt.Print("Введите число для проверки на простоту: ")
    fmt.Scan(&n)
    if numutil.IsPrime(n) {
    fmt.Printf("%d — простое число\n", n)
    } else {
    fmt.Printf("%d — не простое число\n", n)
    }
    fmt.Print("Введите число для вычисления факториала: ")
    fmt.Scan(&f)
    fmt.Printf("%d! = %d\n", f, numutil.Factorial(f))
    fmt.Println("Задание 4")
    id := uuid.New()
    fmt.Println("UUID:", id.String())
    color.Blue("Синий")
}
