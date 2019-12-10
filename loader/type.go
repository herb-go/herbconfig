package loader

//Type field type
type Type string

//TypeUnkonwn field type unkowwn
var TypeUnkonwn = Type("")

//TypeBool field type bool
var TypeBool = Type("loader.bool")

//TypeString field type string
var TypeString = Type("loader.string")

//TypeInt8 field type int8
var TypeInt8 = Type("loader.int8")

//TypeUint8 field type uint8
var TypeUint8 = Type("loader.uint8")

//TypeInt16 field type int16
var TypeInt16 = Type("loader.int16")

//TypeUint16 field type uint16
var TypeUint16 = Type("loader.uint16")

//TypeInt filed type int
var TypeInt = Type("loader.int")

//TypeUint field type uint
var TypeUint = Type("loader.uint")

//TypeInt64 field type int64
var TypeInt64 = Type("loader.int64")

//TypeUint64 field type uint64
var TypeUint64 = Type("loader.uint64")

//TypeFloat32 field type float32
var TypeFloat32 = Type("loader.float32")

//TypeFloat64 field type float64
var TypeFloat64 = Type("loader.float64")

//TypeMap field type map
var TypeMap = Type("loader.map")

//TypeArray field type array
var TypeArray = Type("loader.array")

//TypeSlice field type slice
var TypeSlice = Type("loader.slice")

//TypeStruct field type struct
var TypeStruct = Type("loader.struct")

//TypeStructField field type struct field
var TypeStructField = Type("loader.structFild")

//TypeEmptyInterface field type empty interface
var TypeEmptyInterface = Type("loader.interface{}")

//TypeLazyLoadFunc field type lazyload func
var TypeLazyLoadFunc = Type("loader.lazyloadfunc")

//TypeLazyLoader field type lazyloader
var TypeLazyLoader = Type("loader.lazyloader")

//TypePtr field type pointer
var TypePtr = Type("loader.*")
