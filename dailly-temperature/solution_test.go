package daily_temperature_test

import (
	"fmt"
	"testing"

	daily_temperature "leetcode/dailly-temperature"
)

func TestDo(t *testing.T) {
	x := daily_temperature.DailyTemperaturesStack([]int{73,74,75,71,69,72,76,73})
	fmt.Println(x)
}
