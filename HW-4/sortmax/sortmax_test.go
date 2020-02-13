package sortmax

import (
	"reflect"
	"testing"
)

func TestGetmaxValue(t *testing.T) {

	GetmaxValueTest := []struct {
		idTest string
		inData []interface{}
		result interface{}
		err    error
	}{
		{idTest: "floatTest", inData: []interface{}{10.45, 40.32, 50.00, 100500.00, 100.00, 500.76}, result: 100500.00},

		{idTest: "intTest", inData: []interface{}{10, 30, 55, 1341, 128, 576}, result: 1341},

		{idTest: "stringTest", inData: []interface{}{"Alice", "Bob", "Tomas", "Bendjamin", "Maks", "Henk", "May"}, result: "Bendjamin"},

		{idTest: "emptyTest", inData: []interface{}{}, err: err},
	}

	for _, tt := range GetmaxValueTest {

		t.Run(tt.idTest, func(t *testing.T) {

			got, err := GetMaxValue(tt.inData, GetComparator(tt.inData))

			if err != nil {
								if err != tt.err {
					t.Errorf("GetmaxValue() error = %v, wantErr %v", err, tt.err)
				}
			}

			if !reflect.DeepEqual(got, tt.result) {
				t.Errorf("%#v got %v want %v", tt.idTest, got, tt.result)
			}
		})
	}
}
