	package main

	import (
		"encoding/json"
		"fmt"
		"net/http"
		"time"
		"web-room/lib/datasource"
		"web-room/lib/user"
	)
	func main() {
		ds, dsErr := datasource.Connect()
		if dsErr != nil {
			fmt.Print("error al conectarse")
		} else {
			// user requests
			http.HandleFunc("/user", func (w http.ResponseWriter, r *http.Request){
				if r.Method == "POST" {
					type UserRequest struct {
						BirthDate string `json:"birthdate"`
						FullName  string `json:"fullname"`
					}
					var userRequest UserRequest
					decodedErr := json.NewDecoder(r.Body).Decode(&userRequest)
					if decodedErr != nil {
						panic(decodedErr)
					}
					up := user.UserSaveStruct {
						FullName: userRequest.FullName,
						BirthDate: userRequest.BirthDate,
					}
					_, saveFail := user.SaveInterface(up, ds)
					if saveFail != nil {
						fmt.Fprintf(w,"Error in %s", saveFail.Error())
					} else {
						fmt.Fprintf(w, "saved ")
					}
				} else if r.Method == "GET" {
					userId := r.URL.Query().Get("userId")
					userFound, foundFail := user.FindByIdInterface(userId, ds)
					if foundFail != nil {
						fmt.Fprintf(w,"Error in %s", foundFail.Error())
					} else {
						json.NewEncoder(w).Encode(userFound)
					}
				}
			})
			// starting http server
			srv := &http.Server{
				Addr:    "127.0.0.1:8000",
				// Good practice: enforce timeouts for servers you create!
				WriteTimeout: 15 * time.Second,
				ReadTimeout:  15 * time.Second,
			}
			panic(srv.ListenAndServe().Error())
		}
	}
