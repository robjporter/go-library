package main

import (
	"fmt"
	"time"

	"github.com/robjporter/go-library/xas"
)

func main() {
	time.Sleep(1 * time.Second)
	astostring()
	time.Sleep(1 * time.Second)
	astrimmed()
	time.Sleep(1 * time.Second)
	astofloat()
	time.Sleep(1 * time.Second)
	astorunelength()
	time.Sleep(1 * time.Second)
	astobool()
	time.Sleep(1 * time.Second)
	astobytes()
	time.Sleep(1 * time.Second)
	astoslice()
	time.Sleep(1 * time.Second)
	astoint()
	time.Sleep(1 * time.Second)
	astoip()
	time.Sleep(1 * time.Second)
	astobase64()
	time.Sleep(1 * time.Second)
	asfrombase64()
	time.Sleep(1 * time.Second)
	asisempty()
	time.Sleep(1 * time.Second)
	asiskind()
	time.Sleep(1 * time.Second)
	asofkind()
	time.Sleep(1 * time.Second)
	asoftype()
	time.Sleep(1 * time.Second)
	astotime()
	time.Sleep(1 * time.Second)
	astoduration()
	time.Sleep(1 * time.Second)
	astofixedlengthafter()
	time.Sleep(1 * time.Second)
	astofixedlengthbefore()
	time.Sleep(1 * time.Second)
	astofixedlengthcenter()
	time.Sleep(1 * time.Second)
	asisint()
	time.Sleep(1 * time.Second)
	asisbool()
	time.Sleep(1 * time.Second)
	asisfloat()
	time.Sleep(1 * time.Second)
	asisstring()
	time.Sleep(1 * time.Second)
	asistime()
	time.Sleep(1 * time.Second)
	asisnillable()
	time.Sleep(1 * time.Second)
	astoformattedbytes()
}

func astostring() {
	fmt.Println("")
	fmt.Println("AS TO STRING *******************************************************")
	fmt.Println("STRING: (32)                        >", `"`+xas.ToString(32)+`"`)
	fmt.Println("STRING: (true)                      >", `"`+xas.ToString(bool(true))+`"`)
	fmt.Println("STRING: ('mayonegg')                >", `"`+xas.ToString("mayonegg")+`"`)         // "mayonegg"
	fmt.Println("STRING: (8)                         >", `"`+xas.ToString(8)+`"`)                  // "8"
	fmt.Println("STRING: (8.31)                      >", `"`+xas.ToString(8.31)+`"`)               // "8.31"
	fmt.Println("STRING: ([]byte('one time'))        >", `"`+xas.ToString([]byte("one time"))+`"`) // "one time"
	fmt.Println("STRING: (nil)                       >", `"`+xas.ToString(nil)+`"`)                // ""
	var foo interface{} = "one more time"
	fmt.Println("STRING: (interface{'one more time}) >", `"`+xas.ToString(foo)) // "one more time"
}
func astrimmed() {
	fmt.Println("")
	fmt.Println("AS TRIMMED *******************************************************")
	fmt.Println("TRIMMED: ('    TEST      ')         >", `"`+xas.Trimmed("    TEST      ")+`"`)
}
func astofloat() {
	fmt.Println("")
	fmt.Println("AS TO FLOAT *******************************************************")
	fmt.Println("FLOAT: (32.4400)                    >", xas.ToFloat(32.4400))
	fmt.Println("FLOAT32: (32.4400)                  >", xas.ToFloat32(32.4400))
}
func astorunelength() {
	fmt.Println("")
	fmt.Println("AS TO RUNE LENGTH*******************************************************")
	fmt.Println("RUNELENGTH: ('test')                >", xas.ToRuneLength("test"))
	fmt.Println("RUNELENGTH: ('TEST')                >", xas.ToRuneLength("TEST"))
	fmt.Println("RUNELENGTH: ('iiii')                >", xas.ToRuneLength("iiii"))
	fmt.Println("RUNELENGTH: ('QQKK')                >", xas.ToRuneLength("QQKK"))
	fmt.Println("RUNELENGTH: ('Lllm')                >", xas.ToRuneLength("Lllm"))
}
func astobool() {
	fmt.Println("")
	fmt.Println("AS TO BOOL *******************************************************")
	fmt.Println("BOOL: (1)                           >", xas.ToBool(1))
	fmt.Println("BOOL: (0)                           >", xas.ToBool(0))
	fmt.Println("BOOL: ('1')                         >", xas.ToBool("1"))
	fmt.Println("BOOL: ('true')                      >", xas.ToBool("true"))
	fmt.Println("BOOL: ('down')                      >", xas.ToBool("down"))
}
func astobytes() {
	fmt.Println("")
	fmt.Println("AS TO BYTES *******************************************************")
	fmt.Println("BYTES: ('Testing')                  >", xas.ToBytes("Testing"))
}
func astoslice() {
	fmt.Println("")
	fmt.Println("AS TO SLICE *******************************************************")
	var foo2 []interface{}
	foo2 = append(foo2, "one") //more time"
	fmt.Println("SLICE: ('one')                      >", xas.ToSlice(foo2))
}
func astoint() {
	fmt.Println("")
	fmt.Println("TO INT *******************************************************")
	fmt.Println("INT: ('1')                          >", xas.ToInt("1"))
	fmt.Println("INT64: ('1')                        >", xas.ToInt64("1"))
	fmt.Println("INT32: ('1')                        >", xas.ToInt32("1"))
	fmt.Println("INT16: ('1')                        >", xas.ToInt16("1"))
	fmt.Println("INT8: ('1')                         >", xas.ToInt8("1"))
}

