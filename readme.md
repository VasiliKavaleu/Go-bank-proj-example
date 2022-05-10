# Go project example 


* API via Gin Framework

* sqlc for connection to DB + migrations

* Auth by Tokens (JWT/PASETO)

#### Usage

- Clone the project

- Install requirements (from go.mod + migrate cli)

    ```
    go mod download
    ```

- Copy and fill out config file

    ```
    cp app.env.example app.env
    ```

For rebuilding swagger docs

    ```
    swag init
    ```
    Or
    ```
    $HOME/go/bin/swag init
    ```

Run server (docs available [on this link](http://0.0.0.0:8000/swagger/index.html#))

    ```
    make server
    ```
