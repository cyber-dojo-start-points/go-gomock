package counting

// A Go package is one directory, so this directory is a package of its own
// and needs its own package clause and its own //go:generate line. Both
// commands in cyber-dojo.sh end in ./..., which walks every directory below
// the sandbox, so the mock below is written and the test beside it is run.
//go:generate mockgen -source=counter.go -destination=mock_counter.go -package=counting

type Tally interface {
    OnCount(count int)
}

type Counter struct {
    tally Tally
}

func NewCounter(tally Tally) *Counter {
    return &Counter{tally: tally}
}

func (counter *Counter) CountDigitsOf(n int) {
    digits := 0
    for n > 0 {
        digits++
        n /= 10
    }
    counter.tally.OnCount(digits)
}
