package hiker

import "fmt"

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
    // The learner put a print inside a loop to see what was happening, and
    // it prints far more than the 50K the runner keeps of a stream. go test
    // writes its summary once the tests are over, at the end of that same
    // stream, so the summary is what gets dropped.
    total := 0
    for i := 0; i < 5000; i++ {
        total = total + 1
        fmt.Printf("debug: i is %d, total is %d\n", i, total)
    }
    hiker.listener.OnAnswer(6 * 7)
}
