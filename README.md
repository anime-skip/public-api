# Anime Skip Backend

This is the backend for the Anime Skip web extension.

## Setup

1. Install [Go](https://golang.org/doc/install#download)
1. Install [Postgres](https://www.postgresql.org/download/)
   - Make note of the username and password
1. Install `make`
1. Clone the repo

    ```bash
    git clone git@github.com:aklinker1/anime-skip-backend.git
    ```

1. Generate a `.env` file and see the available `make` commands

    ```bash
    make init
    ```

1. Fill in the `.env` file
1. Run your first build

    ```bash
    make
    ```

1. Create a database in postgres called `anime_skip`
1. Start the server

    ```bash
    make run
    ```

1. Install Modd to auto-restart

    ```bash
    env GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
    make watch
    ```

## Next Goals

- [ ] Add validation directive
- [ ] Delete Account
- [ ] Add foreign key constraints

## Help

If you can't send an email from a gmail account, sign it with it, then go to this link:

https://accounts.google.com/b/0/DisplayUnlockCaptcha
