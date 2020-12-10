package godiff


//RunScript takes a generated script from GenScript and returns the result with the deltas included
func RunScript(f1, s []byte) []byte {

	result := []byte{}
	fIndex := 0
	i := 0

	for i < len(s) {
		b0 := s[i]

		if b0 == '+' {
			result = append(result, s[i+1])

			i += 2
		} else if b0 == '_' {
			if fIndex < len(f1) {
				result = append(result, f1[fIndex])
			}
			i++
			fIndex++
		} else if b0 == '-' {
			i++
			fIndex++
		} else {
			i++
		}
	}
	return result
}

//GenScript creates a delta script which stores the original byte string given the steps needed for the SES(shortest edit script)/LCS(longest common subsequence) to record the deltas.
func GenScript(f1, f2 []byte) []byte {
	script := []byte{}
	steps := Backtrack(f1, f2)
	for i := len(steps) - 1; i >= 0; i-- {
		var opBytes []byte
		s := NewStep(steps[i])

		f2Index := steps[i][1]
		if f2Index < len(f2) {
			opBytes = s.Operation(f2[f2Index])
		} else {
			opBytes = s.Operation(byte(0))
		}
		script = append(script, opBytes...)
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
	s.R = s.opRune()
	return s
}

func (s *Step) opRune() rune {
	if s.xInc && s.yInc {
		return '_'
	} else if s.xInc {
		return '-'
	} else if s.yInc {
		return '+'
	}
	return '_'
}

func (s *Step) Operation(data byte) []byte {
	b := []byte(string(s.R))
	if s.R == '+' {
		return []byte{b[0], data}
	}
	return b
}
