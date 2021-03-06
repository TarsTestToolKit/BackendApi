package stats

import (
	"context"
	"time"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/TarsTestToolKit/BackendApi/client/resfetcher"
	"github.com/TarsTestToolKit/BackendApi/models/mysql"
)

// CollectStats 采集资源使用信息
func CollectStats(testID uint32) error {
	cores, mem, err := resfetcher.FetchResInfo()
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	cpuStats := make([]mysql.CpuStats, 0)
	for _, core := range cores {
		c := mysql.CpuStats{
			TestID:     int(testID),
			Total:      core.Total,
			Idle:       core.Idle,
			Used:       core.Used,
			CreateTime: now,
		}
		cpuStats = append(cpuStats, c)
	}

	memStats := mysql.MemStats{
		TestID:     int(testID),
		Total:      mem.Total,
		Used:       mem.Used,
		Cached:     mem.Cached,
		Free:       mem.Free,
		Active:     mem.Active,
		Inactive:   mem.Inactive,
		SwapTotal:  mem.SwapTotal,
		SwapUsed:   mem.SwapUsed,
		SwapFree:   mem.SwapFree,
		CreateTime: now,
	}

	return mysql.StoreStats(cpuStats, memStats)
}

func WatchStats(ctx context.Context, testID uint32, endTime time.Time) {
	go doWatchStats(ctx, testID, endTime)
}

func doWatchStats(ctx context.Context, testID uint32, endTime time.Time) {
	round := (endTime.Unix()-time.Now().Unix())/5 + 1
	ticker := time.NewTicker(time.Second * 5)
	for i := 0; i < int(round); i++ {
		select {
		case <-ctx.Done():
			return
		default:
			err := CollectStats(testID)
			if err != nil {
				tars.GetLogger("").Errorf("failed to get server stats:%s", err.Error())
			}
		}
		<-ticker.C
	}
}
