package markdown

// implementation to refactor

import (
	"fmt"
	"regexp"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {

	//1 : init some regex
	reStrong := regexp.MustCompile(`(.*)__(.+)__(.*)`)
	reItalic := regexp.MustCompile(`(.*)_(.+)_(.*)`)
	reParagraph := regexp.MustCompile(`^[^*#]`)
	reHeader := regexp.MustCompile(`^(#+) (.+)`)
	reList := regexp.MustCompile(`^\* (.+)`)
	reUl := regexp.MustCompile(`(.*?)(<li>.+</li>)(.*)`)
	//2 : split markdown in lines
	lines := strings.Split(markdown, "\n")

	//3 : iterate over lines
	for i := 0; i < len(lines); i++ {
		//check strong
		if groups := reStrong.FindStringSubmatch(lines[i]); len(groups) == 4 {
			lines[i] = fmt.Sprintf("%s<strong>%s</strong>%s", groups[1], groups[2], groups[3])
		}
		//check italic
		if groups := reItalic.FindStringSubmatch(lines[i]); len(groups) == 4 {
			lines[i] = fmt.Sprintf("%s<em>%s</em>%s", groups[1], groups[2], groups[3])
		}

		//check paragraph
		if reParagraph.MatchString(lines[i]) {
			lines[i] = fmt.Sprintf("<p>%s</p>", lines[i])
		}

		//check headers
		if groups := reHeader.FindStringSubmatch(lines[i]); len(groups) == 3 {
			if len(groups[1]) <= 6 {
				lines[i] = fmt.Sprintf("<h%[1]d>%s</h%[1]d>", len(groups[1]), groups[2])
			} else {
				lines[i] = fmt.Sprintf("<p>%s</p>", lines[i])
			}

		}

		//check list
		if groups := reList.FindStringSubmatch(lines[i]); len(groups) == 2 {
			lines[i] = fmt.Sprintf("<li>%s</li>", groups[1])
		}
	}

	//4 : join lines
	markdown = strings.Join(lines, "")

	//5 : find <ul></ul> on global markdown
	if groups := reUl.FindStringSubmatch(markdown); len(groups) == 4 {
		markdown = fmt.Sprintf("%s<ul>%s</ul>%s", groups[1], groups[2], groups[3])
	}

	return markdown

	/*header := 0
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	pos := 0
	list := 0
	html := ""
	for {
		char := markdown[pos]
		if char == '#' {
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			html += fmt.Sprintf("<h%d>", header)
			pos++
			continue
		}
		if char == '*' {
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
			pos += 2
			continue
		}
		if char == '\n' {
			if list > 0 {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
			pos++
			continue
		}
		html += string(char)
		pos++
		if pos >= len(markdown) {
			break
		}
	}
	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	return "<p>" + html + "</p>"*/

}
