package unpacker

import (
	"errors"
	"strconv"
	"strings"
)

var errRune = errors.New("The string must begin with a letter")

// Unpack Распаковка строки
func Unpack(str string) (string, error) {

	var err error
	ch := []rune(str) //Переводим строку в руны

	if (len(str) == 0) || ('0' <= ch[0] && ch[0] <= '9') { //Если строка пустая или начинается с цифры, возвращаем ошибку

		err := errRune
		return str, err
	}

	for i, v := range ch {

		if ('0' <= v && v <= '9') && ch[i-1] != '\\' { //Ищем цифру в строке и отсутствие '\' перед ней

			runeSt := string(ch[i-1])                                  //От индекса найденной цифры, отнимаем "1", получаем индекс руны перед цифрой
			vInt, _ := strconv.Atoi(string(v))                         //Преобразуем тип руны, содержащей цифру в тип int
			runeStNew := strings.Repeat(runeSt, vInt)                  //Повторяем найденную руну int раз
			str = strings.Replace(str, runeSt+string(v), runeStNew, 1) // Заменяем выражение руна+цифра на runeStNew, и перезаписываем str

		} else if v == '\\' && ch[i+1] != '\\' { // Если rune содержит '\' и не содержит '\' после

			if (int(i+2) < len(ch)) && ('0' <= ch[i+2] && ch[i+2] <= '9') { // После '\' вторая rune цифра и при поиске мы не выходим за пределы slice
				runeSt := string(ch[i+1])                 // Первая цифра после '\'
				vInt, _ := strconv.Atoi(string(ch[i+2]))  // Вторая цифра после '\', переводим в int
				runeStNew := strings.Repeat(runeSt, vInt) // Повторить первую цифру int раз
				str = strings.Replace(str, string(ch[i])+runeSt+string(ch[i+2]), runeStNew, 1)

			} else {

				str = strings.Replace(str, string(ch[i]), "", 1)

			}

		} else if v == '\\' && ch[i+1] == '\\' { // Если rune содержит '\' и содержит '\' после

			if (int(i+2) < len(ch)) && ('0' <= ch[i+2] && ch[i+2] <= '9') { // После '\' вторая rune цифра и при поиске мы не выходим за пределы slice

				runeSt := "\\"
				vInt, _ := strconv.Atoi(string(ch[i+2]))  // Вторая цифра после '\', переводим в int
				runeStNew := strings.Repeat(runeSt, vInt) // Повторить '\' int раз
				str = strings.Replace(str, runeSt+string(ch[i+2]), runeStNew, 1)
			}
		}
	}

	return str, err
}