func astoip() {
	fmt.Println("")
	fmt.Println("TO IP *******************************************************")
	fmt.Println("IP ADDRESS: ('192.168.0.1')          >", xas.ToIP("192.168.0.1"))   // "one more time"
	fmt.Println("IP ADDRESS: ('one more time')        >", xas.ToIP("one more time")) //
	fmt.Println("IP ADDRESS: ('1')                    >", xas.ToIP("1"))             // "one more time"
	fmt.Println("IP ADDRESS: ('1.0')                  >", xas.ToIP("1.0"))           // "one more time"
	fmt.Println("IP ADDRESS: ('1.0.0')                >", xas.ToIP("1.0.0"))         // "one more time"
	fmt.Println("IP ADDRESS: ('1.0.0.0/8')            >", xas.ToIP("1.0.0.0/8"))     // "one more time"
}

func astobase64() {
	fmt.Println("")
	fmt.Println("TO BASE64 *******************************************************")
	fmt.Println("TOBASE64: ('This is a test')         >", xas.ToBase64("This is a test"))
}

func asfrombase64() {
	fmt.Println("")
	fmt.Println("FROM BASE64 *******************************************************")
	fmt.Println("FROMBASE64: ('VGhpcyBpcyBhIHRlc3Q=') >", xas.FromBase64("VGhpcyBpcyBhIHRlc3Q="))
}

func asisempty() {
	fmt.Println("")
	fmt.Println("AS IS EMPTY *******************************************************")
	fmt.Println("IP EMPTY: ('0')                      >", xas.IsEmpty(0))
	fmt.Println("IP EMPTY: ('1')                      >", xas.IsEmpty(1))
	fmt.Println("IP EMPTY: ('')                       >", xas.IsEmpty(""))
	fmt.Println("IP EMPTY: ('sdasdass')               >", xas.IsEmpty("sdasdass"))
	fmt.Println("IP EMPTY: ('[]string{}')             >", xas.IsEmpty([]string{}))
}

func asiskind() {
	fmt.Println("")
	fmt.Println("AS IS KIND *******************************************************")
	fmt.Println("IS KIND: (string,0)                  >", xas.IsKind("string", 0))
	fmt.Println("IS KIND: (string,'')                 >", xas.IsKind("string", ""))
	fmt.Println("IS KIND: (int,0)                     >", xas.IsKind("int", 0))
	fmt.Println("IS KIND: (int,'test')                >", xas.IsKind("int", "test"))
}

func asofkind() {
	fmt.Println("")
	fmt.Println("AS OF KIND *******************************************************")
	fmt.Println("KIND OF: ('string')                  >", xas.OfKind("string"))
	fmt.Println("KIND OF: ([]string{})                >", xas.OfKind([]string{}))
	fmt.Println("KIND OF: (nil)                       >", xas.OfKind(nil))
	fmt.Println("KIND OF: ([]byte('one time))         >", xas.OfKind([]byte("one time")))
	fmt.Println("KIND OF: (bool(true))                >", xas.OfKind(bool(true)))
	fmt.Println("KIND OF: (32)                        >", xas.OfKind(32))
}
func asoftype() {
	fmt.Println("")
	fmt.Println("AS OF TYPE *******************************************************")
	fmt.Println("TYPE: (32)                           >", xas.OfType(32))
	fmt.Println("TYPE: ('')                           >", xas.OfType(""))
	fmt.Println("TYPE: ([]string{}])                  >", xas.OfType([]string{}))
	fmt.Println("TYPE: (true)                         >", xas.OfType(true))
	fmt.Println("TYPE: (1.0f)                         >", xas.OfType(1.00))
	fmt.Println("TYPE: (int64(22))                    >", xas.OfType(int64(22)))
}

