package main

import (
	"io"
	"net/http"
	"regexp"

	"github.com/gofiber/fiber"
)

func main() {
	fiber := fiber.New()

	fiber.Get("/beef/summary", beefSummaryHandler)

	fiber.Listen(":3200")
}

func beefSummaryHandler(c *fiber.Ctx) {

	var beefWords = map[string]bool{
		"beef": true, "t-bone": true, "ribeye": true, "brisket": true,
		"sirloin": true, "shank": true, "flank": true, "short ribs": true,
	}

	text, err := getTextFromApi("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		c.Status(http.StatusInternalServerError).SendString("Error fetching text from API")
		return
	}

	wordCount := countBeefFromKeyWord(text, beefWords)

	c.JSON(wordCount)
}

func countBeefFromKeyWord(text string, keyWord map[string]bool) map[string]int {

	wordCount := make(map[string]int)

	regex := regexp.MustCompile(`\w[\w-]*`)
	words := regex.FindAllString(text, -1)

	for _, word := range words {
		if keyWord[word] {
			wordCount[word]++
		}
	}

	return wordCount
}

func getTextFromApi(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
