package wlbsrv
// wh prefix means Webhook
import (
    "encoding/json"
    "time"
    "gorm.io/gorm"
)

type whcommits struct{
    ID uint `gorm:"PrimaryKey;autoIncrement" json:"Ignore"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"Index"`
    CommitId string `json:"id"`
    Message string
    Url string
    Author whperson `gorm:"foreignKey:ID"`
    Committer whperson `gorm:"foreignKey:ID"`
    timestamp string
}

type whperson struct{
    gorm.Model
    Name string
    Email string
    Username string
}

type whrepository struct{
    ID uint `gorm:"PrimaryKey" json:"Ignore"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"Index"`
    RepoId int `json:"id"`
    Owner whoperator `gorm:"foreignKey:ID"`
    Name string `gorm:"Index"`
    Private bool
    Fork bool
    Website string
    RepoCreatedAt string `json:"created_at"`
    RepoUpdatedAt string `json:"updated_at"`
}

type whoperator struct {
    ID uint `gorm:"PrimaryKey" json:"Ignore"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"Index"`
    UserId int `json:"id"`
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
    Commits []whcommits `gorm:"foreignKey:Name"`
    Repository whrepository `gorm:"foreignKey:ID"`
    Pusher whoperator `gorm:"foreignKey:ID"`
    Sender whoperator `gorm:"foreignKey:ID"`

}

func WebHookParse(text []byte)(webHookMessage, error) {
    var whm webHookMessage
    err := json.Unmarshal(text, &whm)
    return whm, err
}
