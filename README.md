# market_place

Steps to run :
1. go mod init github.com/hiiamninna/market_place
2. go mod tidy (it will generate go.mod and go.sum, if exist, remove it first)
3. go run main.go (make sure in the right folder)

Notes :
- We used golang-migrate, make sure to install it first 
    (ref : https://dev.to/buildhack/golang-database-migration-with-golang-migrate-and-sqlc-4bck)
- Try to implement unit testing, with reference from other github repo, hopefully can understand and implement
    (ref : https://github.com/Fadli2001/go-unit-test-testify)