func astotime() {
	fmt.Println("")
	fmt.Println("AS TO TIME *******************************************************")
	fmt.Println("TIME: ('2016-04-04')                 >", xas.ToTime(false, "2016-04-04"))
	fmt.Println("TIME: ('04-04-2016')                 >", xas.ToTime(false, "04-04-2016"))
	fmt.Println("TIME: ('2016-04-04 16:20:40')        >", xas.ToTime(false, "2016-04-04 16:20:40"))
	fmt.Println("TIME: ('2016-04-04 16:20:40 +1 BST') >", xas.ToTime(false, "2016-04-04 16:20:40 +1 BST"))
	t1 := time.Now()
	fmt.Println("TIME: NOW TO INT                     >", xas.FromTime(t1))
	fmt.Println("TIME: INT TO TIME                    >", xas.ToTime(true, xas.FromTime(t1)))
}
func astoduration() {
	fmt.Println("")
	fmt.Println("AS TO DURATION *******************************************************")
	fmt.Println("DURATION: (1h44m)                    >", xas.ToDuration("1h44m"))
	fmt.Println("DURATION: (44)                       >", xas.ToDuration("44"))
	fmt.Println("DURATION: (44s)                      >", xas.ToDuration("44s"))
	fmt.Println("DURATION: (444h)                     >", xas.ToDuration("444h"))
	fmt.Println("DURATION: (88m)                      >", xas.ToDuration("88m"))
}

func astofixedlengthafter() {
	fmt.Println("")
	fmt.Println("AS TO FIXED LENGTH AFTER *******************************************************")
	fmt.Println("FIXED LENGTH AFTER (*,20):           >", xas.ToFixedLengthAfter("Test String", "*", 20))
	fmt.Println("FIXED LENGTH AFTER (-,50):           >", xas.ToFixedLengthAfter("Test String", "-", 50))
	fmt.Println("FIXED LENGTH AFTER (*,10):           >", xas.ToFixedLengthAfter("Test String", "*", 10))
	fmt.Println("FIXED LENGTH AFTER (*,8):            >", xas.ToFixedLengthAfter("Test String", "*", 8))
}

func astofixedlengthbefore() {
	fmt.Println("")
	fmt.Println("AS TO FIXED LENGTH BEFORE *******************************************************")
	fmt.Println("FIXED LENGTH BEFORE (*,20):          >", xas.ToFixedLengthBefore("Test String", "*", 20))
	fmt.Println("FIXED LENGTH BEFORE (-,50):          >", xas.ToFixedLengthBefore("Test String", "-", 50))
	fmt.Println("FIXED LENGTH BEFORE (*,10):          >", xas.ToFixedLengthBefore("Test String", "*", 10))
	fmt.Println("FIXED LENGTH BEFORE (*,8):           >", xas.ToFixedLengthBefore("Test String", "*", 8))
}

func astofixedlengthcenter() {
	fmt.Println("")
	fmt.Println("AS TO FIXED LENGTH CENTER *******************************************************")
	fmt.Println("FIXED LENGTH CENTER (*,20):          >", xas.ToFixedLengthCenter("Test String", "*", 20))
	fmt.Println("FIXED LENGTH CENTER (-,50):          >", xas.ToFixedLengthCenter("Test String", "-", 50))
	fmt.Println("FIXED LENGTH CENTER (*,10):          >", xas.ToFixedLengthCenter("Test String", "*", 10))
	fmt.Println("FIXED LENGTH CENTER (*,8):           >", xas.ToFixedLengthCenter("Test String", "*", 8))

}

func asisint() {
	fmt.Println("")
	fmt.Println("AS IS INT *******************************************************")
	fmt.Println("INT: (44)                           >", xas.IsInt(44))
	fmt.Println("INT: (true)                         >", xas.IsInt(true))
	fmt.Println("INT: (44.44)                        >", xas.IsInt(44.44))
	fmt.Println("INT: ('test')                       >", xas.IsInt("test"))
	fmt.Println("INT: ('14:14:14')                   >", xas.IsInt(xas.ToTime(false, "14:14:14")))
}

