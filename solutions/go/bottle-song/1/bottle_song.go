// Package bottlesong реализует функциональность для генерации текста песни "Ten Green Bottles"
package bottlesong

import (
	"fmt"
	"strings"
)

// Verse представляет один куплет песни
// Count - количество бутылок
// NumberWord - словесное представление количества
// BottleWord - слово "bottle" в правильном числе
// Verb - глагол для действия
// NextCount - количество бутылок в следующем куплете
// NextNumber - словесное представление следующего количества
// NextBottle - слово "bottle" в правильном числе для следующего куплета
type Verse struct {
	Count        int
	NumberWord   string
	BottleWord   string
	Verb         string
	NextCount    int
	NextNumber   string
	NextBottle   string
}

// NewVerse создает новый куплет для заданного количества бутылок
// count - количество бутылок
func NewVerse(count int) Verse {
	v := Verse{Count: count}
	
	// Заполняем данные для текущего количества
	v.fillCurrent()
	
	// Заполняем данные для следующего количества
	if count > 0 {
		v.NextCount = count - 1
		v.NextNumber = getNumberWord(v.NextCount)
		v.NextBottle = getBottleWord(v.NextCount)
	} else {
		// Для случая с 0 бутылками
		v.NextCount = 10
		v.NextNumber = getNumberWord(v.NextCount)
		v.NextBottle = getBottleWord(v.NextCount)
	}
	
	return v
}

// fillCurrent заполняет данные для текущего количества бутылок
func (v *Verse) fillCurrent() {
	v.NumberWord = getNumberWord(v.Count)
	v.BottleWord = getBottleWord(v.Count)
	v.Verb = getVerb(v.Count)
}

// String возвращает строковое представление куплета
// String возвращает строковое представление куплета
func (v Verse) String() string {
	// Для случая с 0 бутылками возвращаем пустую строку
	if v.Count == 0 {
		return ""
	}
	
	// Определяем правильную форму слова "bottle" для фразы о падении
	fallingBottleWord := "bottle"
	//if v.Count  1 {
	//	fallingBottleWord = "bottles"
	//}
	
	// Форматируем куплет согласно требованиям
	return fmt.Sprintf("%s green %s hanging on the wall,\n%s green %s hanging on the wall,\nAnd if one green %s should accidentally fall,\nThere'll be %s green %s hanging on the wall.",
		Title(v.NumberWord), v.BottleWord, Title(v.NumberWord), v.BottleWord,
		fallingBottleWord, v.NextNumber, v.NextBottle)
}

// getNumberWord возвращает словесное представление числа
// count - количество бутылок
func getNumberWord(count int) string {
	// Для нуля возвращаем "no more"
	if count == 0 {
		return "no"
	}
	// Преобразуем числа в слова
	switch count {
	case 10:
		return "ten"
	case 9:
		return "nine"
	case 8:
		return "eight"
	case 7:
		return "seven"
	case 6:
		return "six"
	case 5:
		return "five"
	case 4:
		return "four"
	case 3:
		return "three"
	case 2:
		return "two"
	case 1:
		return "one"
	default:
		return ""
	}
}

// getBottleWord возвращает правильную форму слова "bottle"
// count - количество бутылок
func getBottleWord(count int) string {
	// Для одной бутылки возвращаем "bottle", для остальных "bottles"
	if count == 1 {
		return "bottle"
	}
	return "bottles"
}

// getVerb возвращает правильный глагол для действия
// count - количество бутылок
func getVerb(count int) string {
	// Для нуля возвращаем фразу о покупке новых бутылок
	if count == 0 {
		return "Go to the store and buy some more"
	}
	// Для остальных случаев возвращаем фразу о падении бутылки
	return "Take one down and pass it around"
}



// Recite возвращает куплеты песни "Ten Green Bottles"
// startBottles - начальное количество бутылок (1-10)
// takeDown - количество куплетов для исполнения (1-10)
// Recite возвращает куплеты песни "Ten Green Bottles"
// startBottles - начальное количество бутылок (1-10)
// takeDown - количество куплетов для исполнения (1-10)
func Recite(startBottles, takeDown int) []string {
	// Проверяем корректность входных параметров
	if startBottles < 1 || startBottles > 10 || takeDown < 1 || takeDown > startBottles {
		return nil
	}
	
	var verses []string
	// Исполняем куплеты от startBottles до takeDown
	for i := 0; i < takeDown; i++ {
		currentBottles := startBottles - i
		verse := NewVerse(currentBottles)
		// Разбиваем куплет на строки
		verseLines := strings.Split(verse.String(), "\n")
		verses = append(verses, verseLines...)
		//// Добавляем пустую строку между куплетами, кроме последнего
		if i < takeDown-1 {
			verses = append(verses, "")
		}
	}
	
	return verses
}


