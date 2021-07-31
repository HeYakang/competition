package main

import (
	"poker/define"
	"poker/src"
)


func main() {

	for _, path := range define.MatchSamplesPaths {
		src.GetMatchesFromMatchFile(path).PrintCompareResult()
	}

}