// Package apitars comment
// This file was generated by tars2go 1.1.4
// Generated from api.tars
package apitars

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

//const as define in tars file
const (
	RespCodeSucc int32  = 0
	FuncTest     string = "func"
	PerfTest     string = "perf"
)

// SimpleResp struct implement
type SimpleResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func (st *SimpleResp) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *SimpleResp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_uint32(&st.Code, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Msg, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *SimpleResp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require SimpleResp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *SimpleResp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_uint32(st.Code, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Msg, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *SimpleResp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// FuncTestDetail struct implement
type FuncTestDetail struct {
	From   string `json:"from"`
	To     string `json:"to"`
	IsSucc bool   `json:"isSucc"`
}

func (st *FuncTestDetail) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *FuncTestDetail) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.From, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.To, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.IsSucc, 2, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *FuncTestDetail) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require FuncTestDetail, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *FuncTestDetail) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.From, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.To, 1)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.IsSucc, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *FuncTestDetail) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// FuncTestResp struct implement
type FuncTestResp struct {
	Code uint32           `json:"code"`
	Msg  string           `json:"msg"`
	Rows []FuncTestDetail `json:"rows"`
}

func (st *FuncTestResp) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *FuncTestResp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_uint32(&st.Code, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Msg, 1, true)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(2, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = _is.Read_int32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Rows = make([]FuncTestDetail, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = st.Rows[i0].ReadBlock(_is, 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SIMPLE_LIST {
			err = fmt.Errorf("not support simple_list type")
			if err != nil {
				return err
			}

		} else {
			err = fmt.Errorf("require vector, but not")
			if err != nil {
				return err
			}

		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *FuncTestResp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require FuncTestResp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *FuncTestResp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_uint32(st.Code, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Msg, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.Rows)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Rows {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}

	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *FuncTestResp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// PerfTestReq struct implement
type PerfTestReq struct {
	Lang       string `json:"lang"`
	ServType   string `json:"servType"`
	ThreadCnt  uint32 `json:"threadCnt"`
	Cores      uint32 `json:"cores"`
	ConnCnt    uint32 `json:"connCnt"`
	ReqFreq    uint32 `json:"reqFreq"`
	KeepAlive  uint32 `json:"keepAlive"`
	PackageLen uint32 `json:"packageLen"`
}

func (st *PerfTestReq) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *PerfTestReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.Lang, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.ServType, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.ThreadCnt, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.Cores, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.ConnCnt, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.ReqFreq, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.KeepAlive, 6, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.PackageLen, 7, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *PerfTestReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require PerfTestReq, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *PerfTestReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Lang, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.ServType, 1)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.ThreadCnt, 2)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.Cores, 3)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.ConnCnt, 4)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.ReqFreq, 5)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.KeepAlive, 6)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.PackageLen, 7)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *PerfTestReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// PerfTestResp struct implement
type PerfTestResp struct {
	Code   uint32 `json:"code"`
	Msg    string `json:"msg"`
	TestID uint32 `json:"testID"`
}

