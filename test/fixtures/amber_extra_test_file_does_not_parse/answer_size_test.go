package hiker

import (
    "testing"

    "go.uber.org/mock/gomock"
)

func Test_the_answer_is_two_digits_long(t *testing.T) {
    controller := gomock.NewController(t)
    listener := NewMockListener(controller)
