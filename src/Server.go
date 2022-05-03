package wlbsrv

import (
    "net/http"
    "fmt"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type wlb struct {
    DbName string
    Db *gorm.DB
}

type dbMessage struct{
    gorm.Model
    Commit whcommits `gorm:"embedded"`
    Repository whrepository `gorm:"foreignKey:ID"`
}

func (wl *wlb)Init() error{
    var err error
    wl.Db, err = gorm.Open(sqlite.Open(wl.DbName),
                            &gorm.Config{})
    if err != nil {
        return err
    }
    wl.Db.AutoMigrate(&dbMessage{})
    return err
}

func (wl *wlb)OnWebhookPost(w http.ResponseWriter, r *http.Request) {
    var whm webHookMessage
    err := json.NewDecoder(r.Body).Decode(&whm)
    if err != nil {
        fmt.Printf("The input message can't be parsed, Syntax Error ? ")
    }
    go wl.UpdateDb(whm)
    
}

func (wl *wlb)UpdateDb(whm webHookMessage) {
    return
}
