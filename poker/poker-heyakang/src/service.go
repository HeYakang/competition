package src

import (
	. "competition/poker/poker-heyakang/define"
)

type Hand struct {
	HandStr  string    // 记录原始手牌字符串
	GhostNum uint64    // 鬼牌数量
	Suits    [4]uint64 // 记录手牌中出现过得所有牌的花色
	Faces    [4]uint64 // 记录手牌中出现过得所有牌的出现的次数（数组下标加1即为出现次数，bit位记录手牌牌面）
}

type MaxHand struct {
	MaxCase uint64  //记录最大牌型（StraightFlush, FourOfAKind, FullHouse...）
	MaxHand uint64  //五张牌的排序 （bit 为牌型顺序，因为重要的牌放高位，所有int值可以直接比较得出胜利方）
	FlushFlag bool  //同花标志位
	FlushSuit int   //如果有同花，记录同花的花色编号
}

func getWinnerCompare(a, b uint64) int {
	return CaseWhen(a == b, 0, a > b, 1, a < b, 2).(int)
}

func getWinner(strA string, strB string) int{
	handA := analysisHandStr(strA)
	handB := analysisHandStr(strB)
	maxHandA := MaxHand{}
	maxHandB := MaxHand{}

	var result int
	if result = util(maxHandA.isStraightFlush(handA),maxHandB.isStraightFlush(handB)); result>=0 {
	}else if result =util(maxHandA.isFourOfAKind(handA),maxHandB.isFourOfAKind(handB)); result>=0 {
	}else if result =util(maxHandA.isFullHouse(handA),maxHandB.isFullHouse(handB)); result>=0 {
	}else if result =util(maxHandA.isFlush(handA),maxHandB.isFlush(handB)); result>=0 {
	}else if result =util(maxHandA.isStraight(handA),maxHandB.isStraight(handB)); result>=0 {
	}else if result =util(maxHandA.isThreeOfAKind(handA),maxHandB.isThreeOfAKind(handB)); result>=0 {
	}else if result =util(handA.GhostNum ==0 && maxHandA.isTwoPair(handA),handB.GhostNum ==0 && maxHandB.isTwoPair(handB)); result>=0 {
	}else if result =util(maxHandA.isOnePair(handA),maxHandB.isOnePair(handB)); result>=0 {
	}else if result =util(maxHandA.isHighCard(handA),maxHandB.isHighCard(handB)); result>=0 {
	}
	if result==0 {
		if maxHandA.MaxCase !=StraightFlush &&maxHandA.MaxCase !=Straight {
			maxHandA.getmaxhand(handA)
			maxHandB.getmaxhand(handB)
		}else {
			scoreA := If(maxHandA.MaxHand == A2345, uint64(0), maxHandA.MaxHand).(uint64)
			scoreB := If(maxHandB.MaxHand == A2345, uint64(0), maxHandB.MaxHand).(uint64)
			return getWinnerCompare(scoreA, scoreB)
		}

		return getWinnerCompare(maxHandA.MaxHand, maxHandB.MaxHand)
	}
	return result
}

func util(a,b bool) int{
	if !a && !b {
		return -1
	}else if a==true && b==false  {
		return 1
	}else if a==false && b==true  {
		return 2
	}else  {
		return 0
	}
}

func analysisHandStr(handStr  string) *Hand{
	hand := Hand{HandStr: handStr}
	var faceValue uint64
	for i :=0; i<len(handStr); i++{
		if handStr[i] =='X'{
			hand.GhostNum++
			i++
			continue
		}

		if i%2 == 0 {
			faceValue = Faces[handStr[i]]
			hand.Faces[3] |= hand.Faces[2] & faceValue
			hand.Faces[2] |= hand.Faces[1] & faceValue
			hand.Faces[1] |= hand.Faces[0] & faceValue
			hand.Faces[0] |=  faceValue
		}else {
			hand.Suits[Suits[handStr[i]]] |= faceValue
		}
	}
	return &hand
}

func (maxHand *MaxHand) isStraightFlush(hand *Hand) bool{
	for i :=0;i < len(hand.Faces);i++{
		if countOne(hand.Suits[i]) >= 5-hand.GhostNum{
			maxHand.FlushFlag = true
			maxHand.FlushSuit = i

			if tempValue := findStraight(hand.Suits[i],hand.GhostNum);tempValue > 0{
				if maxHand.MaxHand == 0 {
					maxHand.MaxHand = tempValue
				} else {
					maxHand.MaxHand = If(tempValue > maxHand.MaxHand && tempValue != A2345, tempValue, maxHand.MaxHand).(uint64)
				}
				maxHand.MaxCase = StraightFlush
			}
		}
	}
	return maxHand.MaxCase == StraightFlush
}


func (maxHand *MaxHand) isFourOfAKind(hand *Hand) bool{
	if hand.Faces[3-hand.GhostNum] > 0{
		maxHand.MaxCase = FourOfAKind
		return true
	}
	return false
}

func (maxHand *MaxHand) isFullHouse(hand *Hand) bool{
	if hand.Faces[2-hand.GhostNum] > 0 && countOne(hand.Faces[1]) >= 2  {
		maxHand.MaxCase = FullHouse
		return true
	}
	return false
}

