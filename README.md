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

How about using a remote network and not only the local one? Yes:

![screen shot 2016-01-24 at 4 01 04 pm](https://cloud.githubusercontent.com/assets/385670/12538901/cee918b8-c2b3-11e5-9974-b4f8a68fea46.png)

## Checking out the repo

```
cd ~/projects
git checkout https://github.com/gophergala2016/supbot.git
```

## Deploying to a server

You can use `sup` to deploy to a server as defined in the `Supfile`.

```
# This is a portion of the supfile
networks:
  ...
  prod:
    hosts:
      - deploy@162.243.9.244
```

```
make deploy
```

If you want to try it locally, use `make docker` to build the docker image,
then use `make docker-run` to run this server locally.

```
SLACK_TOKEN=yyy make docker-run
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

