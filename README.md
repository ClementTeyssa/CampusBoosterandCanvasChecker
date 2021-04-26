# CampusBandCanvasChecker

Go app that will send a Telegram notification when Supinfo LMS will be up again ...

Usage
---

- Clone this repository
- Run ```cp app/.env-example app/.env```
- Set Telegram bot token, if you don't have one :
  - Talk to [BotFather](https://t.me/botfather)
  - Follow the instructions and you will get your bot token
- Launch a conversation with your bot
- Set Telegram channel id, if you don't have it :
  - Go to https://web.telegram.org
  - Click on your channel
  - Look at the URL and find the part that looks like c12112121212_17878787878787878
  - Remove the underscore and after c12112121212
  - Remove the prefixed letter 12112121212
  - Prefix with a -100 so your channel id -10012112121212
- Run ```make prod_up```