package twelve

import (
	"fmt"
	"strings"
)

func Verse(i int) string {
		if i < 1 || i > 12 {
		return ""
	}

	// Порядковые числительные для дней
	ordinals := []string{
		"first", "second", "third", "fourth", "fifth", "sixth",
		"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
	}

	// Подарки для каждого дня
	gifts := []string{
		"a Partridge in a Pear Tree",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming",
	}

	// Начало строки (одинаковое для всех куплетов)
	start := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", ordinals[i-1])

	// Собираем список подарков для данного дня (в обратном порядке)
	var giftList []string
	for i := i; i > 0; i-- {
		giftList = append(giftList, gifts[i-1])
	}

	// Форматируем список подарков
	var giftsStr string
	if i == 1 {
		// Для первого дня просто подарок
		giftsStr = giftList[0]
	} else {
		// Для остальных дней добавляем "and" перед последним подарком
		// Все подарки, кроме последнего, разделяем запятыми
		allButLast := strings.Join(giftList[:len(giftList)-1], ", ")
		last := giftList[len(giftList)-1]
		giftsStr = fmt.Sprintf("%s, and %s", allButLast, last)
	}

	return start + giftsStr + "."
}


func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")

	//panic("Please implement the Song function")
}
