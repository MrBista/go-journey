package arrayandhasing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// region 1. Contains Duplicate

func hashDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func TestContainsDuplicate(t *testing.T) {
	testDuplicate1 := []int{1, 2, 3, 4, 5, 6, 1}

	val1 := hashDuplicate(testDuplicate1)

	assert.Equal(t, true, val1)

	testDuplicate2 := []int{1, 2, 3, 4, 5, 6}

	val2 := hashDuplicate(testDuplicate2)

	assert.Equal(t, false, val2)
}

// endregion

// region 2. Valid anagram
func isAnagram(s string, t string) bool {
	// anagram adalah 2 kata berbeda ketika dibentuk akan

	// kalau panjang karakter tidak sama maka bukan anagram

	if len(s) != len(t) {
		return false
	}

	// countS, countT := make(map[rune]int), make(map[rune]int)

	return true
}
