package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
)

// structs for the JSON
type User struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

type Organization struct {
	Organization string `json:"organization"`
	Users        []User `json:"users"`
}

const filename = "file.csv"

func main() {
	http.HandleFunc("/organize-csv", func(w http.ResponseWriter, r *http.Request) {
		//Open the file
		file, err := os.Open(filename)
		if err != nil {
			http.Error(w, "Error opening the file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Read and process the data
		reader := csv.NewReader(file)
		//discard the first line:
		_, err = reader.Read()
		if err != nil {
			http.Error(w, "Error reading the file", http.StatusInternalServerError)
			return
		}
		organizations := make(map[string]*Organization)
		for {
			row, err := reader.Read()
			if err != nil {
				break
			}

			orgName := row[0]
			username := row[1]
			role := row[2]

			if org, ok := organizations[orgName]; ok {
				found := false
				for i, user := range org.Users {
					if user.Username == username {
						org.Users[i].Roles = append(user.Roles, role) //it found the user. We have to break
						found = true
						break
					}
				}

				//if didn't find the user, we have to append it
				if !found {
					org.Users = append(org.Users, User{Username: username, Roles: []string{role}})
				}

			} else {
				//create the key in the map
				organizations[orgName] = &Organization{
					Organization: orgName,
					Users:        []User{{Username: username, Roles: []string{role}}},
				}
			}

		}

		var result []Organization
		for _, org := range organizations {
			result = append(result, *org)
		}

		// Convert to JSON
		jsonData, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	http.ListenAndServe(":8080", nil)
}
