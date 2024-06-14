package dll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDLList_Insert(t *testing.T) {
	type args[T any] struct {
		elem T
	}
	type testCase[T any] struct {
		name      string
		l         func() DLList[T]
		args      args[T]
		want      T
		wantedPos int
	}
	tests := []testCase[int]{
		{
			name: "insert to empty list",
			l: func() DLList[int] {
				return DLList[int]{}
			},
			args: args[int]{
				elem: 23,
			},

			want:      23,
			wantedPos: 0,
		},
		{
			name: "insert to list with one element",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args[int]{
				elem: 12,
			},
			want:      12,
			wantedPos: 1,
		},
		{
			name: "insert to list with two elements",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: nil,
							val:  121,
						},
						val: 11,
					},
					size: 2,
				}
				l.tail = l.head.next
				return l
			},

			args: args[int]{
				elem: 12,
			},
			want:      12,
			wantedPos: 2,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			l.Insert(tt.args.elem)

			inserted, err := l.At(tt.wantedPos)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, inserted)
		})
	}
}

func TestDLList_GetTail(t *testing.T) {
	type testCase[T any] struct {
		name    string
		l       func() DLList[T]
		want    T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name: "tail in empty list",
			l: func() DLList[int] {
				return DLList[int]{}
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "tail equals head",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			want:    11,
			wantErr: nil,
		},
		{
			name: "just a regular tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: nil,
							val:  121,
						},
						val: 11,
					},
					size: 2,
				}
				l.tail = l.head.next
				return l
			},
			want:    121,
			wantErr: nil,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			got, err := l.GetTail()
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetTail()")
		})
	}
}

func TestDLList_IsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name string
		l    func() DLList[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "check empty list",
			l: func() DLList[int] {
				return DLList[int]{}
			},
			want: true,
		},
		{
			name: "check list where head equals tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			want: false,
		},
		{
			name: "check list with head and tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: nil,
							val:  121,
						},
						val: 11,
					},
					size: 2,
				}
				l.tail = l.head.next
				return l
			},
			want: false,
		},
		{
			name: "check list with head, body and tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: &node[int]{
								next: nil,
								val:  13,
							},
							val: 12,
						},
						val: 11,
					},
					size: 3,
				}
				l.tail = l.head.next.next
				return l
			},
			want: false,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			assert.Equalf(t, tt.want, l.IsEmpty(), "IsEmpty()")
		})
	}
}

func TestDLList_Size(t *testing.T) {
	type testCase[T any] struct {
		name string
		l    func() DLList[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "check empty list",
			l: func() DLList[int] {
				return DLList[int]{}
			},
			want: 0,
		},
		{
			name: "check list where head equals tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			want: 1,
		},
		{
			name: "check list with head and tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: nil,
							val:  121,
						},
						val: 11,
					},
					size: 2,
				}
				l.tail = l.head.next
				return l
			},
			want: 2,
		},
		{
			name: "check list with head, body and tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: &node[int]{
								next: nil,
								val:  13,
							},
							val: 12,
						},
						val: 11,
					},
					size: 3,
				}
				l.tail = l.head.next.next
				return l
			},
			want: 3,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			assert.Equalf(t, tt.want, l.Size(), "Size()")
		})
	}
}

func TestDLList_At(t *testing.T) {
	type args struct {
		idx int
	}
	type testCase[T any] struct {
		name    string
		l       func() DLList[T]
		args    args
		want    T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name: "check empty list",
			l: func() DLList[int] {
				return DLList[int]{}
			},
			args: args{
				idx: 1,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "check negative index",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args{
				idx: -1,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "check index is equal to size",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args{
				idx: 1,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "check index is bigger than size",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args{
				idx: 2,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "check list with head equal to tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args{
				idx: 0,
			},
			want:    11,
			wantErr: nil,
		},
		{
			name: "check list with head and tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: nil,
							val:  121,
						},
						val: 11,
					},
					size: 2,
				}
				l.tail = l.head.next
				return l
			},
			args: args{
				idx: 1,
			},
			want:    121,
			wantErr: nil,
		},
		{
			name: "check list with head, body and tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: &node[int]{
								next: nil,
								val:  13,
							},
							val: 12,
						},
						val: 11,
					},
					size: 3,
				}
				l.tail = l.head.next.next
				return l
			},
			args: args{
				idx: 1,
			},
			want:    12,
			wantErr: nil,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			got, err := l.At(tt.args.idx)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "At(%v)", tt.args.idx)
		})
	}
}

