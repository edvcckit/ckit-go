package option

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApply(t *testing.T) {
	u := &User{}
	Apply[User](u, WithName("chen"), WithAge(18))
	assert.Equal(t, u, &User{name: "chen", age: 18})
}

func TestApplyErr(t *testing.T) {
	u := &User{}
	err := ApplyErr[User](u, WithNameErr("chen"), WithAgeErr(18))
	require.NoError(t, err)
	assert.Equal(t, u, &User{name: "chen", age: 18})

	err = ApplyErr[User](u, WithNameErr(""), WithAgeErr(18))
	assert.Equal(t, errors.New("name 不能为空"), err)

	err = ApplyErr[User](u, WithNameErr("chen"), WithAgeErr(-1))
	assert.Equal(t, errors.New("age 应该是正数"), err)
}

func ExampleApply() {
	u := &User{}
	Apply[User](u, WithName("chen"), WithAge(18))
	fmt.Println(u)
	// Output:
	// &{chen 18}
}

func ExampleApplyErr() {
	u := &User{}
	err := ApplyErr[User](u, WithNameErr("chen"), WithAgeErr(18))
	fmt.Println(err)
	fmt.Println(u)

	err = ApplyErr[User](u, WithNameErr(""), WithAgeErr(18))
	fmt.Println(err)
	err = ApplyErr[User](u, WithNameErr("chen"), WithAgeErr(-1))
	fmt.Println(err)

	// Output:
	// <nil>
	// &{chen 18}
	// name 不能为空
	// age 应该是正数
}

func WithName(name string) Option[User] {
	return func(u *User) {
		u.name = name
	}
}

func WithAge(age int) Option[User] {
	return func(u *User) {
		u.age = age
	}
}

func WithNameErr(name string) OptionErr[User] {
	return func(u *User) error {
		if name == "" {
			return errors.New("name 不能为空")
		}
		u.name = name
		return nil
	}
}

func WithAgeErr(age int) OptionErr[User] {
	return func(u *User) error {
		if age <= 0 {
			return errors.New("age 应该是正数")
		}
		u.age = age
		return nil
	}
}

type User struct {
	name string
	age  int
}
