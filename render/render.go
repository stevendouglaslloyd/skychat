package render

import (
    "io/ioutil"
    "log"
    "net/http"
)

func Page(WriterInstance http.ResponseWriter, RequestInstance *http.Request) {
    content, err := ioutil.ReadFile("render/source/index.html")
    if err != nil {
        log.Printf("Error reading HTML file: %v", err)
        http.Error(WriterInstance, "Could not load the page", http.StatusInternalServerError)
        return
    }
    WriterInstance.Header().Set("Content-Type", "text/html")
    WriterInstance.Write(content)
}

func RenderPage() {
    http.HandleFunc("/", Page)
    FileSource := http.FileServer(http.Dir("render/source"))
    http.Handle("/source/", http.StripPrefix("/source/", FileSource))
}
