package documents

import "strings"

type markdownSection struct {
	Title       string
	HeadingPath string
	Content     string
}

type markdownHeading struct {
	Level int
	Title string
}

func parseMarkdownSections(content string) []markdownSection {
	lines := strings.Split(content, "\n")

	var sections []markdownSection

	var headingStack []markdownHeading

	currentTitle := ""
	currentHeadingPath := ""

	var currentContent []string

	inCodeBlock := false

	flush := func() {
		content := strings.TrimSpace(strings.Join(currentContent, "\n"))

		if content == "" {
			currentContent = nil
			return
		}

		sections = append(sections, markdownSection{
			Title:       currentTitle,
			HeadingPath: currentHeadingPath,
			Content:     content,
		})
		currentContent = nil
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if isCodeFence(trimmed) {
			currentContent = append(currentContent, line)
			inCodeBlock = !inCodeBlock
			continue
		}

		if inCodeBlock {
			currentContent = append(currentContent, line)
			continue
		}
		heading, ok := parseHeading(trimmed)
		if !ok {
			currentContent = append(currentContent, line)
			continue
		}
		flush()

		for len(headingStack) > 0 &&
			headingStack[len(headingStack)-1].Level >= heading.Level {
			headingStack = headingStack[:len(headingStack)-1]
		}

		headingStack = append(headingStack, heading)

		currentTitle = heading.Title
		currentHeadingPath = buildHeadingPath(headingStack)
	}
	flush()
	return sections
}

func parseHeading(line string) (markdownHeading, bool) {
	if line == "" || !strings.HasPrefix(line, "#") {
		return markdownHeading{}, false
	}

	level := 0

	for level < len(line) && line[level] == '#' {
		level++
	}

	// Markdown ATX headings are from # to ######.
	if level < 1 || level > 6 {
		return markdownHeading{}, false
	}

	// A heading must have a space after the # characters.
	if len(line) <= level || line[level] != ' ' {
		return markdownHeading{}, false
	}

	title := strings.TrimSpace(line[level:])

	if title == "" {
		return markdownHeading{}, false
	}

	return markdownHeading{
		Level: level,
		Title: title,
	}, true
}

func buildHeadingPath(stack []markdownHeading) string {
	parts := make([]string, 0, len(stack))

	for _, heading := range stack {
		parts = append(parts, heading.Title)
	}

	return strings.Join(parts, " > ")
}

func isCodeFence(line string) bool {
	return strings.HasPrefix(line, "```") ||
		strings.HasPrefix(line, "~~~")
}
