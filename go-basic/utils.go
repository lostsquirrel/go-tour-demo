package go_basic

import "log"

func printIntSlices(a []int) {
	for i, v := range a {
		log.Printf("%d: %d", i, v)
	}
}

func printSlicesInfo(s []interface{}) {
	log.Printf("cap: %d, len: %d", cap(s), len(s))
}
func printIntSlicesInfo(s []int) {
	log.Printf("cap: %d, len: %d", cap(s), len(s))
}
func printStringSlicesInfo(s []string) {
	log.Printf("cap: %d, len: %d", cap(s), len(s))
}

func printStringSlices(s []string) {
	for i, v := range s {
		log.Printf("%d: %s", i, v)
	}
}
