package option

// Option
// 支持泛型的 Option 设计模式, 避免在代码中定义很多类似的结构体
// 一般情况下 T 应该是一个结构体
type Option[T any] func(t *T)

// Apply
// @Description: 将 opts 应用在 t 之上
// @Author: Edv Chen <edvcc72@gmail.com>
// @param t
// @param opts
func Apply[T any](t *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt(t)
	}
}

// OptionErr
// 与 Option 相同，但会返回一个 error
// 除非在设计 option 模式时需要进行一些校验，否则应优先使用 Option
type OptionErr[T any] func(t *T) error

// ApplyErr
// @Description: 将 opts 应用在 t 之上，如果 opts 中任何一个返回 error，那么它会中断并且返回 error
// @Author: Edv Chen <edvcc72@gmail.com>
// @param t
// @param opts
// @return error
func ApplyErr[T any](t *T, opts ...OptionErr[T]) error {
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return err
		}
	}
	return nil
}
