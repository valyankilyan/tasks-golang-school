package chess

import (
	"errors"
)

var ErrWrongCordsFormat = errors.New("wrong cords format")
var ErrSameCell = errors.New("knights one the same cell")

func OutOfBoundariesHandler(knight string) bool {
	return !(knight[0] >= 'a' && knight[0] <= 'h') ||
		!(knight[1] >= '1' && knight[1] <= '8')
}

func MoveHandler(white, black string) (func(dx, dy int) (ans bool), error) {
	if len(white) != 2 || len(black) != 2 {
		return nil, ErrWrongCordsFormat
	}
	if OutOfBoundariesHandler(white) || OutOfBoundariesHandler(black) {
		return nil, ErrWrongCordsFormat
	}
	if white[0] == black[0] && white[1] == black[1] {
		return nil, ErrSameCell
	}
	w := []int{int(white[0]), int(white[1])}
	b := []int{int(black[0]), int(black[1])}

	return func(dx, dy int) (ans bool) {
		ans = (w[0]+dx == b[0]) && (w[1]+dy == b[1])
		return
	}, nil
}

func CanKnightAttack(white, black string) (bool, error) {
	sign_mask := [][]int{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	move_mask := [][]int{
		{1, 2},
		{2, 1},
	}

	IsHit, err := MoveHandler(white, black)

	if err != nil {
		return false, err
	}

	ans := false
	for _, s := range sign_mask {
		for _, m := range move_mask {
			ans = ans || IsHit(m[0]*s[0], m[1]*s[1])
		}
	}

	return ans, nil
}
