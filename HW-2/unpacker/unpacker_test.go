package unpacker

import (
	"testing"
)

func TestUnpack(t *testing.T) {

	UnpackTests := []struct {
		idTest    string
		inParm    string
		hasUnpack string
		err       error
	}{
		{idTest: "test1", inParm: "a4bc2d5e", hasUnpack: "aaaabccddddde"},
		{idTest: "test2", inParm: "abcd", hasUnpack: "abcd"},
		{idTest: "test3", inParm: "45", err: errRune},
		{idTest: "test4", inParm: "", err: errRune},
		{idTest: "test5", inParm: `qwe\4\5`, hasUnpack: "qwe45"},
		{idTest: "test6", inParm: `qwe\45`, hasUnpack: "qwe44444"},
		{idTest: "test7", inParm: `qwe\\5`, hasUnpack: `qwe\\\\\`},
	}

	for _, tt := range UnpackTests {
		t.Run(tt.idTest, func(t *testing.T) {
			got, err := Unpack(tt.inParm)

			if err != nil {
				if err != tt.err {
					t.Errorf("%#v got %s want %s", tt.inParm, got, tt.err)
				}

			} else {

				if got != tt.hasUnpack {
					t.Errorf("%#v got %s want %s", tt.inParm, got, tt.hasUnpack)
				}
			}
		})

	}

}
