// Пакет keyboard читает данные, введенные потльзоваттелем с клавиатуры
// в командной строке
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Функция GetFloat считывает числа с плавающей точкой
// и возвращает число с плавающей точкой в случае удачного считывания
// и ошибку в случае неудачи
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}
