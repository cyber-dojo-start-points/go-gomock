package hiker

// A .go file added beside hiker.go belongs to the same package, so hiker.go
// calls this without importing anything. It has no //go:generate line of its
// own because it declares no interface to mock.
func checksum() int {
    return 7
}
