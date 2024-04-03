package result

import (
	"fmt"
)

type Result[T any] struct {
	Ok  T
	Err error
}

func (r Result[T]) Raise() T {
	if r.Err != nil {
		panic(r)
	}
	return r.Ok
}

func Catch[T any](f func(error) T) T {
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()
	if r := recover(); r != nil {
		return f(r.(Result[T]).Err)
	}
	return f(nil)
}

func (r Result[T]) Unwrap() (T, error) {
	return r.Ok, r.Err
}

func (r Result[T]) Map(f func(T) T) Result[T] {
	if r.Err != nil {
		return Result[T]{Err: r.Err}
	}
	return Ok(f(r.Ok))
}

func (r Result[T]) Tap(f func(T)) Result[T] {
	if r.Err != nil {
		return Result[T]{Err: r.Err}
	}
	f(r.Ok)
	return r
}

func (r Result[T]) Print() Result[T] {
	if r.Err != nil {
		fmt.Println(r.Err)
	} else {
		fmt.Println(r.Ok)
	}
	return r
}

func Ok[T any](v T) Result[T] {
	return Result[T]{Ok: v}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func Bind[T, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if r.Err != nil {
		return Result[U]{Err: r.Err}
	}
	return f(r.Ok)
}

func Chain[T, U any](f func(T) (U, error)) func(Result[T]) Result[U] {
	return func(r Result[T]) Result[U] {
		if r.Err != nil {
			return Result[U]{Err: r.Err}
		}
		return Wrap(f(r.Ok))
	}
}

func Wrap[T any](ok T, err error) Result[T] {
	return Result[T]{Ok: ok, Err: err}
}
