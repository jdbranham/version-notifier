package slack_notifier

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
	"yuvalpress/version-notifier/internal/release_notes"
)

var (
	// Reset color variables after call
	Reset = "\033[0m"

	// Red color for logs
	Red = "\033[31m"
)

// Notify sends a slack message with the supplied data
func Notify(user, repo, url, oldVer, newVer, updateLevel string) {
	slackClient := slack.New(os.Getenv("SLACK_TOKEN"))

	attachment := slack.Attachment{
		Pretext: "New Version Details:",
		Text:    url,
	}

	notes := release_notes.GetReleaseNotes(url)
	if notes != "" {
		_, _, err := slackClient.PostMessage(
			os.Getenv("SLACK_CHANNEL"),
			slack.MsgOptionText("*New"+updateLevel+"update found for package: "+user+"/"+repo+"*"+"\n"+oldVer+" -> "+newVer, false),
			slack.MsgOptionAttachments(attachment),
			slack.MsgOptionUsername("Version Notifier"),
			slack.MsgOptionBlocks(slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", notes, false, false), nil, nil)),
		)
		if err != nil {
			fmt.Printf(Red+"Faild to post message to slack_notifier with the following error: %s\n"+Reset, err)
			return
		}

	} else {
		_, _, err := slackClient.PostMessage(
			os.Getenv("SLACK_CHANNEL"),
			slack.MsgOptionText("*New"+updateLevel+"update found for package: "+user+"/"+repo+"*"+"\n"+oldVer+" -> "+newVer, false),
			slack.MsgOptionAttachments(attachment),
			slack.MsgOptionUsername("Version Notifier"),
		)
		if err != nil {
			fmt.Printf(Red+"Faild to post message to slack_notifier with the following error: %s\n"+Reset, err)
			return
		}
	}
}