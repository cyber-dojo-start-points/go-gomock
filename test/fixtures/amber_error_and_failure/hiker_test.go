package hiker

import (
    "testing"

    "go.uber.org/mock/gomock"
)

// The first test announces the wrong answer and is reported as a failure.
// The second blows up, and a panic takes the whole test binary with it, so
// it is the last thing in the output whatever comes after it here.
func Test_life_the_universe_and_everything(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnAnswer(42)
    NewHiker(listener).Answer()
}

func Test_the_answer_is_two_digits_long(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnAnswer(2)
    NewHiker(listener).AnswerDigitCount()
}
