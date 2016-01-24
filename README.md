# supbot

> Sup... Slack?

## What is Sup?

[Sup](https://github.com/pressly/sup) let's you execute commands on remote
machines based on predefined and simple rules.

## What is Supbot?

Supbot is a [Slack](https://www.slack.com) bot that listens to Sup commands.

You can easily monitor, deploy, bring-up, bring-down, tail logs... on any
environment, across any network directly from your Slack channel.

You can do anything Sup can do, directly from Slack.

![screen shot 2016-01-24 at 3 27 42 pm](https://cloud.githubusercontent.com/assets/385670/12538719/1a5c1f48-c2af-11e5-94d9-0be574897f67.png)

## Deploying to a server

This deploys to the server specified by the `Supfile`.

```
make deploy
```

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

