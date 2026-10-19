# nanoflux-plugin-reddit

A nanoflux external plugin for Reddit. Reddit's `.rss` feeds are read by the
generic parser; this plugin adds the Reddit-specific behavior: canonical feed
URLs, derived feeds with no extra request, cross-feed post identity, source
attribution, and view-time media resolution.

## Install

```bash
cd nanoflux/plugins/reddit
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../nanoflux-plugin-reddit .
```

Restart nanoflux after putting the binary in its plugins directory
(`NF_PLUGINS_DIR`, `/plugins` in the container).

## Usage

Add a subreddit or user feed, for example `https://www.reddit.com/r/golang/.rss`
or `https://www.reddit.com/user/<name>/submitted.rss`.
