package main

func contains[T comparable](s T, arr []T) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}