func TestDLList_DeleteAt(t *testing.T) {
	type args struct {
		idx int
	}
	type testCase[T any] struct {
		name    string
		l       func() DLList[T]
		args    args
		want    T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name: "check empty list",
			l: func() DLList[int] {
				return DLList[int]{}
			},
			args: args{
				idx: 1,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "check negative index",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args{
				idx: -1,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "check index is equal to size",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args{
				idx: 1,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "check index is bigger than size",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args{
				idx: 2,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "check list with head equal to tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args{
				idx: 0,
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "check list with head and tail, del head",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: nil,
							val:  12,
						},
						val: 11,
					},
					size: 2,
				}
				l.tail = l.head.next
				return l
			},
			args: args{
				idx: 0,
			},
			want:    12,
			wantErr: nil,
		},
		{
			name: "check list with head, body and tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: &node[int]{
								next: nil,
								val:  13,
							},
							val: 12,
						},
						val: 11,
					},
					size: 3,
				}
				l.tail = l.head.next.next
				return l
			},
			args: args{
				idx: 1,
			},
			wantErr: nil,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			err := l.DeleteAt(tt.args.idx)
			assert.ErrorIs(t, err, tt.wantErr)
			if err == nil {
				got, _ := l.At(tt.args.idx)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDLList_PushFront(t *testing.T) {
	type args[T any] struct {
		t T
	}
	type testCase[T any] struct {
		name string
		l    func() DLList[T]
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "push in empty list",
			l: func() DLList[int] {
				return DLList[int]{}
			},
			args: args[int]{
				t: 18,
			},
		},
		{
			name: "push in list where head equal to tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args[int]{
				t: 18,
			},
		},
		{
			name: "push in list with head and tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: nil,
							val:  12,
						},
						val: 11,
					},
					size: 2,
				}
				l.tail = l.head.next
				return l
			},
			args: args[int]{
				t: 18,
			},
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			l.InsertFront(tt.args.t)
			got, _ := l.At(0)
			assert.Equal(t, tt.args.t, got)
		})
	}
}

func TestDLList_InsertAt(t *testing.T) {
	type args[T any] struct {
		idx int
		t   T
	}
	type testCase[T any] struct {
		name       string
		l          func() DLList[T]
		args       args[T]
		wantErr    error
		want       T
		wantSubErr error
	}
	tests := []testCase[int]{
		{
			name: "negative index",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args[int]{
				idx: -1,
				t:   18,
			},
			want:       0,
			wantErr:    ErrIndexIsOutOfSize,
			wantSubErr: nil,
		},
		{
			name: "index is bigger than size",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args[int]{
				idx: 2,
				t:   18,
			},
			wantErr:    nil,
			want:       18,
			wantSubErr: nil,
		},
		{
			name: "index is zero",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			args: args[int]{
				idx: 0,
				t:   18,
			},
			wantErr:    nil,
			want:       18,
			wantSubErr: nil,
		},
		{
			name: "index is zero, list is empty",
			l: func() DLList[int] {
				return DLList[int]{}
			},
			args: args[int]{
				idx: 0,
				t:   18,
			},
			wantErr:    nil,
			want:       18,
			wantSubErr: nil,
		},
		{
			name: "regular insertAt",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: &node[int]{
								next: nil,
								val:  13,
							},
							val: 12,
						},
						val: 11,
					},
					size: 3,
				}
				l.tail = l.head.next.next
				return l
			},
			args: args[int]{
				idx: 1,
				t:   18,
			},
			wantErr:    nil,
			want:       18,
			wantSubErr: nil,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			err := l.InsertAt(tt.args.idx, tt.args.t)
			assert.ErrorIs(t, err, tt.wantErr)
			if tt.wantErr == nil && err == nil {
				var indexVal int
				if tt.args.idx < l.size {
					indexVal = tt.args.idx
				} else {
					indexVal = l.size - 1
				}
				got, newErr := l.At(indexVal)
				assert.Equal(t, got, tt.want)
				assert.ErrorIs(t, newErr, tt.wantSubErr)
			}
		})
	}
}

