package dec23

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

// New instance of the Door for December 23
func New(input string) doors.Solver {
	return dec23{input}
}

type dec23 struct {
	input string
}

// Solve the day's problem
func (d dec23) Solve() (string, error) {
	var count int
	err := fileread.ReadAndApply(d.input, func(s string) error {
		return nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", count), nil
}