func (st *PerfTestResp) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *PerfTestResp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_uint32(&st.Code, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Msg, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.TestID, 2, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *PerfTestResp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require PerfTestResp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *PerfTestResp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_uint32(st.Code, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Msg, 1)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.TestID, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *PerfTestResp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// PerfTestDetail struct implement
type PerfTestDetail struct {
	Timestamp  uint32            `json:"timestamp"`
	Qps        uint32            `json:"qps"`
	TotalReq   uint32            `json:"totalReq"`
	Succ       uint32            `json:"succ"`
	Failed     uint32            `json:"failed"`
	SuccRate   string            `json:"succRate"`
	CostMax    float32           `json:"costMax"`
	CostMin    float32           `json:"costMin"`
	CostAvg    float32           `json:"costAvg"`
	P90        float32           `json:"p90"`
	P99        float32           `json:"p99"`
	P999       float32           `json:"p999"`
	SendByte   uint32            `json:"sendByte"`
	RecvByte   uint32            `json:"recvByte"`
	CostMap    map[string]uint32 `json:"costMap"`
	RetCodeMap map[string]uint32 `json:"retCodeMap"`
}

func (st *PerfTestDetail) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *PerfTestDetail) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_uint32(&st.Timestamp, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.Qps, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.TotalReq, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.Succ, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.Failed, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SuccRate, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_float32(&st.CostMax, 6, true)
	if err != nil {
		return err
	}

	err = _is.Read_float32(&st.CostMin, 7, true)
	if err != nil {
		return err
	}

	err = _is.Read_float32(&st.CostAvg, 8, true)
	if err != nil {
		return err
	}

	err = _is.Read_float32(&st.P90, 9, true)
	if err != nil {
		return err
	}

	err = _is.Read_float32(&st.P99, 10, true)
	if err != nil {
		return err
	}

	err = _is.Read_float32(&st.P999, 11, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.SendByte, 12, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.RecvByte, 13, true)
	if err != nil {
		return err
	}

	err, have = _is.SkipTo(codec.MAP, 14, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&length, 0, true)
	if err != nil {
		return err
	}

	st.CostMap = make(map[string]uint32)
	for i0, e0 := int32(0), length; i0 < e0; i0++ {
		var k0 string
		var v0 uint32

		err = _is.Read_string(&k0, 0, false)
		if err != nil {
			return err
		}

		err = _is.Read_uint32(&v0, 1, false)
		if err != nil {
			return err
		}

		st.CostMap[k0] = v0
	}

	err, have = _is.SkipTo(codec.MAP, 15, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&length, 0, true)
	if err != nil {
		return err
	}

	st.RetCodeMap = make(map[string]uint32)
	for i1, e1 := int32(0), length; i1 < e1; i1++ {
		var k1 string
		var v1 uint32

		err = _is.Read_string(&k1, 0, false)
		if err != nil {
			return err
		}

		err = _is.Read_uint32(&v1, 1, false)
		if err != nil {
			return err
		}

		st.RetCodeMap[k1] = v1
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *PerfTestDetail) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require PerfTestDetail, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *PerfTestDetail) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_uint32(st.Timestamp, 0)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.Qps, 1)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.TotalReq, 2)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.Succ, 3)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.Failed, 4)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SuccRate, 5)
	if err != nil {
		return err
	}

	err = _os.Write_float32(st.CostMax, 6)
	if err != nil {
		return err
	}

	err = _os.Write_float32(st.CostMin, 7)
	if err != nil {
		return err
	}

	err = _os.Write_float32(st.CostAvg, 8)
	if err != nil {
		return err
	}

	err = _os.Write_float32(st.P90, 9)
	if err != nil {
		return err
	}

	err = _os.Write_float32(st.P99, 10)
	if err != nil {
		return err
	}

	err = _os.Write_float32(st.P999, 11)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.SendByte, 12)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.RecvByte, 13)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.MAP, 14)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.CostMap)), 0)
	if err != nil {
		return err
	}

	for k2, v2 := range st.CostMap {

		err = _os.Write_string(k2, 0)
		if err != nil {
			return err
		}

		err = _os.Write_uint32(v2, 1)
		if err != nil {
			return err
		}

	}

	err = _os.WriteHead(codec.MAP, 15)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.RetCodeMap)), 0)
	if err != nil {
		return err
	}

	for k3, v3 := range st.RetCodeMap {

		err = _os.Write_string(k3, 0)
		if err != nil {
			return err
		}

		err = _os.Write_uint32(v3, 1)
		if err != nil {
			return err
		}

	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *PerfTestDetail) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// CoreUsage struct implement
type CoreUsage struct {
	Percent float32 `json:"percent"`
}

func (st *CoreUsage) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *CoreUsage) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_float32(&st.Percent, 0, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *CoreUsage) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require CoreUsage, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *CoreUsage) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_float32(st.Percent, 0)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *CoreUsage) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// MemUsage struct implement
type MemUsage struct {
	Total int64 `json:"total"`
	Used  int64 `json:"used"`
}

func (st *MemUsage) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *MemUsage) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int64(&st.Total, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.Used, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *MemUsage) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require MemUsage, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *MemUsage) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int64(st.Total, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.Used, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *MemUsage) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// PerfResDetail struct implement
type PerfResDetail struct {
	Timestamp uint32      `json:"timestamp"`
	Cpu       []CoreUsage `json:"cpu"`
	Mem       MemUsage    `json:"mem"`
}

func (st *PerfResDetail) ResetDefault() {
	st.Mem.ResetDefault()
}

