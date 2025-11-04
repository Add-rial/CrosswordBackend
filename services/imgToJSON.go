package services

import (
	"CrosswordBackend/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

func ImgToJSON(c *gin.Context) string{
	client, err := genai.NewClient(c, &genai.ClientConfig{
		APIKey:  "AIzaSyDB2iSb0KZJJjQqadHwJygVscJkuESpSEw",
        Backend: genai.BackendGeminiAPI,
	})
	if err != nil{
		log.Fatal(err)
	}

	filePath := "crosswordOfTheDay.png"
	bytes, _ := os.ReadFile(filePath)
	parts := []*genai.Part{
		genai.NewPartFromBytes(bytes, "image/png"),
		genai.NewPartFromText(config.GetPrompt()),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		c,
		"gemini-2.5-flash",
		contents,
		config.GetSchemaConfig(),
	)
	if err != nil{
		log.Fatal(err)
	}
	return result.Text()
}