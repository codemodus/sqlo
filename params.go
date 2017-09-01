package sqlo

import (
	"database/sql"
	"strings"
	"time"
)

var (
	dirAsc = "ASC"
	dirDsc = "DESC"

	lowTime  = time.Time{}
	highTime = time.Time{}.AddDate(9001, 0, 0)

	lowInt32  int32 = -1 << 31
	highInt32 int32 = 1<<31 - 1

	highUint32 uint32 = 1<<32 - 1

	lowInt64  int64 = -1 << 63
	highInt64 int64 = 1<<63 - 1

	highUint64 uint64 = 1<<64 - 1
)

// LimitLapse ...
type LimitLapse struct {
	Limit uint32
	Lapse uint32
}

// MakeLimitLapse ...
func MakeLimitLapse(lim, lap uint32) LimitLapse {
	if lim == 0 {
		lim = 4e+9
	}

	return LimitLapse{
		Limit: lim,
		Lapse: lap,
	}
}

// SpanTime ...
type SpanTime struct {
	First time.Time
	Final time.Time
	Desc  bool
}

// MakeSpanTime ...
func MakeSpanTime(first, final time.Time, each, desc bool) SpanTime {
	if each {
		return SpanTime{
			First: lowTime,
			Final: highTime,
			Desc:  desc,
		}
	}

	if final.Before(first) {
		final = first
	}

	return SpanTime{
		First: first,
		Final: final,
		Desc:  desc,
	}
}

// SwapAsc ...
func (st *SpanTime) SwapAsc(s string) string {
	return swapAsc(s, st.Desc)
}

// SpanInt32 ...
type SpanInt32 struct {
	First int32
	Final int32
	Desc  bool
}

// MakeSpanInt32 ...
func MakeSpanInt32(first, final int32, each, desc bool) SpanInt32 {
	if each {
		return SpanInt32{
			First: lowInt32,
			Final: highInt32,
			Desc:  desc,
		}
	}

	if final < first {
		final = first
	}

	return SpanInt32{
		First: first,
		Final: final,
		Desc:  desc,
	}
}

// SwapAsc ...
func (si *SpanInt32) SwapAsc(s string) string {
	return swapAsc(s, si.Desc)
}

// SpanUint32 ...
type SpanUint32 struct {
	First uint32
	Final uint32
	Desc  bool
}

// MakeSpanUint32 ...
func MakeSpanUint32(first, final uint32, each, desc bool) SpanUint32 {
	if each {
		return SpanUint32{
			First: 0,
			Final: highUint32,
			Desc:  desc,
		}
	}

	if final < first {
		final = first
	}

	return SpanUint32{
		First: first,
		Final: final,
		Desc:  desc,
	}
}

// SwapAsc ...
func (su *SpanUint32) SwapAsc(s string) string {
	return swapAsc(s, su.Desc)
}

// SpanInt64 ...
type SpanInt64 struct {
	First int64
	Final int64
	Desc  bool
}

// MakeSpanInt64 ...
func MakeSpanInt64(first, final int64, each, desc bool) SpanInt64 {
	if each {
		return SpanInt64{
			First: lowInt64,
			Final: highInt64,
			Desc:  desc,
		}
	}

	if final < first {
		final = first
	}

	return SpanInt64{
		First: first,
		Final: final,
		Desc:  desc,
	}
}

// SwapAsc ...
func (si *SpanInt64) SwapAsc(s string) string {
	return swapAsc(s, si.Desc)
}

// SpanUint64 ...
type SpanUint64 struct {
	First uint64
	Final uint64
	Desc  bool
}

// MakeSpanUint64 ...
func MakeSpanUint64(first, final uint64, each, desc bool) SpanUint64 {
	if each {
		return SpanUint64{
			First: 0,
			Final: highUint64,
			Desc:  desc,
		}
	}

	if final < first {
		final = first
	}

	return SpanUint64{
		First: first,
		Final: final,
		Desc:  desc,
	}
}

// SwapAsc ...
func (su *SpanUint64) SwapAsc(s string) string {
	return swapAsc(s, su.Desc)
}

func swapAsc(s string, desc bool) string {
	if !desc {
		return s
	}

	return strings.Replace(s, dirAsc, dirDsc, -1)
}

