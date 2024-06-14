package containers

type List[T any] interface {
	Insert(elem T)
	Traverse(f func(v any))
	IsEmpty() bool
	Size() int
	At(idx int) (T, error)
	DeleteAt(idx int) error
	InsertFront(t T)
	InsertAt(idx int, t T) error
}
