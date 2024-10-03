// Add 返回多个整数的和，并以字符串形式输出。
package mathutil

import "fmt"
import "errors"

// Add 返回多个整数的和，并以字符串形式输出。如果没有提供任何数字，则返回错误。
func Add(numbers ...int) (int, string, error) {
	if len(numbers) == 0 {
		return 0, "", errors.New("没有提供任何数字")
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum, fmt.Sprintf("The sum is: %d", sum), nil
}
