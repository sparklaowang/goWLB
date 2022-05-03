package wlbsrv

import (
    "net/http"
    "fmt"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type Wlb struct {
    DbName string
    Db *gorm.DB
    Port int
}

type dbMessage struct{
    gorm.Model
    Commit whcommits `gorm:"embedded"`
    Repository whrepository `gorm:"foreignKey:ID"`
}

func (wl *Wlb)Init() error{
    var err error
    wl.Db, err = gorm.Open(sqlite.Open(wl.DbName),
                            &gorm.Config{})
    if err != nil {
        return err
    }
    wl.Db.AutoMigrate(&dbMessage{})
    wl.Db.AutoMigrate(&whperson{})
    wl.Db.AutoMigrate(&whoperator{})
    return err
}

func (wl *Wlb)OnWebhookPost(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Recv Web Hook\n")
    var whm webHookMessage
    err := json.NewDecoder(r.Body).Decode(&whm)
    if err != nil {
        fmt.Printf("The input message can't be parsed, Syntax Error ? ")
    }
    go wl.UpdateDb(whm)
}

func (wl *Wlb)UpdateDb(whm webHookMessage) {
    for _, commit:= range whm.Commits {
        wl.Db.Create(&dbMessage{Commit: commit, Repository: whm.Repository})
    }
}
// A very basic start that use net/http as server framework
func (wl *Wlb)Start() {
    http.HandleFunc("/", wl.OnWebhookPost)
    http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", wl.Port), nil)
}
