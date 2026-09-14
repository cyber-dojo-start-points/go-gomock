package hiker

import (
    "testing"

    "go.uber.org/mock/gomock"
)

// cyber-dojo.sh passes -v, so go test names every test it runs. The name
// below appearing in the output is what shows this second file was found
// and its test collected from it.
func Test_the_answer_is_two_digits_long(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnAnswer(gomock.Not(gomock.Eq(0)))
    NewHiker(listener).Answer()
}
