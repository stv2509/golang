package unpacker

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCaseType struct {
	description    string
	testData       string
	expectedResult string
	expectedError  error
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	testCases := []TestCaseType{
		{
			description:    "test data №1",
			testData:       "a4bc2d5e",
			expectedResult: "aaaabccddddde",
		},
		{
			description:    "test data №2",
			testData:       "abcd",
			expectedResult: "abcd",
		},
		{
			description:   "test data №3",
			testData:      "45",
			expectedError: errors.New("некорректная строка"),
		},
		{
			description:    "test data №4",
			testData:       "",
			expectedResult: "",
		},
		{
			description:    "test data №5",
			testData:       `qwe\4\5`,
			expectedResult: "qwe45",
		},
		{
			description:    "test data №6",
			testData:       `qwe\45`,
			expectedResult: "qwe44444",
		},
		{
			description:    "test data №7",
			testData:       `qwe\\5`,
			expectedResult: `qwe\\\\\`,
		},
	}
	for _, testCase := range testCases {
		result, err := Unpack(testCase.testData)
		if testCase.expectedError != nil {
			assert.Equal(testCase.expectedError, err, testCase.description)
		} else {
			assert.Equal(testCase.expectedResult, result, testCase.description)
		}
	}
}