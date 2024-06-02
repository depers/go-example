package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237,
		time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now)) // then在now之前
	p(then.After(now))  // then在now之后
	p(then.Equal(now))  // then和now是相同的

	diff := now.Sub(then) // 两个时间的差
	p(diff)

	// 将时间差距换算成小时
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 时间加差距
	p(then.Add(diff))
	p(then.Add(-diff))
}
