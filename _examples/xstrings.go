package main

import (
	"fmt"
	"strconv"

	"../xstrings"
)

func main() {
	fmt.Println(xstrings.Center("This is a test", "=", 50))
	fmt.Println(xstrings.SubString("This is a test", 5, 6))
	fmt.Println(xstrings.SubStringStart("This is a test", 6))
	fmt.Println(xstrings.SubStringEnd("This is a test", 6))
	fmt.Println(xstrings.Truncate("This is a test", 6, false))
	fmt.Println(xstrings.Truncate("This is a test", 6, true))

	fmt.Println(xstrings.UUID4())
	fmt.Println(xstrings.ToTrain("ThisA_test"))
	fmt.Println(xstrings.ToSpinal("ThisA_test"))
	fmt.Println(xstrings.ToSnake("ThisA_test"))
	fmt.Println(xstrings.ToSnakeUpper("ThisA_test"))
	fmt.Println(xstrings.ToCamel("ThisA_test"))
	fmt.Println(xstrings.ToCamelLower("ThisA_test"))
	fmt.Println(xstrings.IsInSlice("test", []string{"a", "b", "tester", "testing", "test"}))
	fmt.Println(xstrings.PosInSlice("test", []string{"a", "b", "tester", "testing", "test"}))
	fmt.Println(xstrings.StringsBetween("[what is between]", "[", "]"))
	fmt.Println(xstrings.StringBetween("[what is between]", "[", "]"))
	fmt.Println(xstrings.Reverse("This is a test"))
	fmt.Println(xstrings.Format("The {} says {}", "cow", "MOO!"))
	fmt.Println(xstrings.RandStringWithLengthLimit(10))
	fmt.Println(xstrings.RandStringWithLengthLimit(20))
	fmt.Println(xstrings.RandStringWithLengthLimit(50))

	fmt.Println("")
	fmt.Println("SECURITY ***************************************************")
	fmt.Println("SHA 1                             >",xstrings.Sha1("InString"))
	fmt.Println("SHA 256                           >",xstrings.Sha256("InString"))
	fmt.Println("SHA 512                           >",xstrings.Sha512("InString"))
	fmt.Println("MASK String                       >",xstrings.MaskString("TESTING", "TESTINGTESTING", 4, 22))
	a, _ := xstrings.GenerateSecureString(30, xstrings.CharsASCII)
	fmt.Println("Generate Secure String            >", a)
	b, _ := xstrings.GenerateSecureString(30, "abcd")
	fmt.Println("Generate Secure String 2          >", b)

	fmt.Println(xstrings.Announcement("This is an announcement"))

	fmt.Println("")
	fmt.Println("JSON *******************************************************")
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
	fmt.Println("Non pretty JSON                   >", x)
	output, _ := xstrings.PrettyJson(x)
	fmt.Println("Pretty JSON                       >", output)
	fmt.Println("Compact JSON                      >", xstrings.CompactJSON(output))

	fmt.Println("")
	fmt.Println("UUID *******************************************************")
	g := xstrings.UUIDNewGen()
	for i := 0; i < 10; i++ {
		a := g.UUIDNewV4()
		fmt.Println("UUID: ", strconv.Itoa(i),"                     >", a.UUIDString())
	}

	fmt.Println("")
	fmt.Println("ORDINISE ***************************************************")
	fmt.Println("ToOrdinise(1)                     >", xstrings.ToOrdinise(1))
	fmt.Println("ToOrdinise(2)                     >", xstrings.ToOrdinise(2))
	fmt.Println("ToOrdinise(3)                     >", xstrings.ToOrdinise(3))
	fmt.Println("ToOrdinise(4)                     >", xstrings.ToOrdinise(4))
	fmt.Println("ToOrdinise(6)                     >", xstrings.ToOrdinise(6))
	fmt.Println("ToOrdinise(11)                    >", xstrings.ToOrdinise(11))
	fmt.Println("ToOrdinise(21)                    >", xstrings.ToOrdinise(21))
	fmt.Println("ToOrdinise(22)                    >", xstrings.ToOrdinise(22))
	fmt.Println("ToOrdinise(23)                    >", xstrings.ToOrdinise(23))
	fmt.Println("ToOrdinise(24)                    >", xstrings.ToOrdinise(24))

}
func CToGoString(c [16]byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}