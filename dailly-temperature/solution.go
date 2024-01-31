package daily_temperature

func DailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	}
	result := make([]int, len(temperatures))
	i := 0
	for i < len(temperatures) {
		result[i] = 0
		comparable := temperatures[i]
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > comparable {
				result[i] = j - i
				break
			}
		}
		i++
	}
	return result
}

func DailyTemperaturesStack(temperatures []int) []int {
	var stack []int
	ans := make([]int, len(temperatures), len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = i - j
		}
		stack = append(stack, i)
	}
	return ans
}
