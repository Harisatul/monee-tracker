## About

Basic coding, REST API test using Golang make a design system.


## Built with

- Golang
- Postgre
- Goland


## DB Design
see this link :
(https://drive.google.com/file/d/1bB8UgFiCtVVDKFzWuuJ1j42oUHKZrNF8/view?usp=sharing)


1. Clone this project
    ```
    git clone https://github.com/berrylradianh/basic-coding-kulina.git
    ```
2. Configure environment variables for database connection. see file .env-example
    ```
    APP_PORT="${APP_PORT}"
    DB_CONNECTION="${DB_CONNECTION}"
    DB_HOST="${DB_HOST}"
    DB_PORT="${DB_PORT}"
    DB_NAME="${DB_NAME}"
    DB_USERNAME="${DB_USERNAME}"
    DB_PASSWORD="${DB_PASSWORD}"
    SECRET_KEY="${SECRET_KEY}"
    RAJAONGKIR_KEY="${RAJAONGKIR_KEY}"
    MIDTRANS_SERVER_KEY="${MIDTRANS_SERVER_KEY}"
    BINDERBYTE_KEY="${BINDERBYTE_KEY}"
    ```

4.  Run the application
    ```
    go run main.go
    ```


