package configloader

//Type field type
type Type string

//TypeUnkonwn field type unkowwn
var TypeUnkonwn = Type("")

//TypeBool field type bool
var TypeBool = Type("configloader.bool")

//TypeString field type string
var TypeString = Type("configloader.string")

//TypeInt8 field type int8
var TypeInt8 = Type("configloader.int8")

//TypeUint8 field type uint8
var TypeUint8 = Type("configloader.uint8")

//TypeInt16 field type int16
var TypeInt16 = Type("configloader.int16")

//TypeUint16 field type uint16
var TypeUint16 = Type("configloader.uint16")

//TypeInt filed type int
var TypeInt = Type("configloader.int")

//TypeUint field type uint
var TypeUint = Type("configloader.uint")

//TypeInt64 field type int64
var TypeInt64 = Type("configloader.int64")

//TypeUint64 field type uint64
var TypeUint64 = Type("configloader.uint64")

//TypeFloat32 field type float32
var TypeFloat32 = Type("configloader.float32")

//TypeFloat64 field type float64
var TypeFloat64 = Type("configloader.float64")

//TypeMap field type map
var TypeMap = Type("configloader.map")

//TypeArray field type array
var TypeArray = Type("configloader.array")

//TypeSlice field type slice
var TypeSlice = Type("configloader.slice")

//TypeStruct field type struct
var TypeStruct = Type("configloader.struct")

//TypeStructField field type struct field
var TypeStructField = Type("configloader.structFild")

//TypeEmptyInterface field type empty interface
var TypeEmptyInterface = Type("configloader.interface{}")

//TypeLazyLoadFunc field type lazyload func
var TypeLazyLoadFunc = Type("configloader.lazyloadfunc")

//TypeLazyLoader field type lazyloader
var TypeLazyLoader = Type("configloader.lazyloader")

//TypePtr field type pointer
var TypePtr = Type("configloader.*")
