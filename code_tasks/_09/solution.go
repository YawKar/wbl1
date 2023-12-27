package _09

func Connect[T any](in <-chan T, out chan<- T) {
	defer close(out)
	for v := range in {
		out <- v
	}
}

func MapConnect[T, U any](in <-chan T, f func(T) U, out chan<- U) {
	defer close(out)
	for v := range in {
		out <- f(v)
	}
}
