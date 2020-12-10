package godiff


//RunScript takes a generated script from GenScript and returns the result with the deltas included
func RunScript(s []byte) []byte {
	var b0, b1 byte

	result := []byte{}

	for i := 0; i < len(s); i += 2 {
		b0 = s[i]
		b1 = s[i + 1]

		if b0 == '+' {
			result = append(result, b1)
		} else if b0 == ' ' {
			result = append(result, b1)
		}
		//don't need a check for '-'. no action required
	}
	return result
}

//GenScript creates a delta script which stores the original byte string given the steps needed for the SES(shortest edit script)/LCS(longest common subsequence) to record the deltas.
func GenScript(f1, f2 []byte, steps [][]int) []byte {
	script :=[]byte{}

	for i := len(steps) - 1; i >= 0; i-- {
		var opData []byte
		s := NewStep(steps[i])
		if s.R == '+' {
			opData = s.Op(f2[steps[i][1]])
		} else {
			opData = s.Op(f1[steps[i][0]])
		}


		script = append(script, opData...)
	}

	return script
}

type Step struct {
	Step []int
	yInc bool
	xInc bool
	R    rune
}

func NewStep(step []int) Step {
	s := Step{
		Step: step,
		xInc: step[0] < step[2],
		yInc: step[1] < step[3],
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

func (s *Step) Op(data byte) []byte {
	b := []byte(string(s.R))
	return []byte{b[0], data}
}
