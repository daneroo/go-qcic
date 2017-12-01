# go-qcic

*Quis custodiet ipsos custodes*
My own monitoring system for deployed apps.

Just small script to observe that deployed systems are functioning correctly.

- im-weight backup
- gmail backup scripts
- scrobblecast/pocketscrape (matching signatures)

## Running

```
. ENV.sh # .gitignore'd loggly credentials
go run qcicstat/qcic.go
go run slackstat/slack.go
```

## Slack integration

https://github.com/nlopes/slack is causing some import problems for assert with govend

```
. ENV.sh # .gitignore'd  credentials    

# public inhook
curl -X POST --data-urlencode 'payload={"text": "This is my message"}' ${SLACK_INHOOK}        

# From go, with a user token ()
go run slackstat/slack.go 
```
