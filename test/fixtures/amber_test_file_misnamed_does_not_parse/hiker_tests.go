package hiker

// Named hiker_tests.go rather than hiker_test.go, so it is ordinary source
// and never runs as a test, and it is half written so it does not parse
// either. Vanishing is what it must not do: ordinary source is compiled with
// the package, so the compiler names this file and its line.
func three_digits() int {
    return 100 * (6 * 7
}
