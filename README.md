# TaskGetDocument
### Files description: 📋

* StoreFile	       Documents 
* struct.go	       Structure the Document
* aplication.go		 Functions the aplication 
* methods.go       Methods the aplication
* main.go          Api Rest localhost:8000/documents
Implement missing operations
* ("/documents", getDocuments).Methods("GET")
* ("/documents/{ID}", getDocumentsById).Methods("GET")
* ("/CreateDocuments", createDocument).Methods("POST")
* ("/DeleteDocuments/{ID}", deleteDocument).Methods("DELETE")

## How to compile:

-Execute the following command.
* go run main.go aplication.go methods.go struct.go
