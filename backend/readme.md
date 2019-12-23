# Go, GORM & Gin CRUD Example

## Install

1. Clone this repository to `$GOPATH/src/github.com/cryptosalamander` directory:

        git clone https://github.com/CryptoSalamander/Go-Gin-Examples.git

2. Install `glide`:

        https://glide.sh/

3. CD to `gorm_crud_example` folder:

        cd $GOPATH/src/github.com/cryptosalamander/gorm_crud_example

4. Install dependencies using `glide`:

        glide install

5. Open `main.go` and modify this variable values:

        dbUser, dbPassword, dbUrl, dbName := "root", "root", "tcp(%your_ip%)", "test"

6. Login to `MySQL` and create the database:

        create database test;

7. Run `main.go`:

        go run main.go
