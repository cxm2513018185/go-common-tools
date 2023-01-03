package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d) // ParseDuration从字符串中解析出持续时间，其支持的有效单位有ns、us（或μs）、ms、s、m和h，例如，”300ms“, ”-1.5h“ or "2h18m"
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), nil
}
