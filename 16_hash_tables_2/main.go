package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://gist.githubusercontent.com/deekayen/4148741/raw/01c6252ccc5b5fb307c1bb899c95989a8a284616/1-1000.txt")
	if err != nil {
		log.Fatalln(err)
	}

	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	sc.Split(bufio.ScanWords)

	buckets := make([]int, 12)

	for sc.Scan() {
		n := HashBucket(sc.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
}

func HashBucket(word string, buckets int) int {
	//letter := int(word[0])
	//bucket := letter % buckets / 2
	//return bucket
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
