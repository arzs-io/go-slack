# Simple Package that trigger Slack webhook notification

## Installation

`go get github.com/arzs-io/go-slack`

Setup `SLACK_WEBHOOK_URL` environment variable with your webhook url.

## Usage Example 

```
Import "github.com/arzs-io/go-slack"

func main() {

    status, err := slack.SendInfo("Info example msg")
    status, err := slack.SendWarning("Warning example msg")
    status, err := slack.SendAlert("Error example msg")
    
}
```

Status "ok" is returned if the message is sent successfully.


