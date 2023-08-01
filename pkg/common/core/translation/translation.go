package translation

import (
	"fmt"
	"strings"
)

func Translation(cyr string) string {
	result := ""
	translat := map[string]string{
		"ё": "yo", "ж": "zh", "х": "kh", "ц": "ts", "ч": "ch", "щ": "shh", "ш": "sh",
		"э": "eh", "ю": "yu", "я": "ya", "а": "a", "б": "b", "в": "v", "г": "g", "д": "d",
		"е": "e", "з": "z", "и": "i", "й": "j", "к": "k", "л": "l", "м": "m", "н": "n",
		"о": "o", "п": "p", "р": "r", "с": "s", "т": "t", "у": "u", "ф": "f", "Ё": "Yo",
		"Ж": "Zh", "Х": "Kh", "Ц": "Ts", "Ч": "Ch", "Щ": "Shh", "Ш": "Sh", "Э": "Eh", "Ю": "Yu",
		"Я": "Ya", "А": "A", "Б": "B", "В": "V", "Г": "G", "Д": "D", "Е": "E", "З": "Z", "И": "I", "Й": "J",
		"К": "K", "Л": "L", "М": "M", "Н": "N", "О": "O", "П": "P", "Р": "R", "С": "S",
		"Т": "T", "У": "U", "Ф": "F", "0": "0", "1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9",
	}

	for _, s := range cyr {
		symbol := fmt.Sprintf("%c", s)
		if val, ok := translat[symbol]; ok {
			result += val
		} else {
			result += "_"
		}
	}

	return strings.ToLower(result)
}
