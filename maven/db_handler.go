package maven

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB // Database connection pool.

func insertArtifactFile(artRowId int, filename string, checksum string) bool {
	/*
	 * We should keep our groupID, artifactID and version in a seperate table instead of writing them multiple times for each file.
	 * And then we can associate the artifacts table and files table by the artifacts row id(artRowId) eg. 12
	 * If any error occurs the return should be false.
	 */
}

func artifactSaved(groupID string, artifactID string, version string) (bool, int) {
	/*
	 * This function should check if there is a row that contains groupID, artifactID and version exists.
	 * If exists it should return true, nil otherwise it should return false, row id(artRowId)
	 */
}

func insertArtifact(groupID string, artifactID string, version string) (bool, int) {
	/*
	 * This function should insert  groupID, artifactID, version into "artifacts" table. After the insert it should return the row id(artRowId)
	 * It should check if the artifact is already saved with artifactSaved(). If saved it should return the row id returned by artifactSaved()
	 * If any error occurs the first return should be false.
	 */
}

func insertFile(groupID string, artifactID string, version string, filename string) {
	/*
	 * This is the main handler this function should insert the artifact if not exists and then insert the file to files table.
	 * Each required function is explained above.
	 * Function should proceed like:
	 * success, artRowId := insertArtifact(...)
	 * fsuccess := insertArtifactFile(artRowId, ...)
	 */
}

func connect() {
	// Open up our database connection.
	pool, err := sql.Open("mysql", "root:******@tcp(172.17.0.2:3306)/locrep")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("hata aldi")
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer pool.Close()

	res, _ := pool.Query("SHOW TABLES")
	var table string

	for res.Next() {
		res.Scan(&table)
		fmt.Println(table)
	}
}
