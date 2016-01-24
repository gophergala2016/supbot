# supbot
Simple slack bot powered deployment tool with Sup

```
lib/slack // slack bot things
lib/sup // sup wrapper things
lib/hal // instruction parser things
```

`Sup...Slack?`

## What is Supbot/Supslack?

`Sup` (https://github.com/pressly/sup) let's you quickly execute remote commands
on remote machines based on predefined and simple rules. As explained by the
project: 

```
Stack Up is a simple deployment tool that performs given set of commands on multiple hosts in parallel.
It reads Supfile, a YAML configuration file, which defines networks (groups of hosts), commands and targets.

```

## What is Supbot?

Supbot is a slack bot that listens to Sup commands. You can easily monitor,
deploy, bring-up, bring-down, tail logs... on any environment, across any
network directly from your slack channel.

You can do anything Sup can do, directly from Slack.


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


### Changelog:


v 0.1
- now with 100% more HAL9000
- heroku ready!
