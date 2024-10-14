package storage

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"
    "skychat/users"
)

// CreateDatabase initializes the database and creates necessary tables.
func CreateDatabase() {
    db, err := sql.Open("sqlite3", "storage/source/database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create users table
    createTableUsers := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE,
        password TEXT,
        publickey TEXT UNIQUE,
        privatekey TEXT
    );`
    if _, err := db.Exec(createTableUsers); err != nil {
        log.Fatal(err)
    }

    // Create messages table
    createTableMessages := `CREATE TABLE IF NOT EXISTS messages (
        from_username TEXT,
        from_publickey TEXT,
        to_username TEXT,
        to_publickey TEXT,
        message TEXT,
        time TEXT
    );`
    if _, err := db.Exec(createTableMessages); err != nil {
        log.Fatal("wut: ", err)
    }

    // Create contacts table
    createTableContacts := `CREATE TABLE IF NOT EXISTS contacts (
        username TEXT,
        publickey TEXT
    );`
    if _, err := db.Exec(createTableContacts); err != nil {
        log.Fatal(err)
    }
}

// InsertUserAccount inserts a new user account into the database.
func InsertUserAccount(UserAccountObject *users.UserAccount) {
    db, err := sql.Open("sqlite3", "storage/source/database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare the SQL statement
    Stmt, err := db.Prepare("INSERT INTO users (username, password, publickey, privatekey) VALUES (?, ?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer Stmt.Close()

    // Execute the statement with the account details
    _, err = Stmt.Exec(UserAccountObject.UserName, UserAccountObject.UserPassword, UserAccountObject.PublicKey, UserAccountObject.PrivateKey)
    if err != nil {
        log.Fatalf("Failed to insert user %s: %v", UserAccountObject.UserName, err)
    }
}

// InsertUserAccount inserts a new user account into the database.
func InsertContact(ContactObject *users.Contact) {
    db, err := sql.Open("sqlite3", "storage/source/database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare the SQL statement
    Stmt, err := db.Prepare("INSERT INTO contacts (username, publickey) VALUES (?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer Stmt.Close()

    // Execute the statement with the account details
    _, err = Stmt.Exec(ContactObject.UserName, ContactObject.PublicKey)
    if err != nil {
        log.Fatalf("Failed to insert user %s: %v", ContactObject.UserName, err)
    }
}
// InsertMessageSendTo inserts a new message sent to a user.
func InsertMessageSendTo(MessageObject *users.Message) {
    db, err := sql.Open("sqlite3", "storage/source/database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare the SQL statement
    Stmt, err := db.Prepare("INSERT INTO messages (receivefrom, sendto, message, time) VALUES (?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer Stmt.Close()

    // Execute the statement with the message details
    _, err = Stmt.Exec(MessageObject.MessageSendTo, MessageObject.MessageReceiveFrom, MessageObject.MessageBody, MessageObject.MessageTime)
    if err != nil {
        log.Fatalf("Failed to insert message: %v", err)
    }
}


// InsertMessageReceiveFrom inserts a new message received from a user.
func InsertMessageReceiveFrom(MessageObject *users.Message) {
    db, err := sql.Open("sqlite3", "storage/source/database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare the SQL statement
    Stmt, err := db.Prepare("INSERT INTO messages (receivefrom, sendto, message, time) VALUES (?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer Stmt.Close()

    // Execute the statement with the message details
    _, err = Stmt.Exec(MessageObject.MessageReceiveFrom, MessageObject.MessageSendTo, MessageObject.MessageBody, MessageObject.MessageTime)
    if err != nil {
        log.Fatalf("Failed to insert message: %v", err)
    }
}

// ReadUserAccount retrieves a user account from the database and checks if the password matches.
func ReadUserAccount(UserAccountObject *users.UserAccount) bool {
    // Store the submitted username and password for later comparison
    SubmittedUserName := UserAccountObject.UserName
    SubmittedUserPass := UserAccountObject.UserPassword

    // Open the database connection
    db, err := sql.Open("sqlite3", "storage/source/database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare to query the user
    Query := "SELECT username, password FROM users WHERE username = ?"

    // Execute the query
    err = db.QueryRow(Query, SubmittedUserName).Scan(&UserAccountObject.UserName, &UserAccountObject.UserPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No user found with that username.")
            return false // Return false since the user was not found
        }
        log.Fatal(err)
        return false // Return false in case of an unexpected error
    }

    // Compare the provided password with the retrieved password
    if UserAccountObject.UserPassword == SubmittedUserPass {
        fmt.Printf("Username: %s, Password matches!\n", UserAccountObject.UserName)
        // Print the retrieved user information
        fmt.Printf("Username: %s, Password: %s\n", UserAccountObject.UserName, UserAccountObject.UserPassword)
        return true // Return true if the password matches
    } else {
        fmt.Println("Password does not match.")
        return false // Return false if the password does not match
    }
}



// ReadUserAccount retrieves a user account from the database and checks if the password matches.
func ReadMessages(MessageObject *users.Message) {

    // Open the database connection
    db, err := sql.Open("sqlite3", "storage/source/database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare to query the user
    Query := "SELECT from_username, from_publickey, to_username, to_publickey, message, time  FROM messages WHERE message = ?"

    MessageBody       				string `json:"message"`
    MessageTime        				string `json:"time"`
    MessageReceiveFromUsername 		string `json:"from_username"`
    MessageReceiveFromPublickey     string `json:"from_publickey"`
    MessageSendToUsername		string `json:"to_username"`
    MessageSendToPublickey



    // Execute the query
    err = db.QueryRow(Query, SubmittedUserName).Scan(&MessageObject.MessageBody, &MessageObject.MessageTime, &MessageObject.MessageReceiveFromUsername, &MessageObject.MessageReceiveFromPublickey, &MessageObject.MessageSendToUsername)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No user found with that username.")
            return false // Return false since the user was not found
        }
        log.Fatal(err)
        return false // Return false in case of an unexpected error
    }

    // Compare the provided password with the retrieved password
    if UserAccountObject.UserPassword == SubmittedUserPass {
        fmt.Printf("Username: %s, Password matches!\n", UserAccountObject.UserName)
        // Print the retrieved user information
        fmt.Printf("Username: %s, Password: %s\n", UserAccountObject.UserName, UserAccountObject.UserPassword)
        return true // Return true if the password matches
    } else {
        fmt.Println("Password does not match.")
        return false // Return false if the password does not match
    }
}
