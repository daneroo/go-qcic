package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxp-14952710823-14953421296-14957925620-a0023a731b")
	params := slack.PostMessageParameters{
		Username:  "qcicbot",
		IconEmoji: ":eye:", // or :ghost:, :eyes:
	}

	// https://api.slack.com/docs/attachments
	attachment := slack.Attachment{
		Color: "#439FE0", // good, warning, danger, or any hex color code (eg. #439FE0)

		AuthorName:    "Author Name",
		AuthorSubname: "Author Subname",

		Title:     "Title",
		TitleLink: "TitleLink",
		Pretext:   "Attachment pretext",
		Text:      "Attachment Text (with markup).\n*bold* `code` _italic_",
		// ImageURL:  "http://daniel-lauzon.com/img/about-me-trans.png",
		ThumbURL: "https://en.gravatar.com/userimage/2562383/0a6632e59a6821599a51eace02d754ea.jpeg",

		MarkdownIn: []string{"text"}, // ["text", "pretext","fields"]

		// Uncomment the following part to send a field too
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Reported by",
				Value: "Snookr",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Priority",
				Value: "Must Respond",
				Short: true,
			},
		},
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage("#qcic", "Some text", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}
