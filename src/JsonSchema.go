package wlbsrv
// wh prefix means Webhook
import (
    "encoding/json"
    "time"
    "gorm.io/gorm"
)

type whcommits struct{
    CommitId string `json:"id"`
    Message string
    Url string
    AuthorID string
    Author whperson `gorm:"references:Name"`
    CommitterID string
    Committer whperson  `gorm:"references:Name"`
    Timestamp time.Time
}

type whperson struct{
    gorm.Model 
    Name string  `json:"name" gorm:"unique"`
    Email string
    Username string
}

type whrepository struct{
    ID uint `gorm:"PrimaryKey" json:"Ignore"`
    RepoId int `json:"id" gorm:"unique;index"`

    OwnerID int
    Owner whoperator `gorm:"references:UserId"`
    Name string `gorm:"Index"`
    Private bool
    Fork bool
    Website string
    RepoCreatedAt string `json:"created_at"`
    RepoUpdatedAt string `json:"updated_at"`
}

type whoperator struct {
    ID uint `gorm:"PrimaryKey" json:"Ignore"`
    UserId int `json:"id" gorm:"unique"`
    Login string
    FullName string 
    Email  string
    Avatar string  // In fact we don't really care about this
    Username string
}

type webHookMessage struct {
    gorm.Model
    Ref string
    Before string
    After string
    CompareUrl string 
    Commits []whcommits 
    RepositoryID int
    Repository whrepository `gorm:"references:RepoId"`
    PusherID int
    Pusher whoperator 
    SenderID int
    Sender whoperator 

}

func WebHookParse(text []byte)(webHookMessage, error) {
    var whm webHookMessage
    err := json.Unmarshal(text, &whm)
    return whm, err
}
