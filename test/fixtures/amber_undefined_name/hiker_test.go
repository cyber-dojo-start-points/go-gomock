package hiker

import (
    "testing"

    "go.uber.org/mock/gomock"
)

// mockgen names the mock after the interface, so the mock for Listener is
// NewMockListener. The name reached for below is not the one written.
func Test_life_the_universe_and_everything(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockAnswerListener(controller)
    listener.EXPECT().OnAnswer(42)
    NewHiker(listener).Answer()
}
