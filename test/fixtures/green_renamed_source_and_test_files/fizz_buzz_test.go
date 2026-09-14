package fizzbuzz

import (
    "testing"

    "go.uber.org/mock/gomock"
)

// Go finds a test by the _test.go suffix on the file, not by the stem, so a
// renamed test file keeps that suffix and nothing else about the name matters.
func Test_multiples_of_three_and_five_are_fizzbuzz(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnFizzBuzz("FizzBuzz")
    NewFizzBuzz(listener).Say(15)
}
