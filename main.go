package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter name: ")
		scanner.Scan()
		name := scanner.Text()

		fmt.Print("Enter id: ")
		scanner.Scan()
		id := scanner.Text()

		fmt.Print("Enter image name: ")
		scanner.Scan()
		imageName := scanner.Text()

		htmlContent := fmt.Sprintf(`
<div class="content">
    <a href="page.html?id=%s"><img src="drinks/%s.webp"></a>
	<p>%s</p>
</div>
`, id, imageName, name)

		if _, err := writer.WriteString(strings.TrimSpace(htmlContent) + "\n"); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		writer.Flush()

		fmt.Println("Content added to file.")
	}
}
