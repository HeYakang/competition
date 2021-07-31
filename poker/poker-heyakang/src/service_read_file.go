package src

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"time"
)

type Matches struct {
	Matches []*Match `json:"matches"`
}

type Match struct {
	PlayerA string `json:"alice"`
	PlayerB string `json:"bob"`
	Result int `json:"result"`
}

func GetMatchesFromMatchFile(path string) *Matches {
	var file []byte
	var err error
	if file,err = ioutil.ReadFile(path); err !=nil{
		panic("panic: "+ err.Error() )
	}
	matches := Matches{}
	if err = jsoniter.Unmarshal(file , &matches);err != nil {
		panic("panic: "+ err.Error() )
	}
	return &matches
}

func (matches *Matches) PrintCompareResult() {
	beginTime := time.Now()
	for _, v := range matches.Matches {
		res :=getWinner(v.PlayerA, v.PlayerB)

		if res != v.Result {
			fmt.Printf("%s, %s , %d, %d\n", v.PlayerA, v.PlayerB, res, v.Result)
		}
	}
	fmt.Printf("total：%d line\n", len(matches.Matches))
	finishTime := time.Now()
	fmt.Printf("比较共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}








