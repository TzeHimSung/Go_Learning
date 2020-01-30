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

// // Int*, Uint*, Float*, Complex*: Bits
// // Array: Elem, Len
// // Chan: ChanDir, Elem
// // Func: In, NumIn, Out, NumOut, IsVariadic
// // Map: Key, Elem
// // Ptr: Elem
// // Slice: Elem
// // Struct: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField

// // 返回类型的元素类型，该方法只适合Array、Chan、Map、Ptr、Slice类型
// Elem() Type

// // 返回数值型类型内存占用的位数
// Bits() int

// // struct类型专用的方法
// // 返回字段数目
// NumField() int
// // 通过整数索引获取struct字段
// Field(i int) StructField
// // 通过嵌入字段获取struct字段
// FieldByIndex(index []int) StructField
// // 通过名字查找获取struct字段
// FieldByName(name string) (StructField, bool)

// // func类型专用的方法
// // 函数是否为不定参数函数
// IsVariadic() bool
// // 输入参数个数
// NumIn() int
// // 返回值个数
// NumOut() int
// // 返回第i个输入参数类型
// In(i int) Type
// // 返回第i个返回值类型
// Out(i int) Type

// // map类型专用的方法
// // 返回map key的type
// Key() Type

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // Student .
// type Student struct {
// 	Name string "学生姓名"
// 	Age  int    `a:"1111"b:"3333"`
// }

// func main() {
// 	s := Student{}
// 	rt := reflect.TypeOf(s)
// 	fieldName, ok := rt.FieldByName("Name")
// 	// 取tag数据
// 	if ok {
// 		fmt.Println(fieldName.Tag)
// 	}
// 	fieldAge, ok2 := rt.FieldByName("Age")
// 	// 可以像json一样，取tag里的数据，多个tag之间无逗号，tag不需要引号
// 	if ok2 {
// 		fmt.Println(fieldAge.Tag.Get("a"))
// 		fmt.Println(fieldAge.Tag.Get("b"))
// 	}
// 	fmt.Println("type_Name:", rt.Name())
// 	fmt.Println("type_NumField:", rt.NumField())
// 	fmt.Println("type_PkgPath:", rt.PkgPath())
// 	fmt.Println("type_String:", rt.String())
// 	fmt.Println("type.Kind.String:", rt.Kind().String())
// 	fmt.Println("type.String() =", rt.String())
// 	// 获取结构类型的字段名称
// 	for i := 0; i < rt.NumField(); i++ {
// 		fmt.Printf("type.Field[%d].Name := %v \n", i, rt.Field(i).Name)
// 	}
// 	sc := make([]int, 10)
// 	sc = append(sc, 1, 2, 3)
// 	sct := reflect.TypeOf(sc)
// 	// 获取slice元素的type
// 	scet := sct.Elem()
// 	fmt.Println("slice element type.Kind() =", scet.Kind())
// 	fmt.Printf("slice element type.Kind() =%d\n", scet.Kind())
// 	fmt.Println("slice element type.String() =", scet.String())
// 	fmt.Println("slice element type.Name() =", scet.Name())
// 	fmt.Println("slice element type.NumMethod() =", scet.NumMethod())
// 	fmt.Println("slice type.PkgPath() =", scet.PkgPath())
// 	fmt.Println("slice type.PkgPath() =", sct.PkgPath())
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // INT .
// type INT int

// // A ...
// type A struct {
// 	a int
// }

// // B .
// type B struct {
// 	b string
// }

// // Ita .
// type Ita interface {
// 	String() string
// }

// // String .
// func (b B) String() string {
// 	return b.b
// }

