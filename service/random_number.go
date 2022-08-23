package service

import "math/rand"

func Generate() int {
    return rand.Intn(100)
}
