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
    rst, _ := WebHookParse(content)

    w.UpdateDb(rst)


    var persons []whperson
    w.Db.Find(&persons)

    for i, v := range persons {
        fmt.Printf("%d person = %s, %s (%d)\n",i, v.Name, v.Email, v.ID)
    }

    var repos []whrepository
    w.Db.Find(&repos)

    for i, v := range repos {
        fmt.Printf("%d Repo = %s, %d\n", i, v.Name, v.RepoId)
    }

    var msgs []dbMessage
    w.Db.Find(&msgs)
    for i, v := range msgs{
        fmt.Printf("%d Msg = %s, rep=%d, person=%s|%s\n", 
            i, 
            v.Commit.CommitId, 
            v.RepositoryID, 
            v.Commit.AuthorID,
            v.Commit.CommitterID)
    }

    var ops []whoperator
    w.Db.Find(&ops)
    for _, v := range ops {
            fmt.Printf("OP: %d %d: %s\n",
              v.ID, v.UserId, v.Username)
    }
    



    t.Log("Finished")
}
