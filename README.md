<p align="center" width="100%">
  <strong>Dream</strong> is an in-memory key-value data store.
</p>

<p align="center" width="100%">
  <img width="50%" src="./artwork/logo.png">
</p>

---

<div align="center"><p>
  <a href="https://godoc.org/github.com/murtaza-u/dream">
    <img alt="GoDoc" src="https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge&logo=github&color=30b976&logoColor=D9E0EE&labelColor=302D41"/>
  </a>

  <a href="https://github.com/murtaza-u/dream/pulse">
    <img alt="Last commit" src="https://img.shields.io/github/last-commit/murtaza-u/dream?style=for-the-badge&logo=github&color=8bd5ca&logoColor=D9E0EE&labelColor=302D41"/>
  </a>

  <a href="https://github.com/murtaza-u/dream/blob/main/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/murtaza-u/dream?style=for-the-badge&logo=github&color=ee999f&logoColor=D9E0EE&labelColor=302D41" />
  </a>

  <a href="https://github.com/murtaza-u/dream/stargazers">
    <img alt="Stars" src="https://img.shields.io/github/stars/murtaza-u/dream?style=for-the-badge&logo=github&color=c69ff5&logoColor=D9E0EE&labelColor=302D41" />
  </a>

  <a href="https://github.com/murtaza-u/dream/issues">
    <img alt="Issues" src="https://img.shields.io/github/issues/murtaza-u/dream?style=for-the-badge&logo=bilibili&color=F5E0DC&logoColor=D9E0EE&labelColor=302D41" />
  </a>

  <a href="https://github.com/murtaza-u/dream">
    <img alt="Repo Size" src="https://img.shields.io/github/repo-size/murtaza-u/dream?color=%23DDB6F2&label=SIZE&logo=codesandbox&style=for-the-badge&logoColor=D9E0EE&labelColor=302D41" />
  </a>

  <a href="https://twitter.com/intent/follow?screen_name=murtaza_u_">
    <img alt="Follow on Twitter" src="https://img.shields.io/twitter/follow/murtaza_u_?style=for-the-badge&logo=twitter&color=8aadf3&logoColor=D9E0EE&labelColor=302D41" />
  </a>
</p></div>

## Usage

```sh
go get -u github.com/murtaza-u/dream
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/murtaza-u/dream"
)

type User struct {
	Name string
	Age  int
}

func main() {
	s := dream.New()
	// OR
	// With automatic clean up
	// s := dream.New(dream.WithCleanUp(time.Hour))

	s.Put("foo", "bar")
	s.Put("bar", true)
	s.Put("blah", 100)
	s.Put("user", User{
		Name: "Alice",
		Age:  30,
	})

	fmt.Printf("foo=%s\n", s.GetString("foo"))
	fmt.Printf("bar=%v\n", s.GetBool("bar"))
	fmt.Printf("blah=%d\n", s.GetInt("blah"))

	u, ok := s.Get("user").(User)
	if !ok {
		log.Fatal("failed to retrieve user")
	}
	fmt.Printf("user.name=%s | user.age=%d\n", u.Name, u.Age)
}
```

Full API reference: [GoDoc](https://godoc.org/github.com/murtaza-u/dream)
