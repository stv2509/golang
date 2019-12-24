package unpacker

import (
	"strconv"
	"strings"
)

func Unpack(str string) string {
	ch := []rune(str) //Переводим строку в руны
	for i, v := range ch {
		if ('0' <= v && v <= '9') && ch[i-1] != '\\' { //Ищем цифру в строке и отсутствие '\' перед ней
			rune_st := string(ch[i-1])                                    //От индекса найденной цифры, отнимаем "1", получаем индекс руны перед цифрой
			v_int, _ := strconv.Atoi(string(v))                           //Преобразуем тип руны, содержащей цифру в тип int
			rune_st_new := strings.Repeat(rune_st, v_int)                 //Повторяем найденную руну int раз
			str = strings.Replace(str, rune_st+string(v), rune_st_new, 1) // Заменяем выражение руна+цифра на rune_st_new, и перезаписываем str
		} else if v == '\\' && ch[i+1] != '\\' { // Если rune содержит '\' и не содержит '\' после
			if (int(i+2) < len(ch)) && ('0' <= ch[i+2] && ch[i+2] <= '9') { // После '\' вторая rune цифра и при поиске мы не выходим за пределы slice
				rune_st := string(ch[i+1])                    // Первая цифра после '\'
				v_int, _ := strconv.Atoi(string(ch[i+2]))     // Вторая цифра после '\', переводим в int
				rune_st_new := strings.Repeat(rune_st, v_int) // Повторить первую цифру int раз
				str = strings.Replace(str, string(ch[i])+rune_st+string(ch[i+2]), rune_st_new, 1)
			} else {
				str = strings.Replace(str, string(ch[i]), "", 1)
			}
		} else if v == '\\' && ch[i+1] == '\\' { // Если rune содержит '\' и содержит '\' после
			if (int(i+2) < len(ch)) && ('0' <= ch[i+2] && ch[i+2] <= '9') { // После '\' вторая rune цифра и при поиске мы не выходим за пределы slice
				rune_st := "\\"
				v_int, _ := strconv.Atoi(string(ch[i+2]))     // Вторая цифра после '\', переводим в int
				rune_st_new := strings.Repeat(rune_st, v_int) // Повторить '\' int раз
				str = strings.Replace(str, rune_st+string(ch[i+2]), rune_st_new, 1)
			}
		}
	}
	return str
}