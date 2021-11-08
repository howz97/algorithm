package util

import (
	"io/ioutil"
	"strings"
	"time"
)

func ReadAllLines(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func AverageDuration(dur []time.Duration) (avg time.Duration) {
	for _, e := range dur {
		avg += e
	}
	avg /= time.Duration(len(dur))
	return
}
