package MovingOn

func MapStringInterfaceContainKey(domain map[string]interface{}, target string) bool {
	for key, _ := range domain {
		if key == target {
			return true
		}
	}
	return false
}

func StringSliceContainsString(domain []string, target string) bool {
	for _, value := range domain {
		if value == target {
			return true
		}
	}
	return false
}
