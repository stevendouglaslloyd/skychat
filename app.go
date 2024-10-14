package main

import (
    "skychat/storage"
    "skychat/communications"
    "skychat/render"
    "log"
    "net/http"
)

func main() {
    storage.CreateDatabase()
    render.RenderPage()
    communications.CommunicationsRoutes()
    log.Fatal(http.ListenAndServe(":8080", nil)) // Correct placement, no comma before this line
}


