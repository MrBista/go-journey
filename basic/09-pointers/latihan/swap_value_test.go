package latihan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func swapValue(a, b *int) {
	// 1. a b
	temp := *a
	*a = *b
	*b = temp
}

func TestSwapValue(t *testing.T) {

	a, b := 1, 2

	swapValue(&a, &b)

	assert.Equal(t, 2, a)
	assert.Equal(t, 1, b)

}

func safeIncrement(p *int) bool {
	if p == nil {
		return false
	}

	*p++

	return true
}

func TestSafeIncrement(t *testing.T) {
	valP := 0

	safeIncrement(&valP)

	assert.Equal(t, 1, valP)
	safeIncrement(&valP)
	safeIncrement(&valP)
	safeIncrement(&valP)
	safeIncrement(&valP)
	safeIncrement(&valP)
	safeIncrement(&valP)
	safeIncrement(&valP)

	assert.Equal(t, 8, valP)
}

type Student struct {
	Name  string
	Score int
}

func (s *Student) UpdateScore(score int) {
	s.Score = score
}

func UpdateScore(score int, student *Student) {
	if student == nil {
		return
	}

	if score >= 0 && score <= 100 {
		student.Score = score
	}
}

func TestUpdateScore(t *testing.T) {
	person1 := &Student{Name: "Alice", Score: 20}

	fmt.Println(person1.Name)
}
