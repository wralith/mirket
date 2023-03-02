package random

import "crypto/rand"

// From: https://henvic.dev/posts/uuid/
// ID generates a random 11 character base-58 ID.
func ID() string {
	const (
		alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz" // base58
		size     = 11
	)

	var id = make([]byte, size)
	if _, err := rand.Read(id); err != nil {
		panic(err)
	}
	for i, p := range id {
		id[i] = alphabet[int(p)%len(alphabet)] // discard everything but the least significant bits
	}
	return string(id)
}
