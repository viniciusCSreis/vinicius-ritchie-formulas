package calc

import (
	"fmt"
	"time"
)

type WorkHours struct {
	StartWork  string
	StartLunch string
	EndLunch   string
	EndWork    string
}

func (h WorkHours) Calc() (string, error) {
	startWork, err := toDate(h.StartWork)
	if err != nil {
		return "", fmt.Errorf("fail to parse %s", h.StartWork)
	}
	startLunch, err := toDate(h.StartLunch)
	if err != nil {
		return "", fmt.Errorf("fail to parse %s", h.StartLunch)
	}
	endLunch, err := toDate(h.EndLunch)
	if err != nil {
		return "", fmt.Errorf("fail to parse %s", h.EndLunch)
	}
	endWork, err := toDate(h.EndWork)
	if err != nil {
		return "", fmt.Errorf("fail to parse %s", h.EndWork)
	}

	beforeLunch := startLunch.Unix() - startWork.Unix()
	afterLunch := endWork.Unix() - endLunch.Unix()
	result := time.Unix(beforeLunch+afterLunch, 0).UTC()
	return toString(result), nil
}

func (h WorkHours) CalcPerfectEndWork(target string) (string, error) {
	startWork, err := toDate(h.StartWork)
	if err != nil {
		return "", fmt.Errorf("fail to parse %s", h.StartWork)
	}
	startLunch, err := toDate(h.StartLunch)
	if err != nil {
		return "", fmt.Errorf("fail to parse %s", h.StartLunch)
	}
	endLunch, err := toDate(h.EndLunch)
	if err != nil {
		return "", fmt.Errorf("fail to parse %s", h.EndLunch)
	}
	targetTime, err := toDate(target)
	if err != nil {
		return "", fmt.Errorf("fail to parse target %s", target)
	}
	beforeLunch := startLunch.Unix() - startWork.Unix()
	missing := targetTime.Unix() - beforeLunch
	result := time.Unix(endLunch.Unix()+missing, 0)
	return toString(result), nil
}

func toDate(input string) (time.Time, error) {
	return time.Parse("15:04", input)
}

func toString( tt time.Time) string {
	return tt.UTC().Format("15:04")
}
