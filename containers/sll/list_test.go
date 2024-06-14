package sll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Insert(t *testing.T) {
	type args[T any] struct {
		elem T
	}
	type testCase[T any] struct {
		name      string
		l         SLList[T]
		args      args[T]
		want      T
		wantedPos int
	}
	tests := []testCase[int]{
		{
			name: "insert to empty list",
			l:    SLList[int]{},
			args: args[int]{
				elem: 23,
			},

			want:      23,
			wantedPos: 0,
		},
		{
			name: "insert to list with one element",
			l: SLList[int]{
				head: &node[int]{
					next: nil,
					val:  11,
				},
				size: 1,
			},
			args: args[int]{
				elem: 12,
			},
			want:      12,
			wantedPos: 1,
		},
		{
			name: "insert to list with two elements",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
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
			tt.l.Insert(tt.args.elem)

			inserted, err := tt.l.At(tt.wantedPos)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, inserted)
		})
	}
}

func TestList_IsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name string
		l    SLList[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "check empty list",
			l:    SLList[int]{},
			want: true,
		},
		{
			name: "check list with head",
			l: SLList[int]{
				head: &node[int]{
					next: nil,
					val:  11,
				},
				size: 1,
			},
			want: false,
		},
		{
			name: "check list with body",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			want: false,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, tt.want, tt.l.IsEmpty(), "IsEmpty()")
		})
	}
}

func TestList_At(t *testing.T) {
	type args struct {
		idx int
	}
	type testCase[T any] struct {
		name    string
		l       SLList[T]
		args    args
		want    T
		wantErr error
	}
	tests := []testCase[int]{
		{
			name: "empty list",
			l:    SLList[int]{},
			args: args{
				idx: 1,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "non-empty list",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			args: args{
				idx: 1,
			},
			want:    121,
			wantErr: nil,
		},
		{
			name: "non-empty list negative index",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			args: args{
				idx: -1,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "non-empty list index is equal to size",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			args: args{
				idx: 2,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
		{
			name: "non-empty list index is greater than size",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			args: args{
				idx: 3,
			},
			want:    0,
			wantErr: ErrIndexIsOutOfSize,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.l.At(tt.args.idx)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "At(%v)", tt.args.idx)
		})
	}
}

func TestList_PushFront(t *testing.T) {
	type args[T any] struct {
		t T
	}
	type testCase[T any] struct {
		name string
		l    SLList[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "empty list",
			l:    SLList[int]{},
			args: args[int]{123},
			want: 123,
		},
		{
			name: "non-empty list",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			args: args[int]{1},
			want: 1,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.l.InsertFront(tt.args.t)
			inserted, _ := tt.l.At(0)
			assert.Equal(t, tt.want, inserted)
		})
	}
}

func TestList_InsertAt(t *testing.T) {
	type args[T any] struct {
		idx int
		t   T
	}
	type testCase[T any] struct {
		name      string
		l         SLList[T]
		args      args[T]
		want      T
		wantErr   error
		wantedPos int
	}
	tests := []testCase[int]{
		{
			name: "negative index",
			l:    SLList[int]{},
			args: args[int]{
				idx: -1,
				t:   11,
			},
			want:      0,
			wantErr:   ErrIndexIsOutOfSize,
			wantedPos: -1,
		},
		{
			name: "empty list 0 index",
			l:    SLList[int]{},
			args: args[int]{
				idx: 0,
				t:   11,
			},
			want:      11,
			wantErr:   nil,
			wantedPos: 0,
		},
		{
			name: "non-empty list 0 index",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			args: args[int]{
				idx: 0,
				t:   12,
			},
			want:      12,
			wantErr:   nil,
			wantedPos: 0,
		},
		{
			name: "non-empty list",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			args: args[int]{
				idx: 1,
				t:   13,
			},
			want:      13,
			wantErr:   nil,
			wantedPos: 1,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.l.InsertAt(tt.args.idx, tt.args.t)
			got, _ := tt.l.At(tt.wantedPos)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestList_DeleteAt(t *testing.T) {
	type args struct {
		idx int
	}
	type testCase[T any] struct {
		name      string
		l         SLList[T]
		args      args
		wantErr   error
		wantedPos int
		want      T
	}
	tests := []testCase[int]{
		{
			name: "empty list",
			l:    SLList[int]{},
			args: args{
				idx: 1,
			},
			wantErr:   ErrIndexIsOutOfSize,
			wantedPos: 1,
			want:      0,
		},
		{
			name: "index out of range",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			args: args{
				idx: 3,
			},
			wantErr:   ErrIndexIsOutOfSize,
			wantedPos: 3,
			want:      0,
		},
		{
			name: "non-empty list valid index",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: &node[int]{
							next: nil,
							val:  100,
						},
						val: 121,
					},
					val: 11,
				},
				size: 3,
			},
			args: args{
				idx: 1,
			},
			wantErr:   nil,
			wantedPos: 1,
			want:      100,
		},
		{
			name: "non-empty list zero index",
			l: SLList[int]{
				head: &node[int]{
					next: &node[int]{
						next: nil,
						val:  121,
					},
					val: 11,
				},
				size: 2,
			},
			args: args{
				idx: 0,
			},
			wantErr:   nil,
			wantedPos: 0,
			want:      121,
		},
	}

	t.Parallel()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.l.DeleteAt(tt.args.idx)
			assert.ErrorIs(t, err, tt.wantErr)
			got, _ := tt.l.At(tt.wantedPos)
			assert.Equal(t, tt.want, got)
		})
	}
}
