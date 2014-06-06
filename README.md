# twanity

## Overview
Script to graph and track your twitter stats.

## Installation
Clone the repository and then run:
```
% go get
% go install
```

## Usage
You have to register an application in the [twitter development
portal][twitterdev] and then twanity can be configured through an ini file in
`~/.config/twanity/config.ini`:
```
consumerkey       = foo
consumersecret    = bla
accesstoken       = bar
accesstokensecret = baz
graphite          = false
```

or with command line flags:

```
% twanity -h
  -accesstoken="": app access token
  -accesstokensecret="": app access token secret
  -consumerkey="": the app consumer key
  -consumersecret="": the app consumer secret
  -graphite=false: print stats in Graphite format
  -h=false: show this dialog
```

## Send stats to Graphite
twanity has an option to print stats in Graphite compatible format, which you
can enable with a command line flag or by setting the `graphite` key in the
config file to true. And you can use it like this:

```
twanity -graphite mrtazz | nc -w1 graphite.example.com 2003
```

[twitterdev]: https://dev.twitter.com/apps
