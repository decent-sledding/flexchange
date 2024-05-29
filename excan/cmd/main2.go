import (
	"os"
	"excan/config"
	"pkg/server"
)




func main() {

}


func startServer() error {
	user := os.Getenv("DB_USER", "unkex");

	connString := "postgres://"

}


func setupDatabase(connStr string) (conn *sql.DB, error) {
	db, err := sql.Open("postgres", connStr);

	if err != nil {
		return nil, err;
	}

	err = db.Ping();

	if err != nil {
		return nil, err;
	}

	return db, nil
}