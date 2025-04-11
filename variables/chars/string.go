package chars

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func rot13(in rune) rune {
	if in >= 'A' && in <= 'Z' {
		return (((in - 'A') + 13) % 26) + 'A'
	}
	if in >= 'a' && in <= 'z' {
		return (((in - 'a') + 13) % 26) + 'a'
	}

	return in
}

func Example3() {
	fmt.Println("Examples 3	")

	var s string
	s = "Hello, world!"
	fmt.Println(s)

	z := "Hello, \n\t\"world\" \\"
	fmt.Println(z)

	x := `Hello, 
	"world" \`
	fmt.Println(x)

	t := "👋 🌎"
	fmt.Println(t)

	y := s + " " + t
	fmt.Println(y)

	b := s[0]
	b2 := s[4]
	fmt.Println(s, b, string(b), b2, string(b2))
	s2 := s[0:5]  // Hello
	s3 := s[7:12] // world
	s4 := s[:5]   // Hello
	s5 := s[7:]   // world!
	fmt.Println(s, s2, s3, s4, s5)
	fmt.Println(s, len(s), s2, len(s2), s3, len(s3))

	t1 := t[:1]
	t2 := t[2:]
	fmt.Println(t, len(t), t1, len(t1), t2, len(t2))

	var r rune = 127757
	r1 := s[0:7] + string(r)
	fmt.Println(r1)

	g := '🌎'
	r2 := s[0:7] + string(g)
	fmt.Println(r2)

	s = "This is a test 123 🙂"
	s2 = strings.Map(rot13, s)
	fmt.Println(s2, "from 1st rotation")
	s3 = strings.Map(rot13, s2)
	fmt.Println(s3, "from 2nd rotation")

	s = "👋 🌎"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s), "from utf8 package")
}
