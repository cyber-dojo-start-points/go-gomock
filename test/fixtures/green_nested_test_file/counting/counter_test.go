package counting

import (
    "testing"

    "go.uber.org/mock/gomock"
)

// The name below appearing in the output, and the second ok line naming the
// counting package, are what show this nested test really ran.
func Test_the_answer_has_two_digits(t *testing.T) {
    controller := gomock.NewController(t)
    tally := NewMockTally(controller)
    tally.EXPECT().OnCount(2)
    NewCounter(tally).CountDigitsOf(42)
}
