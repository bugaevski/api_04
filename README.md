Language: Go (1.14)

Intent: REST API methods implementation


Functionality:

Initial dummy data population 
    func populateInitialData()

Get all records. Method GET 
    func returnAllInventoryItems(w http.ResponseWriter, r *http.Request)

Get certain record. Method GET 
    func returnSingleInventoryItem(w http.ResponseWriter, r *http.Request)

Create new record. Method POST  
    func createNewInventoryItem(w http.ResponseWriter, r *http.Request)

Update certain record. Data is passed via request body. Method PUT. 
    func updateInventoryItem(w http.ResponseWriter, r *http.Request)

Replace all records by collection passed via body. Method PUT. 
    func replaceAllInventoryItems(w http.ResponseWriter, r *http.Request)

Delete certain record. Method DELETE. 
    func deleteInventoryItem(w http.ResponseWriter, r *http.Request)
 
Delete all records. Method DELETE. 
    deleteAllInventoryItems(w http.ResponseWriter, r *http.Request)
