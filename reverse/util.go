package reverse

import "math/rand"

func randStringRunes(n int) string {
	letterRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
