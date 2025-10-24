package anagram

import (
	"strings"
)

// Detect находит анаграммы целевого слова среди кандидатов
func Detect(subject string, candidates []string) []string {
	var result []string
	subjectLower := strings.ToLower(subject)
	subjectFreq := getCharFrequency(subjectLower)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		
		// Пропускаем, если слово совпадает с целевым (без учета регистра)
		if candidateLower == subjectLower {
			continue
		}
		
		// Проверяем, является ли кандидат анаграммой
		if isAnagram(subjectFreq, candidateLower) {
			result = append(result, candidate)
		}
	}
	
	return result
}

// getCharFrequency возвращает частоту символов в строке
func getCharFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}
	return freq
}

// isAnagram проверяет, является ли строка анаграммой на основе частотного словаря
func isAnagram(subjectFreq map[rune]int, candidate string) bool {
	candidateFreq := getCharFrequency(candidate)
	
	// Если количество уникальных символов разное, это не анаграмма
	if len(subjectFreq) != len(candidateFreq) {
		return false
	}
	
	// Проверяем, что все символы и их частоты совпадают
	for char, count := range subjectFreq {
		if candidateFreq[char] != count {
			return false
		}
	}
	
	return true
}