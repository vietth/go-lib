package db

import(
	"database/sql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

func DbConnect() *gorm.DB {
	var dbURI string
	loadEnv()
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "local" {
		dbURI = getTCPUri()
	} else {
		dbURI = getUnixUri()
	}
	sqlDB, err := sql.Open("mysql", dbURI)
	if err != nil {
		fmt.Printf( "sql.Open error: %v \n", err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
	   fmt.Printf("gorm.Open: %v \n", err)
	}
	if err != nil {
		fmt.Printf("Connect DB error: %v", err)
	}
	return db
}

func loadEnv() {
    // check inside container or not
	_, err := os.ReadFile("serverless_function_source_code/.env.yaml")
	if err != nil {
		_ = godotenv.Load(".env.yaml")
	} else {
		_ = godotenv.Load("serverless_function_source_code/.env.yaml")
	}

}

func getTCPUri() string {
	var (
		dbUser    = os.Getenv("DB_USER")
		dbPwd     = os.Getenv("DB_PASSWORD")
		dbTCPHost = os.Getenv("DB_HOST")
		dbPort    = os.Getenv("DB_PORT")
		dbName    = os.Getenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	return dbURI
}

func getUnixUri() string {
	var (
		dbUser                 = os.Getenv("DB_USER")
		dbPwd                  = os.Getenv("DB_PASSWORD")
		instanceConnectionName = os.Getenv("DB_INSTANCE_CONNECTION_NAME")
		dbName                 = os.Getenv("DB_NAME")
	)
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}
	dbURI := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)
	return dbURI
}