func (maxHand *MaxHand) isFlush(hand *Hand) bool{
	if maxHand.FlushFlag{
		maxHand.MaxCase = Flush
		return true
	}
	return false
}

func (maxHand *MaxHand) isStraight(hand *Hand) bool{
    if maxHand.MaxHand = findStraight(hand.Faces[0],hand.GhostNum);maxHand.MaxHand!=0{
		maxHand.MaxCase = Straight
		return true
	}
	return false
}

func (maxHand *MaxHand) isThreeOfAKind(hand *Hand) bool{

	if hand.Faces[2-hand.GhostNum] > 0{
		maxHand.MaxCase = ThreeOfAKind

		return true
	}
	return false
}

func (maxHand *MaxHand) isTwoPair(hand *Hand) bool{
	if countOne(hand.Faces[1]) >= 2{
		maxHand.MaxCase = TwoPair
		return true
	}
	return false
}

func (maxHand *MaxHand) isOnePair(hand *Hand) bool{
	if hand.Faces[1-hand.GhostNum] > 0{
		maxHand.MaxCase = OnePair
		return true
	}
	return false
}


func (maxHand *MaxHand) isHighCard(hand *Hand) bool{
	maxHand.MaxCase = HighCard
	return true
}

func (maxHand *MaxHand)getmaxhand(hand *Hand)  {
	switch maxHand.MaxCase {
	case HighCard:
		maxHand.MaxHand = deleteLastOne(hand.Faces[0], int(countOne(hand.Faces[0])-5))
	case OnePair:
		firstOne :=getFirstOne(hand.Faces[1-hand.GhostNum])
		tempValue := hand.Faces[0]^firstOne
		maxHand.MaxHand = leftMoveAndAdd(firstOne,2) | deleteLastOne(tempValue,int(countOne(tempValue)-3))
	case TwoPair:
		tempValue := deleteLastOne(hand.Faces[1],int(countOne(hand.Faces[1])-2))
		maxHand.MaxHand = leftMoveAndAdd(tempValue,2) |  getFirstOne(hand.Faces[0]^tempValue)
	case ThreeOfAKind:
		firstOne := getFirstOne(hand.Faces[2-hand.GhostNum])
		tempValue := hand.Faces[0]^firstOne
		maxHand.MaxHand = leftMoveAndAdd(firstOne,3) | deleteLastOne(tempValue,int(countOne(tempValue)-2))
	case Flush:
		tempValue := (hand.Suits[maxHand.FlushSuit] & AKQJT)^AKQJT
		tempValue = deleteLastOne(tempValue,int(countOne(tempValue)-hand.GhostNum))
		tempValue = hand.Suits[maxHand.FlushSuit] | tempValue
		maxHand.MaxHand = deleteLastOne(tempValue,int(countOne(tempValue)-5))
	case FullHouse:
		if hand.GhostNum == 0{
			firstOne := getFirstOne(hand.Faces[2])
			maxHand.MaxHand = leftMoveAndAdd(firstOne,3) | leftMoveAndAdd(getFirstOne(firstOne^hand.Faces[1]),2)
		}else if hand.GhostNum == 1{
			firstOne := getFirstOne(hand.Faces[1])
			maxHand.MaxHand = leftMoveAndAdd(firstOne,3) | leftMoveAndAdd(getFirstOne(firstOne^hand.Faces[1]),2)
		}
	case FourOfAKind:
		if hand.GhostNum == 0{
			firstOne := getFirstOne(hand.Faces[3])
			maxHand.MaxHand = leftMoveAndAdd(firstOne,4) | getFirstOne(firstOne^hand.Faces[0])
		}else {
			firstOne := getFirstOne(hand.Faces[3-hand.GhostNum])
			maxHand.MaxHand = leftMoveAndAdd(firstOne,4) |
				If(firstOne&hand.Faces[3-hand.GhostNum+1] > 0,A,getFirstOne(firstOne^hand.Faces[0])).(uint64)
		}
	}
}


func Compare(strA string, strB string) int {
	playerA := analysisHandStr(strA).getMaxHand()
	playerB := analysisHandStr(strB).getMaxHand()

	if winner := getWinnerCompare(playerA.MaxCase, playerB.MaxCase); winner != 0 {
		return winner
	}

	scoreA := If(playerA.MaxHand == A2345, uint64(0), playerA.MaxHand).(uint64)
	scoreB := If(playerB.MaxHand == A2345, uint64(0), playerB.MaxHand).(uint64)
	return getWinnerCompare(scoreA, scoreB)
}

func (hand *Hand) getMaxHand() *MaxHand{
	maxHand := MaxHand{}

	if maxHand.isStraightFlush(hand) {
	} else if maxHand.isFourOfAKind(hand) {
	} else if maxHand.isFullHouse(hand) {
	} else if maxHand.isFlush(hand) {
	} else if maxHand.isStraight(hand) {
	} else if maxHand.isThreeOfAKind(hand) {
	} else if hand.GhostNum ==0 && maxHand.isTwoPair(hand) {
	} else if maxHand.isOnePair(hand) {
	} else if maxHand.isHighCard(hand) {
	}
	return  &maxHand
}
