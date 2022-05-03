package wlbsrv

import  (
    "testing"
    "fmt"
    "encoding/json"
    "os"
)

func TestParse(t *testing.T) {
    content, err := os.ReadFile("./test/example.json")
    if err != nil{
        t.Log("Should touch a example.json first", err)
        t.Fail()
    }
    var result webHookMessage
    err = json.Unmarshal(content, &result)

    if err != nil {
        t.Log("Parsing Failed ", err)
        t.Fail()
    }
    fmt.Printf("%v\n", result.Sender)
    fmt.Printf("%d\n", len(result.Commits))
    fmt.Printf("%v\n", result.Commits[0].CommitId)

    w, err := WebHookParse(content)
    if err != nil {
        t.Log("Parsing Failed ", err)
        t.Fail()
    }
    fmt.Printf("%v\n", w.Commits[0].CommitId)
    fmt.Printf("%v\n", w.Repository.RepoCreatedAt)
    fmt.Printf("%v\n", w.Repository.RepoUpdatedAt)

}



