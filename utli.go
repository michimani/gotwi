package gotwi

func String(s string) *string {
	return &s
}

func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func Bool(b bool) *bool {
	return &b
}

func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func Int(i int) *int {
	return &i
}

func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
