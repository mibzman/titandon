# Titandon
A Mastodon Client for Gemini.

# Running
1. go get github.com/mibzman/titandon
1. build
1. enter your configs in `.env`
    - use [the RegisterApp function in go-mastodon](https://github.com/mattn/go-mastodon#application) if you don't have a CLIENTID and CLIENTSECRET.
1. start
1. browse to `gemini://localhost/timeline` (or `http://localhost/timeline` if you're feeling basic)

# Features:
- Toot
- Home Timeline
- Pagination
- View Toot
    - Comments
    - Parents if it's a thread
    - Comment, Boost, and Fav numbers
- Reply to Toot
- Boost Toot
- Fav Toot

# Missing Features
- Images (profile and post)
- Correct CW implementation
    - currently shows CW/subject but does not hide content
- Indicate if a toot is a reply or boost
- Don't break if the first character of a toot is `#`
- login
    - currently requires username and password in .env file
- verification to work on domains outside of localhost  
    - I don't know how TLS works

# FAQ

## Why?
It was fun!  Also I needed a testbed for [Titan](https://github.com/mibzman/titan)
