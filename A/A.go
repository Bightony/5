package A
import "5/B"
func A() string {
    return "A вызывает " + B.B()
}