func asisbool() {
	fmt.Println("")
	fmt.Println("AS IS BOOL *******************************************************")
	fmt.Println("BOOL: (44)                           >", xas.IsBool(44))
	fmt.Println("BOOL: (true)                         >", xas.IsBool(true))
	fmt.Println("BOOL: (44.44)                        >", xas.IsBool(44.44))
	fmt.Println("BOOL: ('test')                       >", xas.IsBool("test"))
	fmt.Println("BOOL: ('14:14:14')                   >", xas.IsBool(xas.ToTime(false, "14:14:14")))
}

func asisfloat() {
	fmt.Println("")
	fmt.Println("AS IS FLOAT *******************************************************")
	fmt.Println("FLOAT: (44)                           >", xas.IsFloat(44))
	fmt.Println("FLOAT: (true)                         >", xas.IsFloat(true))
	fmt.Println("FLOAT: (44.44)                        >", xas.IsFloat(44.44))
	fmt.Println("FLOAT: ('test')                       >", xas.IsFloat("test"))
	fmt.Println("FLOAT: ('14:14:14')                   >", xas.IsFloat(xas.ToTime(false, "14:14:14")))
}

func asisstring() {
	fmt.Println("")
	fmt.Println("AS IS STRING *******************************************************")
	fmt.Println("STRING: (44)                         >", xas.IsString(44))
	fmt.Println("STRING: (true)                       >", xas.IsString(true))
	fmt.Println("STRING: (44.44)                      >", xas.IsString(44.44))
	fmt.Println("STRING: ('test')                     >", xas.IsString("test"))
	fmt.Println("STRING: ('14:14:14')                 >", xas.IsString(xas.ToTime(false, "14:14:14")))
}

func asistime() {
	fmt.Println("")
	fmt.Println("AS IS TIME *******************************************************")
	fmt.Println("TIME: (44)                           >", xas.IsTime(44))
	fmt.Println("TIME: (true)                         >", xas.IsTime(true))
	fmt.Println("TIME: (44.44)                        >", xas.IsTime(44.44))
	fmt.Println("TIME: ('test')                       >", xas.IsTime("test"))
	fmt.Println("TIME: ('14:14:14')                   >", xas.IsTime(xas.ToTime(false, "14:14:14")))
}

func asisnillable() {
	fmt.Println("")
	fmt.Println("AS IS NILLABLE *******************************************************")
	fmt.Println("NILLABLE: ('')                       >", xas.IsNillable(""))
	fmt.Println("NILLABLE: ([]string{})               >", xas.IsNillable([]string{}))
}

func astoformattedbytes() {
	fmt.Println("")
	fmt.Println("AS TO FORMATTED BYTES *******************************************************")
	fmt.Println("FORMAT: (44)                       >", xas.FormatIntToByte(44))
	fmt.Println("FORMAT: (444)                      >", xas.FormatIntToByte(444))
	fmt.Println("FORMAT: (4444)                     >", xas.FormatIntToByte(4444))
	fmt.Println("FORMAT: (44444)                    >", xas.FormatIntToByte(44444))
	fmt.Println("FORMAT: (444444)                   >", xas.FormatIntToByte(444444))
	fmt.Println("FORMAT: (4444444)                  >", xas.FormatIntToByte(4444444))
	fmt.Println("FORMAT: (44444444)                 >", xas.FormatIntToByte(44444444))
	fmt.Println("FORMAT: (444444444)                >", xas.FormatIntToByte(444444444))
	fmt.Println("FORMAT: (4444444444)               >", xas.FormatIntToByte(4444444444))
	fmt.Println("FORMAT: (44444444444)              >", xas.FormatIntToByte(44444444444))
	fmt.Println("FORMAT: (444444444444)             >", xas.FormatIntToByte(444444444444))
	fmt.Println("FORMAT: (4444444444444)            >", xas.FormatIntToByte(4444444444444))
	fmt.Println("FORMAT: (44444444444444)           >", xas.FormatIntToByte(44444444444444))
	fmt.Println("FORMAT: (444444444444444)          >", xas.FormatIntToByte(444444444444444))
	fmt.Println("FORMAT: (4444444444444444)         >", xas.FormatIntToByte(4444444444444444))
	fmt.Println("FORMAT: (44444444444444444)        >", xas.FormatIntToByte(44444444444444444))
	fmt.Println("FORMAT: (444444444444444444)       >", xas.FormatIntToByte(444444444444444444))
	fmt.Println("FORMAT: (999999999999999999)       >", xas.FormatIntToByte(999999999999999999))
	fmt.Println("FORMAT: (1000000000000000000)      >", xas.FormatIntToByte(1152921504606846976))
}
