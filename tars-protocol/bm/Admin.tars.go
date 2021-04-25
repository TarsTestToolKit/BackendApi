// Package bm comment
// This file was generated by tars2go 1.1.4
// Generated from BTAdmin.tars
package bm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/protocol/tup"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8
var _ = unsafe.Pointer(nil)
var _ = bytes.ErrTooLarge

//Admin struct
type Admin struct {
	s m.Servant
}

//Startup is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) Startup(req *BenchmarkUnit, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "startup", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//StartupWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) StartupWithContext(tarsCtx context.Context, req *BenchmarkUnit, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "startup", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//StartupOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) StartupOneWayWithContext(tarsCtx context.Context, req *BenchmarkUnit, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "startup", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Query is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) Query(req *BenchmarkUnit, stat *ResultStat, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*stat).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "query", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*stat).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//QueryWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) QueryWithContext(tarsCtx context.Context, req *BenchmarkUnit, stat *ResultStat, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*stat).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "query", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*stat).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//QueryOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) QueryOneWayWithContext(tarsCtx context.Context, req *BenchmarkUnit, stat *ResultStat, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*stat).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "query", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Shutdown is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) Shutdown(req *BenchmarkUnit, stat *ResultStat, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*stat).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "shutdown", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*stat).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//ShutdownWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) ShutdownWithContext(tarsCtx context.Context, req *BenchmarkUnit, stat *ResultStat, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*stat).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "shutdown", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*stat).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//ShutdownOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) ShutdownOneWayWithContext(tarsCtx context.Context, req *BenchmarkUnit, stat *ResultStat, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*stat).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "shutdown", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Test is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) Test(req *BenchmarkUnit, rsp *string, errmsg *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*errmsg), 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "test", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*rsp), 2, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*errmsg), 3, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//TestWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) TestWithContext(tarsCtx context.Context, req *BenchmarkUnit, rsp *string, errmsg *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*errmsg), 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "test", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*rsp), 2, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*errmsg), 3, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//TestOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Admin) TestOneWayWithContext(tarsCtx context.Context, req *BenchmarkUnit, rsp *string, errmsg *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*errmsg), 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "test", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *Admin) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *Admin) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol sets the protocol for the servant.
func (_obj *Admin) TarsSetProtocol(p m.Protocol) {
	_obj.s.TarsSetProtocol(p)
}

//AddServant adds servant  for the service.
func (_obj *Admin) AddServant(imp _impAdmin, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *Admin) AddServantWithContext(imp _impAdminWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impAdmin interface {
	Startup(req *BenchmarkUnit) (ret int32, err error)
	Query(req *BenchmarkUnit, stat *ResultStat) (ret int32, err error)
	Shutdown(req *BenchmarkUnit, stat *ResultStat) (ret int32, err error)
	Test(req *BenchmarkUnit, rsp *string, errmsg *string) (ret int32, err error)
}
type _impAdminWithContext interface {
	Startup(tarsCtx context.Context, req *BenchmarkUnit) (ret int32, err error)
	Query(tarsCtx context.Context, req *BenchmarkUnit, stat *ResultStat) (ret int32, err error)
	Shutdown(tarsCtx context.Context, req *BenchmarkUnit, stat *ResultStat) (ret int32, err error)
	Test(tarsCtx context.Context, req *BenchmarkUnit, rsp *string, errmsg *string) (ret int32, err error)
}

// Dispatch is used to call the server side implemnet for the method defined in the tars file. _withContext shows using context or not.
func (_obj *Admin) Dispatch(tarsCtx context.Context, _val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, _withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	_os := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "startup":
		var req BenchmarkUnit

		if tarsReq.IVersion == basef.TARSVERSION {

			err = req.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			_decoder_ := json.NewDecoder(bytes.NewReader(_is.ToBytes()))
			_decoder_.UseNumber()
			err = _decoder_.Decode(&_jsonDat_)
			if err != nil {
				return fmt.Errorf("Decode reqpacket failed, error: %+v", err)
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impAdmin)
			_funRet_, err = _imp.Startup(&req)
		} else {
			_imp := _val.(_impAdminWithContext)
			_funRet_, err = _imp.Startup(tarsCtx, &req)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "query":
		var req BenchmarkUnit
		var stat ResultStat

		if tarsReq.IVersion == basef.TARSVERSION {

			err = req.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			_decoder_ := json.NewDecoder(bytes.NewReader(_is.ToBytes()))
			_decoder_.UseNumber()
			err = _decoder_.Decode(&_jsonDat_)
			if err != nil {
				return fmt.Errorf("Decode reqpacket failed, error: %+v", err)
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impAdmin)
			_funRet_, err = _imp.Query(&req, &stat)
		} else {
			_imp := _val.(_impAdminWithContext)
			_funRet_, err = _imp.Query(tarsCtx, &req, &stat)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = stat.WriteBlock(_os, 2)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = stat.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("stat", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_
			_rspJson_["stat"] = stat

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "shutdown":
		var req BenchmarkUnit
		var stat ResultStat

		if tarsReq.IVersion == basef.TARSVERSION {

			err = req.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			_decoder_ := json.NewDecoder(bytes.NewReader(_is.ToBytes()))
			_decoder_.UseNumber()
			err = _decoder_.Decode(&_jsonDat_)
			if err != nil {
				return fmt.Errorf("Decode reqpacket failed, error: %+v", err)
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impAdmin)
			_funRet_, err = _imp.Shutdown(&req, &stat)
		} else {
			_imp := _val.(_impAdminWithContext)
			_funRet_, err = _imp.Shutdown(tarsCtx, &req, &stat)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = stat.WriteBlock(_os, 2)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = stat.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("stat", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_
			_rspJson_["stat"] = stat

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "test":
		var req BenchmarkUnit
		var rsp string
		var errmsg string

		if tarsReq.IVersion == basef.TARSVERSION {

			err = req.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			_decoder_ := json.NewDecoder(bytes.NewReader(_is.ToBytes()))
			_decoder_.UseNumber()
			err = _decoder_.Decode(&_jsonDat_)
			if err != nil {
				return fmt.Errorf("Decode reqpacket failed, error: %+v", err)
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impAdmin)
			_funRet_, err = _imp.Test(&req, &rsp, &errmsg)
		} else {
			_imp := _val.(_impAdminWithContext)
			_funRet_, err = _imp.Test(tarsCtx, &req, &rsp, &errmsg)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = _os.Write_string(rsp, 2)
			if err != nil {
				return err
			}

			err = _os.Write_string(errmsg, 3)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = _os.Write_string(rsp, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("rsp", _os.ToBytes())

			_os.Reset()
			err = _os.Write_string(errmsg, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("errmsg", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_
			_rspJson_["rsp"] = rsp
			_rspJson_["errmsg"] = errmsg

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(tarsCtx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(tarsCtx)
	if ok && c != nil {
		_context = c
	}
	*tarsResp = requestf.ResponsePacket{
		IVersion:     tarsReq.IVersion,
		CPacketType:  0,
		IRequestId:   tarsReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}

	_ = _is
	_ = _os
	_ = length
	_ = have
	_ = ty
	return nil
}
