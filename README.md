# Anime Skip Backend

This is the backend for the Anime Skip web extension.

## Setup

1. Install [Go](https://golang.org/doc/install#download)
2. Install [Postgres](https://www.postgresql.org/download/)
   - Make note of the username and password
3. Install `make`
4. Clone the repo

    ```bash
    git clone git@github.com:aklinker1/anime-skip-backend.git
    ```

5. Generate a `.env` file and see the available `make` commands

    ```bash
    make init
    ```

6. Fill in the `.env` file
7. Run your first build

    ```bash
    make
    ```

8. Create a database in postgres called `anime_skip`
9. Start the server

    ```bash
    make run
    ```

## TODO

- [x] Timestamp Types
- [x] Episode URLs
- [x] [Email helper](https://medium.com/glottery/sending-emails-with-go-golang-and-gmail-39bc20423cf0)
- [x] Create Account
- [x] Validate Email
- [ ] Delete Account
- [ ] Add validation check to isShowAdmin
- [ ] Add foreign key constraints