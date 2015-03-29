package stringutils

func IsEmpty(s string) bool {
	return len(s)==0
}

func IsNotEmpty(s string) bool {
	return len(s)!=0
}

func IsBlank(s string) bool {
	if len(s)==0 {
		return true
	}
	
	for _, rune := range s {
		if rune != ' ' {
			return false
		}
	}
	
	return true
}

func IsNotBlank(s string) bool {
	if len(s)==0 {
		return false
	}
	
	for _, rune := range s {
		if rune != ' ' {
			return true
		}
	}
	
	return false
}

func DefaultIfEmpty(s string, d string) string {
	if IsEmpty(s) {
		return d
	}
	
	return s
}

func DefaultIfBlank(s string, d string) string {
	if IsBlank(s) {
		return d
	}
	
	return s
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}