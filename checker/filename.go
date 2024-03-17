package checker

import "slices"

func IsNameMatch(fileName string, listNames []string) bool {
	return slices.Contains(listNames, fileName)
}
