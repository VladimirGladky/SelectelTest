package utils

func IsLogMethod(methodName string) bool {
	logMethods := []string{
		"Debug", "Info", "Warn", "Error",
		"DebugContext", "InfoContext", "WarnContext", "ErrorContext",
		"DPanic", "Panic", "Fatal",
		"Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf",
		"Debugw", "Infow", "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw",
		"Print", "Println", "Printf",
	}

	for _, method := range logMethods {
		if methodName == method {
			return true
		}
	}
	return false
}

func IsLatinLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func IsEmoji(r rune) bool {
	return (r >= 0x1F600 && r <= 0x1F64F) ||
		(r >= 0x1F300 && r <= 0x1F5FF) ||
		(r >= 0x1F680 && r <= 0x1F6FF) ||
		(r >= 0x1F700 && r <= 0x1F77F) ||
		(r >= 0x1F780 && r <= 0x1F7FF) ||
		(r >= 0x1F800 && r <= 0x1F8FF) ||
		(r >= 0x1F900 && r <= 0x1F9FF) ||
		(r >= 0x1FA00 && r <= 0x1FA6F) ||
		(r >= 0x1FA70 && r <= 0x1FAFF) ||
		(r >= 0x2600 && r <= 0x26FF) ||
		(r >= 0x2700 && r <= 0x27BF)
}

func IsForbiddenPunctuation(r rune) bool {
	forbiddenChars := "!?…:;"
	for _, char := range forbiddenChars {
		if r == char {
			return true
		}
	}
	return false
}

func ContainsSensitiveData(message string) bool {
	sensitiveKeywords := []string{
		"password", "passwd", "pwd",
		"token", "jwt", "bearer",
		"api_key", "apikey", "api-key",
		"secret", "private_key", "private-key",
		"credit_card", "card_number", "cvv",
		"ssn", "social_security",
		"authorization",
	}

	for _, keyword := range sensitiveKeywords {
		for i := 0; i <= len(message)-len(keyword); i++ {
			match := true
			for j := 0; j < len(keyword); j++ {
				if toLower(rune(message[i+j])) != rune(keyword[j]) {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}

	return false
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}
