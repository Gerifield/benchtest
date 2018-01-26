package benchtest

import (
	"github.com/dgryski/go-metro"
)

var sMap = make(map[string]struct{})
var iMap = make(map[int]struct{})
var i32Map = make(map[int32]struct{})
var i64Map = make(map[int64]struct{})
var f32Map = make(map[float32]struct{})
var f64Map = make(map[float64]struct{})
var ui32Map = make(map[uint32]struct{})
var ui64Map = make(map[uint64]struct{})

var metroMap = make(map[uint64]struct{})

func sMapInsert(k string) {
	sMap[k] = struct{}{}
}

func sMapGet(k string) struct{} {
	return sMap[k]
}

func iMapInsert(k int) {
	iMap[k] = struct{}{}
}

func iMapGet(k int) struct{} {
	return iMap[k]
}

func i32MapInsert(k int32) {
	i32Map[k] = struct{}{}
}

func i32MapGet(k int32) struct{} {
	return i32Map[k]
}

func i64MapInsert(k int64) {
	i64Map[k] = struct{}{}
}

func i64MapGet(k int64) struct{} {
	return i64Map[k]
}

func f32MapInsert(k float32) {
	f32Map[k] = struct{}{}
}

func f32MapGet(k float32) struct{} {
	return f32Map[k]
}

func f64MapInsert(k float64) {
	f64Map[k] = struct{}{}
}

func f64MapGet(k float64) struct{} {
	return f64Map[k]
}

func ui64MapInsert(k uint64) {
	ui64Map[k] = struct{}{}
}

func ui64MapGet(k uint64) struct{} {
	return ui64Map[k]
}

func ui32MapInsert(k uint32) {
	ui32Map[k] = struct{}{}
}

func ui32MapGet(k uint32) struct{} {
	return ui32Map[k]
}

func metroMapInsert(s string) {
	metroMap[metro.Hash64Str(s, 0)] = struct{}{}
}

func metroMapGet(s string) struct{} {
	return metroMap[metro.Hash64Str(s, 0)]
}
