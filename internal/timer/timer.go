package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d) // 最好对输入的时间进行处理
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), nil
}
