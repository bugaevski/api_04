package main

// test. Burgess Group
// S:\MiGo\mibGOPATH\src\mi-code\api\api_04

import (
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type inventoryItem struct {
	Name string `json:"name"`
	Quantity int `json:"quantity"`
}

var inventory []inventoryItem

func populateInitialData() {
	var currItem *inventoryItem

	currItem = new(inventoryItem)
	currItem.Name = "Apples"
	currItem.Quantity = 3
  inventory = append(inventory, *currItem)

	currItem = new(inventoryItem)
	currItem.Name = "Oranges"
	currItem.Quantity = 7
  inventory = append(inventory, *currItem)

	currItem = new(inventoryItem)
	currItem.Name = "Pomegranates"
	currItem.Quantity = 55
  inventory = append(inventory, *currItem)

	//DEBUG
	//fmt.Println(inventory)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test. Burgess Group")
}

func returnAllInventoryItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if len(inventory) > 0 {
			json.NewEncoder(w).Encode(inventory) 
		} else {
			fmt.Fprintf(w, "no records found")
		}
	} else {
		fmt.Fprintf(w, "Expected request method: GET. Actual: " + r.Method)
	}
}

func returnSingleInventoryItem(w http.ResponseWriter, r *http.Request) {
	var isFound bool = false
	vars := mux.Vars(r)
	key := vars["id"] //retrieve {id} from request URL
  var item inventoryItem
	
	if r.Method == http.MethodGet {
		for _, currItem := range inventory {
			if currItem.Name == key {
				item = currItem
				isFound = true
				break
			}
		}
		if isFound == true {
			json.NewEncoder(w).Encode(item) 
		} else {
			fmt.Fprintf(w, "no records found")
		}
	} else {
		fmt.Fprintf(w, "Expected request method: GET. Actual: " + r.Method)
	}
}

func createNewInventoryItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var item inventoryItem
		json.Unmarshal(reqBody, &item)
	  inventory = append(inventory, item)
		//DEBUG
		//fmt.Fprintf(w, "%+v", string(reqBody))
	}
}

func updateInventoryItem(w http.ResponseWriter, r *http.Request) {
	var targetItemIndex int
	vars := mux.Vars(r)
	key := vars["id"]

	// get index of the item
	for index, currItem := range inventory {
		if currItem.Name == key {
      targetItemIndex = index
			break
		}
	}

	// get item contents
	if r.Method == http.MethodPut {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var item inventoryItem
		json.Unmarshal(reqBody, &item)
		inventory[targetItemIndex] = item
	}
}

func replaceAllInventoryItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		reqBody, _ := ioutil.ReadAll(r.Body)
		
		inventory = nil
		json.Unmarshal(reqBody, &inventory)

		//DEBUG
	  //fmt.Println(inventory)
	}
}

func deleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if r.Method == http.MethodDelete {
		for index, currItem := range inventory {
			if currItem.Name == key {
				inventory = append(inventory[:index], inventory[index+1:]...)
				break
			}
		}
	} else {
		fmt.Fprintf(w, "Expected request method: DELETE. Actual: " + r.Method)
	}
	//fmt.Println("After DELETE \n")
	//fmt.Println(inventory)
}

func deleteAllInventoryItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
    inventory = nil
	} else {
		fmt.Fprintf(w, "Expected request method: DELETE. Actual: " + r.Method)
	}
}

func handleRequests() {
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/", homePage)
	
	//GET http://localhost:8081/inventory
	apiRouter.HandleFunc("/inventory", returnAllInventoryItems).Methods("GET")

	//GET http://localhost:8081/inventory/Apples
	apiRouter.HandleFunc("/inventory/{id}", returnSingleInventoryItem).Methods("GET")

	//POST http://localhost:8081/inventory + body
	apiRouter.HandleFunc("/inventory", createNewInventoryItem).Methods("POST") 

	//DELETE http://localhost:8081/inventory/Oranges
	apiRouter.HandleFunc("/inventory/{id}", deleteInventoryItem).Methods("DELETE")

	//DELETE http://localhost:8081/inventory - this is bad design per task description!
	apiRouter.HandleFunc("/inventory", deleteAllInventoryItems).Methods("DELETE")

	//PUT http://localhost:8081/inventory/Apples
	apiRouter.HandleFunc("/inventory/{id}", updateInventoryItem).Methods("PUT")

	//PUT http://localhost:8081/inventory
	apiRouter.HandleFunc("/inventory", replaceAllInventoryItems).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", apiRouter))
}

func main() {
	populateInitialData()
	handleRequests()
}
