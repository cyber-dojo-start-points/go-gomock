package hiker

import (
    "testing"

    "go.uber.org/mock/gomock"
)

func Test_life_the_universe_and_everything(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnAnswer(42)
    NewHiker(listener).Answer()
}

// The hiker does announce an answer, so nothing is called that was not
// expected; it is the second announcement that never comes. gomock says so
// only when the controller checks its expectations at the end of the test.
func Test_the_answer_is_announced_twice(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnAnswer(gomock.Any()).Times(2)
    NewHiker(listener).Answer()
}

func Test_the_answer_is_not_the_question(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnAnswer(gomock.Not(gomock.Eq(54)))
    NewHiker(listener).Answer()
}
