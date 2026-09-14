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
    hiker.listener.OnAnswer(6 * 9)
}

func (hiker *Hiker) AnswerDigitCount() {
    digits := []int{4, 2}
    hiker.listener.OnAnswer(digits[0] + digits[5])
}
