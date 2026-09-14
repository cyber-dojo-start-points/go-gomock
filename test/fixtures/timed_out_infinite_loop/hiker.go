package hiker

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
    n := 0
    // The learner meant to count up to 42 and never moves n. go test has a
    // timeout of its own, ten minutes by default, which is far longer than
    // the manifest's max_seconds, so the runner is what stops this.
    for n != 42 {
    }
    hiker.listener.OnAnswer(n)
}
