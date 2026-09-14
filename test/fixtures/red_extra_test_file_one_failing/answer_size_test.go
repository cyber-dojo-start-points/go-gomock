package hiker

import (
    "testing"

    "go.uber.org/mock/gomock"
)

// The file this expectation was set in is named in the failure, which is
// what shows a second test file is picked up.
func Test_the_answer_is_three_digits_long(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnAnswer(gomock.Eq(420))
    NewHiker(listener).Answer()
}