//ReadFrom reads  from _is and put into struct.
func (st *PerfResDetail) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_uint32(&st.Timestamp, 0, true)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(1, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Cpu = make([]CoreUsage, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = st.Cpu[i0].ReadBlock(_is, 0, false)
			if err != nil {
				return err
			}

		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}

	}

	err = st.Mem.ReadBlock(_is, 2, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *PerfResDetail) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require PerfResDetail, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *PerfResDetail) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_uint32(st.Timestamp, 0)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.Cpu)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Cpu {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}

	}

	err = st.Mem.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *PerfResDetail) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// TestDetailResp struct implement
type TestDetailResp struct {
	Code       uint32           `json:"code"`
	Msg        string           `json:"msg"`
	PerfDetail []PerfTestDetail `json:"perfDetail"`
	ResUsage   []PerfResDetail  `json:"resUsage"`
}

func (st *TestDetailResp) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *TestDetailResp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_uint32(&st.Code, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Msg, 1, true)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}

		st.PerfDetail = make([]PerfTestDetail, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = st.PerfDetail[i0].ReadBlock(_is, 0, false)
			if err != nil {
				return err
			}

		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}

	}

	err, have, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}

		st.ResUsage = make([]PerfResDetail, length)
		for i1, e1 := int32(0), length; i1 < e1; i1++ {

			err = st.ResUsage[i1].ReadBlock(_is, 0, false)
			if err != nil {
				return err
			}

		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}

	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *TestDetailResp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require TestDetailResp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *TestDetailResp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_uint32(st.Code, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Msg, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.PerfDetail)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.PerfDetail {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}

	}

	err = _os.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.ResUsage)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.ResUsage {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}

	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *TestDetailResp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// QueryTestHistoryReq struct implement
type QueryTestHistoryReq struct {
	PageSize uint32 `json:"pageSize"`
	Page     uint32 `json:"page"`
}

func (st *QueryTestHistoryReq) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *QueryTestHistoryReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_uint32(&st.PageSize, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.Page, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *QueryTestHistoryReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require QueryTestHistoryReq, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *QueryTestHistoryReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_uint32(st.PageSize, 0)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.Page, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *QueryTestHistoryReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// TestHistory struct implement
type TestHistory struct {
	TestID    uint32 `json:"testID"`
	StartTime uint32 `json:"startTime"`
	EndTime   uint32 `json:"endTime"`
	ServType  string `json:"servType"`
	Lang      string `json:"lang"`
	Cores     uint32 `json:"cores"`
	Threads   uint32 `json:"threads"`
	ConnCnt   uint32 `json:"connCnt"`
	KeepAlive uint32 `json:"keepAlive"`
	PkgLen    uint32 `json:"pkgLen"`
}

func (st *TestHistory) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *TestHistory) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_uint32(&st.TestID, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.StartTime, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.EndTime, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.ServType, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Lang, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.Cores, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.Threads, 6, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.ConnCnt, 7, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.KeepAlive, 8, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.PkgLen, 9, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *TestHistory) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require TestHistory, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *TestHistory) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_uint32(st.TestID, 0)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.StartTime, 1)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.EndTime, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.ServType, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Lang, 4)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.Cores, 5)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.Threads, 6)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.ConnCnt, 7)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.KeepAlive, 8)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.PkgLen, 9)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *TestHistory) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// QueryTestHistoryResp struct implement
type QueryTestHistoryResp struct {
	Code      uint32        `json:"code"`
	Msg       string        `json:"msg"`
	Total     uint32        `json:"total"`
	Page      uint32        `json:"page"`
	Histories []TestHistory `json:"histories"`
}

func (st *QueryTestHistoryResp) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *QueryTestHistoryResp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_uint32(&st.Code, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Msg, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.Total, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.Page, 3, true)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(4, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Histories = make([]TestHistory, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = st.Histories[i0].ReadBlock(_is, 0, false)
			if err != nil {
				return err
			}

		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}

	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *QueryTestHistoryResp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require QueryTestHistoryResp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *QueryTestHistoryResp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_uint32(st.Code, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Msg, 1)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.Total, 2)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.Page, 3)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 4)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(st.Histories)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Histories {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}

	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *QueryTestHistoryResp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
