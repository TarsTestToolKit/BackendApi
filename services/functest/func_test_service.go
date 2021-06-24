package functest

import (
	"context"
	"sync"
	"time"

	"github.com/TarsTestToolKit/BackendApi/client/benchmark"
	"github.com/TarsTestToolKit/BackendApi/data"
	"github.com/TarsTestToolKit/BackendApi/tars-protocol/apitars"
)

// DoFuncTest 执行功能性测试
func DoFuncTest(ctx context.Context) (apitars.FuncTestResp, error) {
	start := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(5)
	ch := make(chan *data.FuncTestResult, 5)

	go pingCpp(ctx, wg, ch)
	go pingJava(ctx, wg, ch)
	go pingGo(ctx, wg, ch)
	go pingNodeJs(ctx, wg, ch)
	go pingPhp(ctx, wg, ch)
	wg.Wait()
	close(ch)
	resp := apitars.FuncTestResp{
		Code:      0,
		Msg:       "succ",
		Rows:      make([]apitars.FuncTestDetail, 0),
		StartTime: start.UnixNano() / int64(time.Millisecond),
	}
	for result := range ch {
		row := apitars.FuncTestDetail{
			From:      "cpp",
			To:        result.Lang,
			IsSucc:    result.Ret == nil,
			StartTime: start.UnixNano() / int64(time.Millisecond),
			EndTime:   time.Now().UnixNano() / int64(time.Millisecond),
		}
		resp.Rows = append(resp.Rows, row)
	}
	resp.EndTime = time.Now().UnixNano() / int64(time.Millisecond)

	return resp, nil
}

func pingCpp(ctx context.Context, wg *sync.WaitGroup, ch chan<- *data.FuncTestResult) {
	defer wg.Done()
	ch <- &data.FuncTestResult{
		Lang: "cpp",
		Ret:  benchmark.PingCpp(ctx),
	}
}

func pingJava(ctx context.Context, wg *sync.WaitGroup, ch chan<- *data.FuncTestResult) {
	defer wg.Done()
	ch <- &data.FuncTestResult{
		Lang: "java",
		Ret:  benchmark.PingJava(ctx),
	}
}

func pingGo(ctx context.Context, wg *sync.WaitGroup, ch chan<- *data.FuncTestResult) {
	defer wg.Done()
	ch <- &data.FuncTestResult{
		Lang: "golang",
		Ret:  benchmark.PingGo(ctx),
	}
}

func pingNodeJs(ctx context.Context, wg *sync.WaitGroup, ch chan<- *data.FuncTestResult) {
	defer wg.Done()
	start := time.Now()
	ch <- &data.FuncTestResult{
		Lang:  "nodejs",
		Ret:   benchmark.PingNodeJs(ctx),
		Start: start,
		End:   time.Now(),
	}
}

func pingPhp(ctx context.Context, wg *sync.WaitGroup, ch chan<- *data.FuncTestResult) {
	defer wg.Done()
	start := time.Now()
	ch <- &data.FuncTestResult{
		Lang:  "php",
		Ret:   benchmark.PingPhp(ctx),
		Start: start,
		End:   time.Now(),
	}
}
