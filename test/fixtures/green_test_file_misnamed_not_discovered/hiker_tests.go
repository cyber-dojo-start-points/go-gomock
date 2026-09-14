package hiker

import (
    "testing"

    "go.uber.org/mock/gomock"
)

// Named hiker_tests.go rather than hiker_test.go, so go takes it for ordinary
// source: it is compiled with the rest of the package and the function below
// is never run as a test. The expectation is one the hiker does not meet, so
// a green says the function really did not run.
func Test_the_answer_is_the_question(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnAnswer(54)
    NewHiker(listener).Answer()
}
