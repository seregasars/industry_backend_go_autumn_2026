package main

func rotateRunes(s string, shift int) string {
	new_rune := []rune(s)
	length := len(new_rune)
	
	if length == 0{
		return ""
	}

	pos := shift % length
	if pos < 0{
		pos += length
	}

	res := make([]rune, length)

	for i := 0; i < length; i++ {
		res[i] = new_rune[(i + pos) % length]
	}
	return string(res)
}
