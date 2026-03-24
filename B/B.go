package B
import "5/A"
func B() string {
    return "B вызывает " + A.A()
}