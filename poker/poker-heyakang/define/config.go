package define

//存储输入牌数据的切片
var MatchSamplesPaths = []string{
	//"./match/match.json",
	//"./match/match_result.json",
	//"./match/seven_cards_with_ghost.json",
	"./match/seven_cards_with_ghost.result.json",
}


//手牌花色对应 黑桃(s:spades)、红心(h:hearts)、方块(d:diamonds)、草花(:clubs)
var Suits = map[byte]int{
	'c' : 0,
	'd' : 1,
	'h' : 2,
	's' : 3,
}

//用一个uint64存储扑克牌的牌面数字（uint16应该也是可以的）
var Faces = map[byte]uint64{
	                 //AKJQT98765432
	'2' : 1 << 0,	 //0000000000001
	'3' : 1 << 1,	 //0000000000010
	'4' : 1 << 2,	 //0000000000100
	'5' : 1 << 3,	 //0000000001000
	'6' : 1 << 4,	 //0000000010000
	'7' : 1 << 5, 	 //0000000100000
	'8' : 1 << 6, 	 //0000001000000
	'9' : 1 << 7, 	 //0000010000000
	'T' : 1 << 8, 	 //0000100000000
	'J' : 1 << 9, 	 //0001000000000
	'Q' : 1 << 10, 	 //0010000000000
	'K' : 1 << 11, 	 //0100000000000
	'A' : 1 << 12, 	 //1000000000000
}

const (
	StraightFlush = 8 // 皇家同花顺&同花顺
	FourOfAKind   = 7 // 四条
	FullHouse     = 6 // 葫芦
	Flush         = 5 // 同花
	Straight      = 4 // 顺子
	ThreeOfAKind  = 3 // 三条
	TwoPair       = 2 // 两对
	OnePair       = 1 // 一对
	HighCard      = 0 // 散牌
)

const (
	// 特殊值        AKQJT98765432
	A2345 uint64= 4111 // 1000000001111
	AKQJT uint64= 7936 // 1111100000000
	A     uint64= 4096 // 1000000000000
)