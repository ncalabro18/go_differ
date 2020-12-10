package godiff

var (
	scriptHeader = []byte{0xaa, 0x0f, 0xca, 0xab}
)

//GenScript creates a delta script which stores the original byte string given the steps needed for the SES(shortest edit script)/LCS(longest common subsequence) to record the deltas.
func GenScript(f1, f2 []byte, steps [][]int) []byte {
	script := append([]byte{}, scriptHeader...)

	for i := len(steps) - 1; i > 0; i-- {
		var s Step
		if s.R == '+' {
			s = NewStep(steps[i], f2[steps[i][1]])
		} else {
			s = NewStep(steps[i], f1[steps[i][0]])
		}

		opData := s.Op()
		script = append(script, opData...)
	}

	return script
}

type Step struct {
	Step []int
	yInc bool
	xInc bool
	R    rune
	Data byte
}

func NewStep(step []int, data byte) Step {
	s := Step{
		Step: step,
		xInc: step[0] < step[2],
		yInc: step[1] < step[3],
		Data: data,
	}
	s.R = s.OpRune()
	return s
}

func (s *Step) OpRune() rune {
	if s.xInc && s.yInc {
		return ' '
	} else if s.xInc {
		return '-'
	} else if s.yInc {
		return '+'
	}
	return ' '
}

func (s *Step) Op() []byte {
	b := []byte(string(s.R))
	return []byte{b[0], s.Data}
}
