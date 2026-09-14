package fizzbuzz

// Both files here are renamed, and so is the package. The //go:generate line
// names the file it reads and the file it writes, so renaming the source
// means editing this line too, and -package must name the package as well.
//go:generate mockgen -source=fizz_buzz.go -destination=mock_fizz_buzz.go -package=fizzbuzz

type Listener interface {
    OnFizzBuzz(word string)
}

type FizzBuzz struct {
    listener Listener
}

func NewFizzBuzz(listener Listener) *FizzBuzz {
    return &FizzBuzz{listener: listener}
}

func (fizzBuzz *FizzBuzz) Say(n int) {
    switch {
    case n%15 == 0:
        fizzBuzz.listener.OnFizzBuzz("FizzBuzz")
    case n%3 == 0:
        fizzBuzz.listener.OnFizzBuzz("Fizz")
    case n%5 == 0:
        fizzBuzz.listener.OnFizzBuzz("Buzz")
    }
}
