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
