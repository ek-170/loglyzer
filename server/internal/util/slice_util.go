package util

import (
	"log"
)

func RemoveSlice[T any](arr []T, i int) ([]T) {
    if i < 0 || i >= len(arr) {
		log.Print("index out of range occured.")
        return arr
    }
    return arr[:i+copy(arr[i:], arr[i+1:])]
}