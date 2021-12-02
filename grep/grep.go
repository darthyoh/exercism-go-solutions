package grep

import "fmt"
import "regexp"
import "strings"

func getFile(filename string) (file []string) {
	file = make([]string, 0)

	retext := regexp.MustCompile(`\|(.*)*\|`)
	for i, line := range fileContentData {
		if matched, _ := regexp.MatchString(" "+filename, line); matched {
			for j := i + 2; j < len(fileContentData); j++ {
				if rematched, _ := regexp.MatchString("---", fileContentData[j]); rematched {
					return
				}
				file = append(file, strings.TrimRight(retext.FindStringSubmatch(fileContentData[j])[1], " "))
			}
		}
	}
	return
}

//Search is a grep-like function to perform search of a pattern in files with different flags
func Search(pattern string, flags []string, filenames []string) []string {

	files := make(map[string][]string)

	for _, filename := range filenames {
		files[filename] = getFile(filename)
	}

	grep := make([]string, 0)

	var isInverted, isInsensitive, isFileName, isEntireLine, isLineFlag bool

	for _, f := range flags {
		switch f {
		case "-v":
			isInverted = true
		case "-x":
			isEntireLine = true
		case "-l":
			isFileName = true
		case "-i":
			isInsensitive = true
		case "-n":
			isLineFlag = true
		}
	}

	re := regexp.MustCompile(pattern)
	if isInsensitive {
		re = regexp.MustCompile("(?i)" + pattern)
	}

	matchLine := func(line string) bool {
		match := false

		if isEntireLine {
			if isInsensitive {
				if strings.ToLower(line) == strings.ToLower(pattern) {
					match = true
				}
			} else {
				if line == pattern {
					match = true
				}
			}
		} else {
			if re.MatchString(line) {
				match = true
			}
		}

		if isInverted {
			return !match
		}

		return match

	}

	for _, fileName := range filenames {
		for lineIndice, line := range files[fileName] {
			if matchLine(line) {
				if isFileName {
					grep = append(grep, fileName)
					break
				}
				toAdd := ""
				if len(files) > 1 {
					toAdd = fmt.Sprintf("%v:", fileName)
				}
				if isLineFlag {
					toAdd += fmt.Sprintf("%v:", lineIndice+1)
				}
				toAdd += line
				grep = append(grep, toAdd)
			}
		}
	}

	return grep
}
