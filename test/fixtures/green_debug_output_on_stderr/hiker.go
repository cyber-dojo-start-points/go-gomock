package hiker

import (
    "fmt"
    "os"
)

//go:generate mockgen -source=hiker.go -destination=mock_hiker.go -package=hiker

type Listener interface {
    OnAnswer(answer int)
}

type Hiker struct {
    listener Listener
}

func NewHiker(listener Listener) *Hiker {
    return &Hiker{listener: listener}
}

func (hiker *Hiker) Answer() {
    // The learner is watching when the listener gets told and has not taken
    // this line out yet. go test gives the test binary one stream for both
    // of its own, so this arrives among the results rather than beside them.
    fmt.Fprintln(os.Stderr, "about to announce the answer")
    hiker.listener.OnAnswer(6 * 7)
}
