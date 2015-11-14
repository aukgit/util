package byteutil

import "sort"

// ByteCount is a struct store count of byte
type ByteCount struct {
	Byte  byte
	Count int
}

// ByteCountList is slice of Bytecount
type ByteCountList []ByteCount

func (b ByteCountList) Len() int           { return len(b) }
func (b ByteCountList) Less(i, j int) bool { return b[i].Count < b[j].Count || b[i].Byte < b[j].Byte }
func (b ByteCountList) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// ReversedByteCountList is Reversed ByteCountList
type ReversedByteCountList struct {
	ByteCountList
}

// Less is different from the Less of ByteCountList
func (b ReversedByteCountList) Less(i, j int) bool {
	return b.ByteCountList[i].Count > b.ByteCountList[j].Count || b.ByteCountList[i].Byte < b.ByteCountList[j].Byte
}

// CountOfByte returns the count of byte for a byte slice
func CountOfByte(s []byte) map[byte]int {
	count := make(map[byte]int)
	for _, b := range s {
		count[b]++
	}
	return count
}

// SortCountOfByte sorts count of byte
func SortCountOfByte(count map[byte]int, reverse bool) ByteCountList {
	countList := make(ByteCountList, len(count))
	i := 0
	for b, c := range count {
		countList[i] = ByteCount{b, c}
		i++
	}
	if reverse {
		sort.Sort(ReversedByteCountList{countList})
	} else {
		sort.Sort(countList)
	}
	return countList
}
