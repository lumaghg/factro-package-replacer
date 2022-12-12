package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Factro Package Replacer v1")
	firstIteration := true
	for {
		if !firstIteration {
			fmt.Println("Programm beenden? (J/N)")
			fmt.Print("-> ")
			userInput, _ := reader.ReadString('\n')
			// clean console input from returns and newlines
			userInput = strings.Replace(userInput, "\r\n", "", -1)
			if userInput == "J" {
				break
			}
		}
		firstIteration = false

		//get testuser api token from file
		//TODO ggf. refactoring: do everything with the admin token and obliterate config file
		curDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			time.Sleep(10 * time.Minute)
			return
		}
		filePath := filepath.Join(curDir, "config", "api_user_token.txt")
		fmt.Println("Using filepath: ", filePath)
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			time.Sleep(10 * time.Minute)
			return
		}
		getJWT := string(fileContent)
		getJWT = strings.Replace(getJWT, "\n", "", -1)
		getJWT = strings.Replace(getJWT, "\r", "", -1)
		fmt.Println("API Token used: ", getJWT)

		fmt.Println("Konsolenausgaben zwecks Debugging aktivieren? (J/N)")
		fmt.Print("-> ")
		debuggingOutputs, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		debuggingOutputs = strings.Replace(debuggingOutputs, "\r\n", "", -1)

		fmt.Println("Mit welchem JWT (Token API Admin) sollen Aufgaben geupdated werden?")
		fmt.Print("-> ")
		postJWT, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		postJWT = strings.Replace(postJWT, "\r\n", "", -1)

		fmt.Println("Wie lautet die ProjectId der Packages, für die Werte ersetzt werden sollen?")
		fmt.Print("-> ")
		projectId, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		projectId = strings.Replace(projectId, "\r\n", "", -1)

		fmt.Println("Welches Feld soll ersetzt werden? (Um den Titel zu ändern, hier title eingeben)")
		fmt.Print("-> ")
		fieldToReplace, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		fieldToReplace = strings.Replace(fieldToReplace, "\r\n", "", -1)

		fmt.Println("Welcher Wert soll ersetzt werden? (Beispiel: Kursnummer = MM.YYYY)")
		fmt.Print("-> ")
		valueToFind, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		valueToFind = strings.Replace(valueToFind, "\r\n", "", -1)

		fmt.Println("Durch welchen Wert soll ersetzt werden? (Beispiel: Kursnummer = MM.YYYY)")
		fmt.Print("-> ")
		valueToInsert, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		valueToInsert = strings.Replace(valueToInsert, "\r\n", "", -1)

		client := &http.Client{}
		//request packages from factro
		req, err := http.NewRequest("GET", "https://cloud.factro.com/api/core/projects/"+projectId+"/packages", nil)
		if err != nil {
			fmt.Println(err)
			time.Sleep(10 * time.Minute)
			return
		}
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", getJWT)
		response, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if debuggingOutputs == "J" {
			fmt.Printf("Packages received from Server: %v", string(responseBody))
		}

		var packages []map[string]interface{}

		err = json.Unmarshal(responseBody, &packages)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//find and replace
		//var is named pack instead of package, because package is reserved
		for _, pack := range packages {
			fieldValue, ok := pack[fieldToReplace].(string)
			if !ok {
				fmt.Println("Das angefragte Feld enthält keinen String")
				continue
			}

			if strings.Contains(fieldValue, valueToFind) {
				newFieldValue := strings.ReplaceAll(fieldValue, valueToFind, valueToInsert)
				pack[fieldToReplace] = newFieldValue
			}
		}
		if debuggingOutputs == "J" {
			fmt.Printf("new packages: %v\n", packages)
		}

		//update packages in factro
		reqBody, err := json.Marshal(packages)
		if err != nil {
			fmt.Println(err)
			continue
		}

		req, err = http.NewRequest("PUT", "https://cloud.factro.com/api/core/projects/"+projectId+"/packages/packages", bytes.NewBuffer(reqBody))
		if err != nil {
			fmt.Println(err)
			continue
		}
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", postJWT)
		req.Header.Add("Content-Type", "application/json")

		response, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer response.Body.Close()
		responseBody, err = io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if debuggingOutputs == "J" {
			fmt.Printf("Server responded updating with: %v", string(responseBody))
		}

		fmt.Println("Packages wurden erfolgreich aktualisiert! Ansicht in Factro aktualisieren und überprüfen.")

	}
}
