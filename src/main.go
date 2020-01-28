// // 返回包含包名的类型名字，对于未命名类型，返回nil
// Name() string

// // Kind返回该类型的底层基础类型
// Kind() Kind

// // 确定当前类型是否实现了u接口类型
// // 注意这里的u必须是接口类型的Type
// Implements(u Type) bool

// // 判断当前类型的实例是否能赋值给type为u的类型变量
// AssignableTo(u Type) bool

// // 判断当前类型的实例是否能强制类型转换为u类型变量
// ConvertibleTo(u Type) bool

// // 判断当前类型是否支持比较(等于或不等于)
// // 支持等于的类型可以作为map的key
// Comparable() bool

// // 返回一个类型的方法的个数
// NumMethod() int

// // 通过索引值访问方法，索引值必须属于[0, NumMethod())，否则引发panic
// Method(int) Method

// // 通过方法名获取Method
// MethodByName(string)(Method, bool)

// // 返回类型的包路径，如果类型是预声明类型或未命名类型，则返回空字符串
// PkgPath() string

// // 返回存放该类型的实例需要多大的字节空间
// Size() uintptr

// -----------------------------------------------------------------------------------

// Int*, Uint*, Float*, Complex*: Bits
// Array: Elem, Len
// Chan: ChanDir, Elem
// Func: In, NumIn, Out, NumOut, IsVariadic
// Map: Key, Elem
// Ptr: Elem
// Slice: Elem
// Struct: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField