package impl

import (
	"context"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsTestToolKit/BackendApi/constants/errors"
	"github.com/TarsTestToolKit/BackendApi/services/functest"
	"github.com/TarsTestToolKit/BackendApi/services/perftest"
	"github.com/TarsTestToolKit/BackendApi/tars-protocol/apitars"
)

// APIImpl servant implementation
type APIImpl struct {
}

func (imp *APIImpl) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Init servant init
func (imp *APIImpl) InitFramework(ctx context.Context) (ret apitars.SimpleResp, err error) {
	//initialize servant here:
	//...
	return
}

func (imp *APIImpl) DoPerfTest(tarsCtx context.Context, req *apitars.PerfTestReq) (ret apitars.PerfTestResp, err error) {
	tars.GetLogger("").Debugf("request DoPerfTest:%+v", *req)
	return perftest.DoPerfTest(tarsCtx, req)
}

func (imp *APIImpl) DoFuncTest(tarsCtx context.Context) (ret apitars.FuncTestResp, err error) {
	tars.GetLogger("").Debugf("request DoFuncTest")
	return functest.DoFuncTest(tarsCtx)
}

func (imp *APIImpl) GetTestDetail(tarsCtx context.Context, testID uint32, timestamp uint32, showWarmUp bool) (
	ret apitars.TestDetailResp, err error) {
	tars.GetLogger("").Debugf("request GetTestDetail id:%v time:%v showWarmUp:%v", testID, timestamp, showWarmUp)
	var finished = false
	finished, ret.PerfDetail, ret.ResUsage, err = perftest.GetTestDetail(tarsCtx, testID, timestamp, showWarmUp)
	if err != nil {
		ret.Code = uint32(tars.GetErrorCode(err))
		ret.Msg = err.Error()
		return ret, err
	}

	ret.Code = 0
	if finished == false {
		// code 1 表示running
		ret.Code = 1
	}
	ret.Msg = "succ"
	return ret, err
}

func (imp *APIImpl) GetTestHistories(tarsCtx context.Context, req *apitars.QueryTestHistoryReq) (ret apitars.QueryTestHistoryResp, err error) {
	tars.GetLogger("").Debugf("request GetTestHistories req:%+v", *req)
	total, perfTests, err := perftest.QueryHistories(tarsCtx, req.Page, req.PageSize)
	ret.Total = uint32(total)
	ret.Page = req.Page
	ret.Histories = perfTests

	return ret, err
}

func (imp *APIImpl) IsPerfExists(tarsCtx context.Context, req *apitars.IsPerfExistsReq) (ret apitars.IsPerfExistsResp, err error) {
	tars.GetLogger("").Debugf("request IsPerfExists req:%+v", *req)
	exists, err := perftest.IsPerfExists(tarsCtx, req.ServType)
	if err != nil {
		ret.Code = uint32(tars.GetErrorCode(err))
		ret.Msg = err.Error()
		return
	}
	if exists {
		ret.Code = errors.ErrCodeDuplicatePerfTest
		ret.Msg = "Duplicate Performance Test"
		return
	}

	return
}
