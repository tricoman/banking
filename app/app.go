package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/tricoman/banking/domain"
	"github.com/tricoman/banking/service"
	"log"
	"net/http"
	"os"
	"time"
)

// env vars names
const serverAddressEnvVar = "BANKING_SERVER_ADDRESS"
const serverPortEnvVar = "BANKING_SERVER_PORT"
const dbUserEnvVar = "BANKING_APP_DB_USER"
const dbPasswordEnvVar = "BANKING_APP_DB_PASSWORD"
const dbAddressEnvVar = "BANKING_APP_DB_ADDRESS"
const dbPortEnvVar = "BANKING_APP_DB_PORT"
const dbNameEnvVar = "BANKING_APP_DB_NAME"

// API routes
const customersEndpoint = "/customers"
const getCustomerEndpoint = "/customers/{customer_id:[0-9]+}"
const customerAccountEndpoint = "/customers/{customer_id:[0-9]+}/accounts"
const transactionEndpoint = "/customers/{customer_id:[0-9]+}/accounts/{account_id:[0-9]+}"

func Start() {
	startServer()
}

func startServer() {
	sanityCheck()
	router := mux.NewRouter()
	dbClient := initDBClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	transactionRepositoryDB := domain.NewTransactionRepositoryDB(dbClient)
	customerHandlers := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	accountHandlers := AccountHandlers{service.NewAccountService(accountRepositoryDB)}
	transactionHandlers := TransactionHandlers{service.NewTransactionService(transactionRepositoryDB)}
	router.HandleFunc(customersEndpoint, customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc(getCustomerEndpoint, customerHandlers.getCustomer).Methods(http.MethodGet)
	router.HandleFunc(customerAccountEndpoint, accountHandlers.newAccount).Methods(http.MethodPost)
	router.HandleFunc(transactionEndpoint, transactionHandlers.newTransaction).Methods(http.MethodPost)
	address := os.Getenv(serverAddressEnvVar)
	port := os.Getenv(serverPortEnvVar)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func initDBClient() *sqlx.DB {
	//db data "root:qwerty12345@tcp(localhost:3306)/banking"
	dbUser := os.Getenv("BANKING_APP_DB_USER")
	dbPassword := os.Getenv("BANKING_APP_DB_PASSWORD")
	dbAddress := os.Getenv("BANKING_APP_DB_ADDRESS")
	dbPort := os.Getenv("BANKING_APP_DB_PORT")
	dbName := os.Getenv("BANKING_APP_DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	mysqlClient, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	mysqlClient.SetConnMaxLifetime(time.Minute * 3)
	mysqlClient.SetMaxOpenConns(10)
	mysqlClient.SetMaxIdleConns(10)
	return mysqlClient

}

func sanityCheck() {
	errorMessage := "Environment variable %s not defined"
	if os.Getenv(serverAddressEnvVar) == "" {
		log.Fatal(fmt.Sprintf(errorMessage, serverAddressEnvVar))
	}
	if os.Getenv(serverPortEnvVar) == "" {
		log.Fatal(fmt.Sprintf(errorMessage, serverPortEnvVar))
	}
	if os.Getenv(dbUserEnvVar) == "" {
		log.Fatal(fmt.Sprintf(errorMessage, dbUserEnvVar))
	}
	if os.Getenv(dbPasswordEnvVar) == "" {
		log.Fatal(fmt.Sprintf(errorMessage, dbPasswordEnvVar))
	}
	if os.Getenv(dbAddressEnvVar) == "" {
		log.Fatal(fmt.Sprintf(errorMessage, dbAddressEnvVar))
	}
	if os.Getenv(dbPortEnvVar) == "" {
		log.Fatal(fmt.Sprintf(errorMessage, dbPortEnvVar))
	}
	if os.Getenv(dbNameEnvVar) == "" {
		log.Fatal(fmt.Sprintf(errorMessage, dbNameEnvVar))
	}
}
