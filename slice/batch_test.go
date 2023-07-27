package slice

import (
	"reflect"
	"testing"
)

func TestBatch_GetChunks(t *testing.T) {
	type testCase[T any] struct {
		name string
		b    Batch[T]
		size int
		want [][]T
	}
	tests := []testCase[int]{
		{
			name: "wrong input",
			b:    NewBatch[int](1, 2, 3),
			size: 0,
			want: nil,
		},
		{
			name: "single",
			b:    NewBatch[int](1, 2, 3),
			size: 5,
			want: [][]int{
				{
					1, 2, 3,
				},
			},
		},
		{
			name: "multiple",
			b:    NewBatch[int](1, 2, 3),
			size: 2,
			want: [][]int{
				{
					1, 2,
				},
				{
					3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetChunks(tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
