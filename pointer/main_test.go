package main_test

import "testing"

type MyStruct struct {
	F1, F2, F3, F4, F5, F6, F7 string
	I1, I2, I3, I4, I5, I6, I7 int64
}

func accept(s []MyStruct) {
}

func acceptPointer(s []*MyStruct) {
}
func BenchmarkAppendingStructs(b *testing.B) {
	var s []MyStruct

	for i := 0; i < b.N; i++ {
		s = append(s, MyStruct{})
	}
	accept(s)
	accept(s)
	accept(s)
}

func BenchmarkAppendingPointers(b *testing.B) {
	var s []*MyStruct

	for i := 0; i < b.N; i++ {
		s = append(s, &MyStruct{})
	}
	acceptPointer(s)
	acceptPointer(s)
	acceptPointer(s)
}

func BenchmarkAppendingStructsPreallocatedCap(b *testing.B) {
	var s []MyStruct = make([]MyStruct, 0, b.N)

	for i := 0; i < b.N; i++ {
		s = append(s, MyStruct{})
	}
	accept(s)
	accept(s)
	accept(s)
}

func BenchmarkAppendingPointersPreallocatedCap(b *testing.B) {
	var s []*MyStruct = make([]*MyStruct, 0, b.N)

	for i := 0; i < b.N; i++ {
		s = append(s, &MyStruct{})
	}
	acceptPointer(s)
	acceptPointer(s)
	acceptPointer(s)
}

func BenchmarkAppendingStructsPreallocatedLen(b *testing.B) {
	var s []MyStruct = make([]MyStruct, b.N)

	for i := 0; i < b.N; i++ {
		s[i] = MyStruct{}
	}
	accept(s)
	accept(s)
	accept(s)
}

func BenchmarkAppendingPointersPreallocatedLen(b *testing.B) {
	var s []*MyStruct = make([]*MyStruct, b.N)

	for i := 0; i < b.N; i++ {
		s[i] = &MyStruct{}
	}
	acceptPointer(s)
	acceptPointer(s)
	acceptPointer(s)
}
