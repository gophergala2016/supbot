# supbot
Chatops bot for Sup

```
lib/slack // slack bot things
lib/sup // sup wrapper things
lib/hal // instruction parser things
```

`Sup...Slack?`

v 0.1
- now with 100% more HAL9000
- heroku ready!

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```
or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

Make sure to set your `SLACK_TOKEN` env variable

```sh
$ heroku config:set SLACK_TOKEN=xxx-xxyz-xxzy

Setting config vars and restarting xxxx... done
SLACK_TOKEN: xxx-xxyz-xxzy

$ heroku config
=== xxxx Config Vars
SLACK_TOKEN: xxx-xxyz-xxzy
```

Some pointers:

- you might need to run `heroku ps:scale worker=1` to scale dyno manually
