package hiker

import (
    "testing"

    "go.uber.org/mock/gomock"
)

// The learner has taken the one test out and left the helper that set the
// mock up behind. A function only counts as a test when its name begins
// Test, so nothing here runs, and go test prints PASS for a binary that ran
// nothing: the warning about it is the only sign anything is wrong.
func answering_hiker(t *testing.T) *Hiker {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
    listener.EXPECT().OnAnswer(42)
    return NewHiker(listener)
}
