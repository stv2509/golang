package topwords

import (
	"reflect"
	"testing"
)

func TestTop10(t *testing.T) {

	/*	assertCorrectMessage := func(t *testing.T, got, want []string) {
			t.Helper()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}

		t.Run("test1", func(t *testing.T) {
			got := Top10("111 222 111 444 222 111 333")
			want := []string{"111", "222"}
			assertCorrectMessage(t, got, want)
		})

		t.Run("test2", func(t *testing.T) {
			got := Top10(`giraffe cat rat cheetah dog monkey kitten  rat ferret lion
						rat pig giraffe hamster rat kangaroo elephant giraffe leopard
						lion monkey ostrich rat tiger giraffe lion turtle`)
			want := []string{"rat", "giraffe", "lion", "monkey"}
			assertCorrectMessage(t, got, want)
		})

	}*/

	Top10Test := []struct {
		idTest   string
		inString string
		hasTop   []string
	}{
		{idTest: "test1", inString: "111 222 111 444 222 111 333", hasTop: "111 222"},
		{idTest: "test2", inString: `giraffe cat rat cheetah dog monkey kitten  rat ferret lion
					rat pig giraffe hamster rat kangaroo elephant giraffe leopard 
					lion monkey ostrich rat tiger giraffe lion turtle`, hasTop: {"rat", "giraffe", "lion", "monkey"}},
	}

	for _, tt := range Top10Test {
		t.Run(tt.idTest, func(t *testing.T) {
			got := Top10(tt.inString)
			if !reflect.DeepEqual(got, tt.hasTop) {
				t.Errorf("%#v got %s want %s", tt.idTest, got, tt.hasTop)
			}
		})
	}
}
