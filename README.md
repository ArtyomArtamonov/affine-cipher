# Affine Cipher

[More on Wikipedia](https://en.wikipedia.org/wiki/Affine_cipher)

Repository contains simple Affine cipher encrypt and decrypt functions

---

Dictionary
```go
const dictionary = "abcdefghijklmnopqrstuvwxyz "
```

And input data
```go
    // Input data
	str := "hello world"
	// Keys
	a, b := 4, 4
```
Could be modified.

---

Make sure key 'a' is co-prime with dictionary length (program will exit with panic if not).
