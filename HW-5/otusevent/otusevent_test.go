package otusevent

import (
	"bytes"
	"reflect"
	"testing"
)

func TestLogOtusEvent(t *testing.T) {

	LogOtusEventTest := []struct {
		idTest    string
		otusEvent OtusEvent
		result    string
	}{
		{idTest: "HwSubmittedTest", otusEvent: &HwSubmitted{ID: 3456, Code: "My Code", Comment: "please take a look at my homework"}, result: data + " submitted 3456 \"please take a look at my homework\"\n"},

		{idTest: "HwAcceptedTest", otusEvent: &HwAccepted{ID: 3456, Grade: 4}, result: data + " accepted 3456 4\n"},
	}

	for _, tt := range LogOtusEventTest {

		t.Run(tt.idTest, func(t *testing.T) {

			var b bytes.Buffer

			LogOtusEvent(tt.otusEvent, &b)

			got := b.String()

			if !reflect.DeepEqual(got, tt.result) {
				t.Errorf("%#v got %q want %q", tt.idTest, got, tt.result)
			}
		})
	}
}