// Int32OrEach ...
type Int32OrEach struct {
	Val *int32
}

// MakeInt32OrEach ...
func MakeInt32OrEach(n int32, each bool) Int32OrEach {
	ret := Int32OrEach{Val: nil}

	if each {
		return ret
	}

	ret.Val = &n

	return ret
}

// Uint32OrEach ...
type Uint32OrEach struct {
	Val *uint32
}

// MakeUint32OrEach ...
func MakeUint32OrEach(n uint32, each bool) Uint32OrEach {
	ret := Uint32OrEach{Val: nil}

	if each {
		return ret
	}

	ret.Val = &n

	return ret
}

// Int64OrEach ...
type Int64OrEach struct {
	Val *int64
}

// MakeInt64OrEach ...
func MakeInt64OrEach(n int64, each bool) Int64OrEach {
	ret := Int64OrEach{Val: nil}

	if each {
		return ret
	}

	ret.Val = &n

	return ret
}

// Uint64OrEach ...
type Uint64OrEach struct {
	Val *uint64
}

// MakeUint64OrEach ...
func MakeUint64OrEach(n uint64, each bool) Uint64OrEach {
	ret := Uint64OrEach{Val: nil}

	if each {
		return ret
	}

	ret.Val = &n

	return ret
}

// BoolOrEach ...
type BoolOrEach struct {
	Val sql.NullBool
}

// MakeBoolOrEach ...
func MakeBoolOrEach(b, each bool) BoolOrEach {
	return BoolOrEach{
		Val: sql.NullBool{
			Bool:  b,
			Valid: !each,
		},
	}
}

// StringOrEach ...
type StringOrEach struct {
	Val sql.NullString
}

// MakeStringOrEach ...
func MakeStringOrEach(str string, each bool) StringOrEach {
	if !each && str == "" {
		str = ":|:"
	}

	return StringOrEach{
		Val: sql.NullString{
			String: str,
			Valid:  !each,
		},
	}
}

// MakeBoolOrEach ...
func (s *SQLO) MakeBoolOrEach(b, each bool) BoolOrEach {
	return MakeBoolOrEach(b, each)
}

// MakeInt32OrEach ...
func (s *SQLO) MakeInt32OrEach(n int32, each bool) Int32OrEach {
	return MakeInt32OrEach(n, each)
}

// MakeInt64OrEach ...
func (s *SQLO) MakeInt64OrEach(n int64, each bool) Int64OrEach {
	return MakeInt64OrEach(n, each)
}

// MakeLimitLapse ...
func (s *SQLO) MakeLimitLapse(lim, lap uint32) LimitLapse {
	return MakeLimitLapse(lim, lap)
}

// MakeSpanInt32 ...
func (s *SQLO) MakeSpanInt32(first, final int32, each, desc bool) SpanInt32 {
	return MakeSpanInt32(first, final, each, desc)
}

// MakeSpanInt64 ...
func (s *SQLO) MakeSpanInt64(first, final int64, each, desc bool) SpanInt64 {
	return MakeSpanInt64(first, final, each, desc)
}

// MakeSpanTime ...
func (s *SQLO) MakeSpanTime(first, final time.Time, each, desc bool) SpanTime {
	return MakeSpanTime(first, final, each, desc)
}

// MakeSpanUint32 ...
func (s *SQLO) MakeSpanUint32(first, final uint32, each, desc bool) SpanUint32 {
	return MakeSpanUint32(first, final, each, desc)
}

// MakeSpanUint64 ...
func (s *SQLO) MakeSpanUint64(first, final uint64, each, desc bool) SpanUint64 {
	return MakeSpanUint64(first, final, each, desc)
}

// MakeStringOrEach ...
func (s *SQLO) MakeStringOrEach(str string, each bool) StringOrEach {
	return MakeStringOrEach(str, each)
}

// MakeUint32OrEach ...
func (s *SQLO) MakeUint32OrEach(n uint32, each bool) Uint32OrEach {
	return MakeUint32OrEach(n, each)
}

// MakeUint64OrEach ...
func (s *SQLO) MakeUint64OrEach(n uint64, each bool) Uint64OrEach {
	return MakeUint64OrEach(n, each)
}
