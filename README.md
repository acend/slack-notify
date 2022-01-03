# Slack-Notify

Slack-Notify is a very simple tool for sending a Slack notification via a
Slack Incoming Webhook.

It is designed to function as a 12-factor app, receiving configuration via
environment variables. To keep things simple, it is rigid in what it allows you
to set. More complex slackbots might want to customize this a little more.

## Slack Incoming Webhooks

This tool uses [Slack Incoming Webhooks](https://api.slack.com/messaging/webhooks)
to send a message to your slack channel.

Before you can use this tool, you need to log into your Slack account and configure
this.

## Usage

Running `slack-notify` in a shell prompt goes like this:

```console
export SLACK_WEBHOOK=https://hooks.slack.com/services/Txxxxxx/Bxxxxxx/xxxxxxxx
SLACK_MESSAGE="hello" slack-notify
```

Running the Docker container goes like this:

```bash
export SLACK_WEBHOOK=https://hooks.slack.com/services/Txxxxxx/Bxxxxxx/xxxxxxxx
docker run -e SLACK_WEBHOOK=$SLACK_WEBHOOK -e SLACK_MESSAGE="hello" ghcr.io/acend/slack-notify:latest
```

## Environment Variables

```bash
# The Slack-assigned webhook
SLACK_WEBHOOK=https://hooks.slack.com/services/Txxxxxx/Bxxxxxx/xxxxxxxx
# The body of the message
SLACK_MESSAGE="Today is a fine day"

```


## Build It

```bash
docker build -t ghcr.io/acend/slack-notify:latest .
```
