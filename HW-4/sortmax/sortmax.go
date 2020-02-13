package sortmax

import (
	"errors"
)

var err = errors.New("Interface is empty")

// GetComparator - sort string, int, float64
func GetComparator(list []interface{}) func(i, j int) bool {

	return func(i, j int) bool {

		switch list[len(list)-1].(type) {

		case string:
			return len(list[i].(string)) > len(list[j].(string))

		case int:
			return list[i].(int) > list[j].(int)

		case float64:
			return list[i].(float64) > list[j].(float64)

		}
		return false
	}
}

//GetMaxValue maximum get
func GetMaxValue(slice []interface{}, comparator func(value1, value2 int) bool) (interface{}, error) {

	if len(slice) == 0 {
		return nil, err
	}

	var max int

	for idx := range slice {

		if max != 0 {

			if comparator(idx, max) {

				max = idx
			}

		} else {

			max = idx
		}
	}
	return slice[max], nil

}