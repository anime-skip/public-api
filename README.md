# Anime Skip Backend

This is the backend for the Anime Skip web extension.

## Setup

The backend is written in Go. However, all builds are done in a docker container, so you only need Go installed for editor support.

1. Install tooling:
    - [`docker`](https://docs.docker.com/get-docker/)
    - [`docker-compose`](https://docs.docker.com/compose/install/)
    - `make`
    - [`go ≥ v11`](https://golang.org/doc/install#download)
1. Clone the repo
    ```bash
    git clone git@github.com:anime-skip/backend.git
    ```
1. Generate a `.env` file and get an overview of the different `make` commands
    ```bash
    make init
    ```
1. Spin up postgres and other services the backend depends on via `docker-compose`
    ```bash
    make services
    ```
1. Run the server
    ```bash
    make run
    ```
1. (Optional) Install [Modd](https://github.com/cortesi/modd) to use `make watch` (restart the server on change)
    ```bash
    env GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
    make watch
    ```

#### VS Code

Install the recommended extensions, and install all go tooling by  and type in

1. Open command pallet: `ctrl+shift+P` (`mcd+shift+P` for Mac)
1. Run `Go: Install/Update Tools`
1. Select all tools and press OK
