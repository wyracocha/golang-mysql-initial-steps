package user

import (
	"database/sql"
	"time"
)
 /*
 	Interface for web
 */

func SaveInterface (userInput UserSaveStruct, db *sql.DB) (sql.Result, error){
	bdLayout := "2006-01-02"
	birthdate, bdErr := time.Parse(bdLayout, userInput.BirthDate)
	if bdErr != nil {
		panic(bdErr.Error())
	}
	// building user struct
	user := UserStruct{
		Id: "",
		FullName: userInput.FullName,
		BirthDate: birthdate,
		Status: false,
		CreatedAt: time.Now(),
	}
	saved, savedErr := SaveDriver(user, db)
	if savedErr != nil {
		panic(savedErr.Error())
	} else {
		return saved, nil
	}
}


func FindByIdInterface (userId string, db *sql.DB) (UserStruct, error){
	found, foundErr := FindByIdDriver(userId, db)
	user := UserStruct{}
	if foundErr != nil {
		panic(foundErr.Error())
	} else {
		for found.Next() {
			var Id, FullName string
			var CreatedAt, BirthDate time.Time
			var Status bool
			mapError := found.Scan(&Id, &FullName, &CreatedAt, &BirthDate, &Status)
			if mapError != nil {
				panic(mapError.Error())
			}
			user.Id = Id
			user.FullName = FullName
			user.CreatedAt = CreatedAt
			user.BirthDate = BirthDate
			user.Status = Status
		}
		return user, nil
	}
}

/*

func ChangeStatusInterface (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, cErr := strconv.ParseBool(vars["userStatus"]) // converting status from str to bool
	if cErr != nil {
		panic(cErr.Error()) // throwing error
	}
	changeStatus := struct {
		Id string
		Status bool
	}{
		Id: vars["userId"],
		Status: status,
	}
}

func ChangeBirthDateInterface (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ChangeBirthDate := struct {
		Id string
		BirthDate string
	}{
		Id: vars["userId"],
		BirthDate: vars["birthdate"],
	}
}

func ChangeFullNameInterface (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ChangeFulNmae := struct {
		Id string
		FullName string
	}{
		Id: vars["userId"],
		FullName: vars["fullName"],
	}
}

*/