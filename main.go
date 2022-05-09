package main

import (
	github "GORep/Repository/GitHub"
	"GORep/cmd"
	"fmt"
	githook "github.com/go-playground/webhooks/v6/github"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	fmt.Println("Before Init")
	cmd.Init()
	hook, _ := githook.New(githook.Options.Secret(github.MyGithubAccess.TOKEN))
	app := fiber.New()
	app.Post("/webhook", func(c *fiber.Ctx) error {
		fmt.Println("Recieving webhook...")

		httpRequest := new(http.Request)
		// err := fasthttpadaptor.ConvertRequest(c.Context(), httpRequest, true)
		payload, err := hook.Parse(httpRequest, githook.IssueCommentEvent)
		if err != nil {
			fmt.Println("Error parsing", err)
		}

		switch payload.(type) {
		case githook.IssueCommentPayload:
			comment := payload.(githook.IssueCommentPayload)
			commentText := comment.Comment.Body
			userName := comment.Comment.User.Login

			fmt.Printf("user %s posted %s\n", userName, commentText)
		}
		return c.SendStatus(200)
	})
	app.Listen(":3390")
}
