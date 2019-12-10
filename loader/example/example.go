package example

import (
	"reflect"
)

type ExampleStruct struct {
	FieldBasicFieldStruct      BasicFieldStruct
	FieldBasicFieldStructSlice []BasicFieldStruct
	FieldBasicFieldStructMap   map[string]BasicFieldStruct
}
type BasicFieldStruct struct {
	FieldBool         bool
	FieldBoolSlice    []bool
	FieldBoolMap      map[string]bool
	FieldInt          int
	FieldIntSlice     []int
	FieldIntMap       map[string]int
	FieldInt64        int64
	FieldInt64Slice   []int64
	FieldInt64Map     map[string]int64
	FieldFloat32      float32
	FieldFloat32Slice []float32
	FieldFloat32Map   map[string]float32
	FieldFloat64      float64
	FieldFloat64Slice []float64
	FieldFloat64Map   map[string]float64
	FieldString       string
	FieldStringSlice  []string
	FieldStringMap    map[string]string
}

func (s *ExampleStruct) Equal(target *ExampleStruct) bool {
	return reflect.DeepEqual(*s, *target)
}

var BasicFieldData = &BasicFieldStruct{
	FieldBool:         true,
	FieldBoolSlice:    []bool{true, false},
	FieldBoolMap:      map[string]bool{"true": true, "false": false},
	FieldInt:          int(32),
	FieldIntSlice:     []int{int(33), int(34)},
	FieldIntMap:       map[string]int{"36": int(36), "37": int(37)},
	FieldInt64:        int64(64),
	FieldInt64Slice:   []int64{int64(65), int64(66)},
	FieldInt64Map:     map[string]int64{"67": int64(67), "68": int64(68)},
	FieldFloat32:      float32(32.0),
	FieldFloat32Slice: []float32{float32(33.0), float32(34.0)},
	FieldFloat32Map:   map[string]float32{"36": float32(36.0), "37": float32(37.0)},
	FieldFloat64:      float64(64.0),
	FieldFloat64Slice: []float64{float64(65.0), float64(66.0)},
	FieldFloat64Map:   map[string]float64{"67": float64(67.0), "68": float64(68.0)},
	FieldString:       "stringvalue1",
	FieldStringSlice:  []string{"stringvalue2", "stringvalue3"},
	FieldStringMap:    map[string]string{"stringkey4": "stringvalue4", "stringkey5": "stringvalue5"},
}

var ExampleData = &ExampleStruct{
	FieldBasicFieldStruct:      *BasicFieldData,
	FieldBasicFieldStructSlice: []BasicFieldStruct{*BasicFieldData, *BasicFieldData},
	FieldBasicFieldStructMap:   map[string]BasicFieldStruct{"value1": *BasicFieldData, "value2": *BasicFieldData},
}
