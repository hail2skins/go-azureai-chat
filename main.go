package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hail2skins/go-azureai-chat/openaiclient"
	"github.com/hail2skins/go-azureai-chat/setup"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Load environment variables
	setup.LoadEnv()

	// Create an Azure OpenAI client using the environment variables
	client := openaiclient.CreateAzureOpenAIClient(os.Getenv("AI_API_KEY"), os.Getenv("AI_ENDPOINT"), os.Getenv("AI_MODEL"))

	// Set up the Gin web framework with default settings
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Define a route for the root path ("/") that serves the index.html template
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Define a route for the "/api/chat" path that handles chat completions
	r.POST("/api/chat", func(c *gin.Context) {
		prompt := c.PostForm("prompt")
		response, err := openaiclient.ChatCompletion(client, prompt, openai.GPT3Dot5Turbo0301)

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, gin.H{"response": response})
		}
	})

	// Run the Gin server on port 8080
	r.Run(":8080")
}