func TestDLList_Reverse(t *testing.T) {
	type testCase[T any] struct {
		name string
		l    func() DLList[T]
	}
	tests := []testCase[int]{
		{
			name: "head and tail",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: nil,
							val:  12,
						},
						val: 11,
					},
					size: 2,
				}
				l.tail = l.head.next
				return l
			},
		},
		{
			name: "head, tail and body",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: &node[int]{
								next: nil,
								val:  13,
							},
							val: 12,
						},
						val: 11,
					},
					size: 3,
				}
				l.tail = l.head.next.next
				return l
			},
		},
		{
			name: "head, tail and body 2.0",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: &node[int]{
								next: &node[int]{
									next: nil,
									val:  14,
								},
								val: 13,
							},
							val: 12,
						},
						val: 11,
					},
					size: 4,
				}
				l.tail = l.head.next.next.next
				return l
			},
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			wasHead, wasTail := l.head.val, l.tail.val
			l.Reverse()
			nowHead, nowTail := l.head.val, l.tail.val
			assert.Equal(t, wasHead, nowTail)
			assert.Equal(t, wasTail, nowHead)
		})
	}
}

func TestDLList_Traverse(t *testing.T) {
	type args struct {
		f func(m []int) func(v any)
	}
	type testCase[T any] struct {
		name               string
		l                  func() DLList[T]
		args               args
		wantTraversedNodes []int
	}

	tests := []testCase[int]{
		{
			name: "traverse forward",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: &node[int]{
								next: &node[int]{
									next: nil,
									val:  14,
								},
								val: 13,
							},
							val: 12,
						},
						val: 11,
					},
					size: 4,
				}
				l.tail = l.head.next.next.next
				return l
			},
			args: args{
				f: func(m []int) func(v any) {
					pos := 0
					return func(v any) {
						m[pos] = v.(int)
						pos++
					}
				},
			},
			wantTraversedNodes: []int{11, 12, 13, 14},
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()

			traversedNodes := make([]int, l.size)

			l.Traverse(tt.args.f(traversedNodes))
			assert.Equal(t, tt.wantTraversedNodes, traversedNodes)
		})
	}
}

func TestDLList_Delete(t *testing.T) {
	type testCase[T any] struct {
		name     string
		l        func() DLList[T]
		wantErr  error
		wantTail T
	}
	tests := []testCase[int]{
		{
			name: "delete in empty list",
			l: func() DLList[int] {
				return DLList[int]{}
			},
			wantErr:  ErrIndexIsOutOfSize,
			wantTail: 0,
		},
		{
			name: "delete in ddl with one node",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: nil,
						val:  11,
					},
					size: 1,
				}
				l.tail = l.head
				return l
			},
			wantErr:  nil,
			wantTail: 0,
		},
		{
			name: "delete in ddl with two nodes",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: nil,
							val:  12,
						},
						val: 11,
					},
					size: 2,
				}
				l.tail = l.head.next
				return l
			},
			wantErr:  nil,
			wantTail: 11,
		},
		{
			name: "delete in ddl with three nodes",
			l: func() DLList[int] {
				l := DLList[int]{
					head: &node[int]{
						next: &node[int]{
							next: &node[int]{
								next: nil,
								val:  13,
							},
							val: 12,
						},
						val: 11,
					},
					size: 3,
				}
				l.tail = l.head.next.next
				return l
			},
			wantErr:  nil,
			wantTail: 12,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := tt.l()
			err := l.DeleteFromTail()
			assert.ErrorIs(t, err, tt.wantErr)
			if err == nil && l.size > 0 {
				nowTail, _ := l.GetTail()
				assert.Equal(t, nowTail, tt.wantTail)
			}

		})
	}
}
