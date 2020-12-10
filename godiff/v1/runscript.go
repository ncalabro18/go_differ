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
