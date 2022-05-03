package wlbsrv

import (
    "testing"
    "os"
    "fmt"
)

func TestCreateDb(t *testing.T) {
    w := Wlb{}
    w.DbName="test.db"
    err := w.Init()
    if err != nil {
        t.Log("Db file creation failed")
        t.Fail()
    }
    t.Log("Db init success")
    content, err := os.ReadFile("./test/example_2.json")
    var aMsg dbMessage
    rst, _ := WebHookParse(content)
    aMsg.Commit = rst.Commits[0]
    aMsg.Repository = rst.Repository
    w.Db.Create(&aMsg)
    fmt.Printf("%v\n", aMsg.Commit.ID)
    fmt.Printf("%v\n", aMsg.Commit.Timestamp)

    w.UpdateDb(rst)

    t.Log("Finished")
}
