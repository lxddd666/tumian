package test

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"testing"
	"time"
)

// GetRecentWeekday 返回最近的工作日（周一到周五），时间部分清零
// 如果当天是工作日，返回当天；否则返回上一个工作日
func GetRecentWeekday() *gtime.Time {
	now := time.Now()
	weekday := now.Weekday()

	// 计算需要回溯的天数
	daysToSubtract := 0

	switch weekday {
	case time.Saturday:
		// 周六：回溯1天到周五
		daysToSubtract = 1
	case time.Sunday:
		// 周日：回溯2天到周五
		daysToSubtract = 2
	default:
		// 周一到周五：不需要回溯
		daysToSubtract = 0
	}

	// 计算目标日期
	targetDate := now.AddDate(0, 0, -daysToSubtract)

	// 清零时间部分，只保留年月日
	dateOnly := time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(),
		0, 0, 0, 0, targetDate.Location())

	return gtime.NewFromTime(dateOnly)
}

// 更简洁的版本，使用数组查找
func GetRecentWeekdaySimple() *gtime.Time {
	now := time.Now()
	weekday := now.Weekday()

	// 定义需要回溯的天数：索引对应 time.Weekday 的值
	// Sunday=0, Monday=1, Tuesday=2, Wednesday=3, Thursday=4, Friday=5, Saturday=6
	daysBack := []int{2, 0, 0, 0, 0, 0, 1} // 周日回溯2天，周六回溯1天，其他不回溯

	// 计算目标日期
	targetDate := now.AddDate(0, 0, -daysBack[weekday])

	// 清零时间部分
	dateOnly := time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(),
		0, 0, 0, 0, targetDate.Location())

	return gtime.NewFromTime(dateOnly)
}

// 处理特定输入时间的版本
func GetRecentWeekdayForTime(inputTime time.Time) *gtime.Time {
	weekday := inputTime.Weekday()
	daysBack := []int{2, 0, 0, 0, 0, 0, 1}

	targetDate := inputTime.AddDate(0, 0, -daysBack[weekday])
	dateOnly := time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(),
		0, 0, 0, 0, targetDate.Location())

	return gtime.NewFromTime(dateOnly)
}

func TestDate(t *testing.T) {
	// 示例使用
	recentWorkday := GetRecentWeekday()
	fmt.Printf("最近的工作日: %s\n", recentWorkday.Format("Y-m-d"))
	fmt.Printf("完整时间: %s\n", recentWorkday.Format("Y-m-d H:i:s"))

	// 测试不同日期
	testDates := []time.Time{
		time.Date(2026, 1, 17, 14, 30, 0, 0, time.Local), // 周一
		time.Date(2026, 1, 18, 10, 0, 0, 0, time.Local),  // 周六
		time.Date(2024, 1, 21, 9, 0, 0, 0, time.Local),   // 周日
	}

	fmt.Println("\n测试不同日期:")
	for _, testDate := range testDates {
		result := GetRecentWeekdayForTime(testDate)
		fmt.Printf("%s (周%d) -> 最近工作日: %s\n",
			testDate.Format("2006-01-02"),
			testDate.Weekday(),
			result.Format("Y-m-d"))
	}
}
