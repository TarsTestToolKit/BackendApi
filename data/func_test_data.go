package data

import "time"

// FuncTestResult 单语言功能测试结果
type FuncTestResult struct {
	Lang  string
	Ret   error
	Start time.Time
	End   time.Time
}
