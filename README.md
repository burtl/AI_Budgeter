




* Install and use `golang-migrate` for SQL migrations.
    ```bash 
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```
  
    SQL migrations files will be in here:
    ```sh 
    mkdir -p migrations
    ```  

    Example SQL migrations file creation:
    ```sh 
    migrate create -ext sql -dir migrations -seq create_users_table
    ```
  

* Schema design for scalability.
  * `Balance` is updated frequently (whenever a transaction is made).
    This could lead to high write contention on the users table, especially in a high-traffic app.
    * Compute dynamically instead of having a `balance` column in the `Users` table.
      * High scalability (immutable transactions)
      * Tradeoff: Slower for high-volume queries.
    * Alternative: ledger table.
      * Faster balance retrieval.
      * Tradeoff: Can have race-conditions.



Note: Make sure you have docker desktop installed on your localhost. 
```sh 
docker compose --profile database up
docker compose --profile demo --profile database up

# To teardown everything:
docker compose down -v
```



Add Gin to the project.
```sh 
go get github.com/gin-gonic/gin
```

Add bcrypt to the project.
```sh 
go get golang.org/x/crypto/bcrypt
```

