package brokennode

// асимптотика тут О(2^n): очень плохая, но
// другого решения придумать не получилось.

var ans []byte

func isRight(reports, state []bool, cur int) bool {
	if cur <= 0 {
		return true
	}

	return !state[cur-1] ||
		(state[cur-1] && reports[cur-1] == state[cur%len(state)])
}

func fillTheNode(cur int, state []bool) {
	if cur >= len(state) {
		return
	}

	var nw byte
	if state[cur] {
		nw = byte('W')
	} else {
		nw = byte('B')
	}

	if ans[cur] == 0 {
		ans[cur] = nw
	}
	if ans[cur] != byte('?') && ans[cur] != nw {
		ans[cur] = '?'
	}
}

func recFind(reports, state []bool, cur, brokenleft int) bool {
	if cur == len(reports) {
		return isRight(reports, state, cur)
	}
	out := false
	if brokenleft < len(state)-cur {
		if cur < len(state) {
			state[cur] = true
		}
		if isRight(reports, state, cur) &&
			recFind(reports, state, cur+1, brokenleft) {
			fillTheNode(cur, state)
			out = true
		}
	}
	if brokenleft > 0 {
		if cur < len(state) {
			state[cur] = false
		}
		if isRight(reports, state, cur) &&
			recFind(reports, state, cur+1, brokenleft-1) {
			fillTheNode(cur, state)
			out = true
		}
	}
	return out
}

func FindBrokenNodes(brokenNodes int, reports []bool) string {
	ans = make([]byte, len(reports))
	state := make([]bool, len(reports))
	recFind(reports, state, 0, brokenNodes)
	return string(ans)
}
