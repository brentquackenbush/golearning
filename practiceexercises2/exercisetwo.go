package practiceexercises2

/** Hands-on exercise #2
 * Using the following operators, write expressions and assign their values to variables:
 * g. == h. <= i. >= j. != k. < l. >
 * Now print each of the variables.
 */
import "fmt"

func main() {
	i := 3
	j := 3
	if i == j {
		fmt.Println("Call options are the best. GO BULLS")
	} else if i <= j {
		fmt.Println("Put options are the best. GO BEARS")
	} else {
		fmt.Println("Too easy to finish this exercise...")
	}
}