// func main() {
// 	var a INT = 12
// 	var b int = 14
// 	// 实参是具体类型，reflect.TypeOf返回其静态类型
// 	ta := reflect.TypeOf(a)
// 	tb := reflect.TypeOf(b)
// 	// INT和int是两个类型，不相等
// 	if ta == tb {
// 		fmt.Println("ta == tb")
// 	} else {
// 		fmt.Println("ta != tb")
// 	}
// 	fmt.Println(ta.Name())
// 	fmt.Println(tb.Name())
// 	// 底层基础类型
// 	fmt.Println(ta.Kind().String())
// 	fmt.Println(tb.Kind().String())
// 	s1 := A{1}
// 	s2 := B{"tata"}
// 	// 实参是具体类型，reflect.TypeOf返回其静态类型
// 	fmt.Println(reflect.TypeOf(s1).Name())
// 	fmt.Println(reflect.TypeOf(s2).Name())
// 	// Type的Kind()方法返回的是基础类型，类型A和B的底层基础都是struct
// 	fmt.Println(reflect.TypeOf(s1).Kind().String()) // struct
// 	fmt.Println(reflect.TypeOf(s2).Kind().String()) // struct
// 	ita := new(Ita)
// 	var itb Ita = s2
// 	// 实参是未绑定具体变量的接口类型，reflect.TypeOf返回接口类型本身
// 	// 也就是接口的静态类型
// 	fmt.Println(reflect.TypeOf(ita).Elem().Name())          // Ita
// 	fmt.Println(reflect.TypeOf(ita).Elem().Kind().String()) // interface
// 	// 实参是绑定了具体变量的接口类型，reflect.TypeOf返回绑定的具体类型
// 	// 也就是接口的动态类型
// 	fmt.Println(reflect.TypeOf(itb).Name())          // B
// 	fmt.Println(reflect.TypeOf(itb).Kind().String()) // struct
// }

// type Value struct{
// 	// typ holds the type of the value represented by a value
// 	typ *rtype
// 	// Pointer-valued data or, if flagIndir is set, pointer to data.
// 	// Valid when either flagIndir is set or typ.pointers() is true
// 	ptr unsafe.Pointer
// 	flag
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // User .
// type User struct {
// 	ID, Age int
// 	Name    string
// }

// func (user User) String() {
// 	fmt.Println("User:", user.ID, user.Name, user.Age)
// }

// // Info .
// func Info(o interface{}) {
// 	// 获取value信息
// 	v := reflect.ValueOf(o)
// 	// 通过value获取type
// 	t := v.Type()
// 	// 类型名称
// 	fmt.Println("Type:", t.Name())
// 	// 访问接口字段名、字段类型和字段值信息
// 	fmt.Println("Fields:")
// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		value := v.Field(i).Interface()
// 		// 类型查询
// 		switch value := value.(type) {
// 		case int:
// 			fmt.Printf(" %6s: %v = %d\n", field.Name, field.Type, value)
// 		case string:
// 			fmt.Printf(" %6s: %v = %s\n", field.Name, field.Type, value)
// 		default:
// 			fmt.Printf(" %6s: %v = %d\n", field.Name, field.Type, value)
// 		}
// 	}
// }

// func main() {
// 	u := User{1, 30, "Tom"}
// 	Info(u)
// }

// type A struct{
// 	a int
// }

// type Aa A

// type B struct{
// 	b int
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // User .
// type User struct {
// 	ID, Age int
// 	Name    string
// }

// func main() {
// 	u := User{ID: 1, Name: "andes", Age: 20}
// 	va := reflect.ValueOf(u)
// 	vb := reflect.ValueOf(&u)
// 	// 值类型是可修改的
// 	fmt.Println(va.CanSet(), va.FieldByName("Name").CanSet())        // false false
// 	fmt.Println(vb.CanSet(), vb.Elem().FieldByName("Name").CanSet()) // false true
// 	fmt.Printf("%v\n", vb)
// 	name := "shine"
// 	vc := reflect.ValueOf(name)
// 	// 通过set函数修改变量的值
// 	vb.Elem().FieldByName("Name").Set(vc)
// 	fmt.Printf("%v\n", vb)
// }

