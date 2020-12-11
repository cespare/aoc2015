package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func init() {
	addSolutions(1, problem1)
}

func problem1(ctx *problemContext) {
	b, err := ioutil.ReadAll(ctx.f)
	if err != nil {
		log.Fatal(err)
	}
	b = bytes.TrimSpace(b)
	ctx.reportLoad()

	ctx.reportPart1(bytes.Count(b, []byte{'('}) - bytes.Count(b, []byte{')'}))
	ctx.reportPart2(elevatorPos(b))
}

func elevatorPos(b []byte) int {
	var level int
	for i, c := range b {
		if c == '(' {
			level++
		} else {
			level--
		}
		if level == -1 {
			return i + 1
		}
	}
	panic("unreached")
}
