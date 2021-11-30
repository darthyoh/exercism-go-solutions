package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools

type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"

type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	r, ok := cb[rank]
    if !ok {
        return 0
    }
	sum := 0
    for _,v := range r {
        if v {
            sum++
        }
    }
	return sum 
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
        return 0
    }
	sum := 0
	for _,v := range cb {
        if v[file-1] {
            sum+=1
        }
    }
	return sum
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
    sum := 0
	for _,v := range cb {
        sum += len(v)
    }
	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	sum := 0
	for _,v := range cb {
        for _,o := range v {
            if o {
                sum ++
            }
        }    
    }
	return sum
}
