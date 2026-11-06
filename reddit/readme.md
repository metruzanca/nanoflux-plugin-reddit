# reddit

reddit feeds are plain RSS, so the generic parser fetches them. This plugin adds
reddit-specific behavior on top.

## Post media

A post's real content (an external destination, an embedded player, or a
multi-image gallery) is marked in reddit's feed HTML in a way a stored item
cannot carry, so it is resolved when the item modal opens.

## Source attribution

A post is shown as "r/cats by u/sam". Both the subreddit and the poster link to
your own subscribed feed for them when you have one, so you stay inside
nanoflux; otherwise they link to reddit.

## Cross-feed dedup

The same post seen through a subreddit feed and through the poster's user feed is
stored once, so reading it in one leaves it read in the other.

A user who reposts the same link to several subreddits under different titles is
also shown once: the plugin keys each link post by its poster and external
destination, and the host collapses the duplicates into a single item that lists
the other subreddits as sources. This applies on the author and collection pages.

## Deleted profiles

If a user profile or subreddit is deleted, banned, or suspended, reddit answers
the feed request with a 404 whose Atom body still says so in its title, for
example `u/someuser: deleted`. The plugin reads that title and pauses the feed
with the reason shown on the feed card.

A paused feed is skipped by the poller, so it stops spending reddit's tight
anonymous request budget. A manual refresh still makes one request, so you can
check whether the profile came back. If it did, re-enable the feed from its
settings; a 404 without the title confirmation stays a normal retryable error.

## Poll interval

Reddit's anonymous `.rss` endpoints are tightly rate-limited and every subreddit
shares the same host, so polling one feed more often spends the budget the others
need. This plugin therefore declares a **one-hour minimum poll interval** for the
feeds it matches. Nanoflux uses the larger of that value and its own configured
minimum, so lowering the global default cannot make a reddit feed poll more often
than once an hour.

## Categories (tags)

Every reddit item carries labels, surfaced as tags in the in-feed tag filter, so
you can narrow a list to a subreddit, an author, or the site a link post points
at.

| Feed | Category labels |
| --- | --- |
| Subreddit (`r/foo`) | the post's author, e.g. `u/someuser` |
| User (`u/someuser`) | the destination subreddit, e.g. `r/foo` |
| Link post | the destination's domain, e.g. `imgur.com` |

So on a subreddit feed the tags identify who posted, and on a user feed they
identify where they posted. A link post adds the site it links to, grouped by
registrable domain: `i.imgur.com` and `www.imgur.com` both become `imgur.com`.
Reddit's own destinations (an image on `i.redd.it`, a gallery page) get no domain
tag, and a self post has no external site. The same labels work for any feed
whose parser provides categories, not just reddit.
