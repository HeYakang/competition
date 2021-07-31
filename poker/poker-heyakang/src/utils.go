package src

import "poker/define"

/*
功能：获取二进制最高位的1 返回这个值
目的：获得大的牌
*/
func getFirstOne(data uint64)(result uint64) {
	for data >0{
		result = data
		data = data & (data - 1)
	}
	return
}

/*
功能：获取二进制最前面n个1 返回这个值
目的：
*/
//func getFirstOneNum(data uint64, deleteOneNum int)(result uint64) {
//	coutOne := countOne(data)
//	result := deleteLastOne(data,int(countOne)-deleteOneNum)
//	return
//}


/***********************用于排序输出**********/
/*
功能：递归删除二进制最后几个1
目的：删除多余的手牌（7选5，9选5等）
 */
func deleteLastOne(data uint64, deleteOneNum int) uint64 {
	if deleteOneNum <= 0 {
		return data
	} else {
		deleteOneNum--
		return deleteLastOne(data&(data-1), deleteOneNum)
	}
}

/*
功能：让一个数据<<13位  次数是moveCount  如leftMoveAndAdd（0000 0100 0000 00000，3） 00000100 实现0000 0100 0000 00000 |0000 0100 0000 00000|0000 0100 0000 00000|...
目的：牌型排版
*/
func leftMoveAndAdd(data uint64, moveCount int) (result uint64) {
	for i := 0; i < moveCount; i++ {
		result |= data << uint(i*13)
	}
	return
}

// 统计二进制中1的个数（最大有效位数为16位）
//variable-precision SWAR算法
func countOne(a uint64) uint64 {
	// 这里用了分治思想：先将相邻两个比特位１的个数相加，再将相邻四各比特位值相加...
	a = ((a & 0xAAAA) >> 1) + (a & 0x5555) // 1010101010101010  0101010101010101
	a = ((a & 0xCCCC) >> 2) + (a & 0x3333) // 1100110011001100  0011001100110011
	a = ((a & 0xF0F0) >> 4) + (a & 0x0F0F) // 1111000011110000  0000111100001111
	a = ((a & 0xFF00) >> 8) + (a & 0x00FF) // 1111111100000000  0000000011111111
	return a
}
//如  1111 1100 1010 0101 ->  1010 1000 0101 0101 -> 0100 0010 0010 0010 ->0000 0110 0000 0100 -> 0000 0000 0000 1010  10个1


//判断一组uint64里是否存在顺子 并返回最大的顺子
func findStraight(data uint64,ghostNum uint64) uint64 {
	var cardFaces uint64

	// 定义模板模板,从最大顺子"AKQJT"开始依次与牌面做匹配,例:
	// cardface	0000011011111    0000011011111    		  0000011011111    0000011011111
	// cardMold 1111100000000 -> 0111110000000 -> ... ->  0000011111000 -> 0000000011111
	// superCard

	cardFaces = define.AKQJT
	for cardFaces>0{
        if countOne(data & cardFaces) >= 5-ghostNum{
        	return cardFaces
		}
		cardFaces = cardFaces >> 1
	}

	cardFaces = define.A2345
	if countOne(data & cardFaces) >= 5-ghostNum{
		return cardFaces
	}
	return 0
}


// 三目表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// Case When Then
func CaseWhen(whenThen ...interface{}) interface{} {
	//循环比较多个条件，成立返回后一个参数
	//step long=2
	for i := 0; i < len(whenThen)-1; i += 2 {
		if whenThen[i].(bool) {
			return whenThen[i+1]
		}
	}
	return nil
}
