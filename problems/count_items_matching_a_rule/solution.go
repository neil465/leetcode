func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	count := 0
	num := -1
	if ruleKey == "type" {
		num = 0
	}
	if ruleKey == "color" {
		num = 1
	}
	if ruleKey == "name" {
		num = 2
	}
	for _, item := range items {
		if item[num] == ruleValue {
			count++
		}
	}
	return count
}
