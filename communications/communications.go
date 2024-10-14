package communications

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "skychat/users"
    "skychat/storage"
)

// CreateAccountRequest handles account creation requests.
func CreateAccountRequest(WriterInstance http.ResponseWriter, RequestInstance *http.Request) {
    if RequestInstance.Method != http.MethodPost {
        http.Error(WriterInstance, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    keyPair := users.GenKeyPairs(1)[0]
    UserAccountObject := &users.UserAccount{}
    if err := json.NewDecoder(RequestInstance.Body).Decode(&UserAccountObject); err != nil {
        http.Error(WriterInstance, "Bad request", http.StatusBadRequest)
        return
    }
    defer RequestInstance.Body.Close()

    UserAccountObject.PublicKey = keyPair.PK.String() // Convert to string
    UserAccountObject.PrivateKey = keyPair.SK.String() // Convert to string

    log.Printf("Received data: %+v\n", UserAccountObject)
    WriterInstance.WriteHeader(http.StatusOK)
    fmt.Fprintf(WriterInstance, "%+v\n", UserAccountObject)
    storage.InsertUserAccount(UserAccountObject)
}

// LoginAccountRequest handles login requests. read account messages and contacts,
func LoginAccountRequest(WriterInstance http.ResponseWriter, RequestInstance *http.Request) {
    if RequestInstance.Method != http.MethodPost {
        http.Error(WriterInstance, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    UserAccountObject := &users.UserAccount{}
    if err := json.NewDecoder(RequestInstance.Body).Decode(&UserAccountObject); err != nil {
        http.Error(WriterInstance, "Bad request", http.StatusBadRequest)
        return
    }
    defer RequestInstance.Body.Close()

    IsInDatabase := storage.ReadUserAccount(UserAccountObject)
    if IsInDatabase {
        WriterInstance.WriteHeader(http.StatusOK)
        fmt.Fprintf(WriterInstance, "Received: %s, %s\n", UserAccountObject.UserName, UserAccountObject.UserPassword)
        
    } else {
        WriterInstance.WriteHeader(http.StatusOK)
        fmt.Fprintf(WriterInstance, "Received: Password does not match")
    }
}


// MessageSendToRequest handles messages sent to a user.
func MessageSendToRequest(WriterInstance http.ResponseWriter, RequestInstance *http.Request) {
    if RequestInstance.Method != http.MethodPost {
        http.Error(WriterInstance, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    MessageObject := &users.Message{}
    if err := json.NewDecoder(RequestInstance.Body).Decode(&MessageObject); err != nil {
        http.Error(WriterInstance, "Bad request", http.StatusBadRequest)
        return
    }
    defer RequestInstance.Body.Close()

    WriterInstance.WriteHeader(http.StatusOK)
    fmt.Fprintf(WriterInstance, "Received: %s, %s\n", MessageObject.MessageBody, MessageObject.MessageTime, MessageObject.MessageReceiveFrom, MessageObject.MessageSendTo)
    storage.InsertMessageSendTo(MessageObject)
}


// MessageRecieveFromRequest handles messages received from a user.
func MessageReceiveFromRequest(WriterInstance http.ResponseWriter, RequestInstance *http.Request) {
    if RequestInstance.Method != http.MethodPost {
        http.Error(WriterInstance, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    MessageObject := &users.Message{}
    if err := json.NewDecoder(RequestInstance.Body).Decode(&MessageObject); err != nil {
        http.Error(WriterInstance, "Bad request", http.StatusBadRequest)
        return
    }
    defer RequestInstance.Body.Close()

    WriterInstance.WriteHeader(http.StatusOK)
    fmt.Fprintf(WriterInstance, "Received: %s, %s\n", MessageObject.MessageBody, MessageObject.MessageTime, MessageObject.MessageReceiveFrom, MessageObject.MessageSendTo)
    storage.InsertMessageReceiveFrom(MessageObject)
}

// MessageSendToRequest handles messages sent to a user.
func AddContactRequest(WriterInstance http.ResponseWriter, RequestInstance *http.Request) {
    if RequestInstance.Method != http.MethodPost {
        http.Error(WriterInstance, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    ContactObject := &users.Contact{}

    if err := json.NewDecoder(RequestInstance.Body).Decode(&ContactObject); err != nil {
        http.Error(WriterInstance, "Bad request", http.StatusBadRequest)
        return
    }
    defer RequestInstance.Body.Close()

    WriterInstance.WriteHeader(http.StatusOK)
    fmt.Fprintf(WriterInstance, "Received: %s, %s\n", ContactObject.UserName, ContactObject.PublicKey)
    storage.InsertContact(ContactObject)
}

// CommunicationsRoutes sets up the routes for the communication.
func CommunicationsRoutes() {
    http.HandleFunc("/LogonAccount", LoginAccountRequest)
    http.HandleFunc("/CreateAccount", CreateAccountRequest)
    http.HandleFunc("/MessageRecieveFrom", MessageReceiveFromRequest)
    http.HandleFunc("/MessageSendTo", MessageSendToRequest)
    http.HandleFunc("/AddContact", AddContactRequest)
}
