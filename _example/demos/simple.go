package main

// profile run:
//
// go build -gcflags '-m -l' simple.go
func main() {
	// stackIt()
	// _ = stackIt2()
	slogTest()
}

//go:noinline
func stackIt() int { _ = "STUB: not implemented"; return 0 }

//go:noinline
func stackIt2() *int { _ = "STUB: not implemented"; return nil }

func slogTest() { _ = "STUB: not implemented"; return }

// slog.WithFields(slog.M{
// 	"omg":    true,
// 	"number": 122,
// }).Infof("slog %s", "message message")
