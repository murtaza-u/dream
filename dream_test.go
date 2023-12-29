package dream_test

import (
	"fmt"
	"time"

	"github.com/murtaza-u/dream"
)

func ExampleStore() {
	s := dream.New()
	s.Put("foo", "bar")
	s.Put("blah", 100)

	fmt.Printf("foo=%s\n", s.GetString("foo"))

	s.Delete("blah")
	s.Delete("blah")
	fmt.Printf("blah=%d\n", s.GetInt("blah"))

	// Output:
	// foo=bar
	// blah=0
}

func ExampleWithCleanUp() {
	s := dream.New(dream.WithCleanUp(time.Millisecond))
	defer s.StopCleanUp()

	s.Put("foo", true)
	s.Put("bar", true)
	s.Put("blah", true)

	time.Sleep(time.Millisecond * 2)

	fmt.Printf("foo=%v\n", s.GetBool("foo"))
	fmt.Printf("bar=%v\n", s.GetBool("bar"))
	fmt.Printf("blah=%v\n", s.GetBool("blah"))

	// Output:
	// foo=false
	// bar=false
	// blah=false
}
