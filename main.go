package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/rosedblabs/rosedb/v2"
)

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

var db *rosedb.DB // 声明RoseDB实例

func main() {
    // 复制DefaultOptions并修改DirPath
    options := rosedb.DefaultOptions
    options.DirPath = "C:/Users/Olivia/Desktop/RoseDBData" // 更新数据存储路径

    var err error
    db, err = rosedb.Open(options)
    if err != nil {
        log.Fatalf("failed to open RoseDB: %v", err)
    }
    defer db.Close()

    // 启动Go服务
    http.HandleFunc("/put", putHandler)
    http.HandleFunc("/get", getHandler)
    log.Println("Server started at :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    key := []byte(user.Email) // 使用Email作为键
    value, _ := json.Marshal(user)

    if err := db.Put(key, value); err != nil {
        http.Error(w, "Failed to store data", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User data saved successfully"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")
    key := []byte(email)

    value, err := db.Get(key)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(value)
}
