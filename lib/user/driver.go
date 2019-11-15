package user

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
)

type UserDriver struct {

}
func FindByIdDriver(userId string, driver *sql.DB) (*sql.Rows, error) {
	found, foundErr := driver.Query(
		Query().findById, userId)
	if foundErr != nil { // validating query
		panic(foundErr.Error())
	}
	return found, nil
}

func SaveDriver (userInput UserStruct, driver *sql.DB) (sql.Result, error) {
	// generating uuid
	uid, uidErr := uuid.NewV4()
	if uidErr != nil { // validating uuid
		panic(uidErr.Error())
	}

	userInput.Id = uid.String()
	// creating prepared statement
	saved, savedErr := driver.Prepare(
		Query().save)

	if savedErr != nil { // validating query
		panic(savedErr.Error())
	}

	// execting query
	savedOk, savedFail := saved.Exec(
		userInput.Id, userInput.FullName, userInput.CreatedAt,
		userInput.BirthDate, userInput.Status)

	if savedFail != nil { // validating query execution
		panic(savedFail.Error())
	}
	return savedOk, nil
}