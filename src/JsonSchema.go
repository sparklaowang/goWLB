package wlbsrv
// wh prefix means Webhook
type whcommits struct{
    Id string
    Message string
    Url string
    Author whperson
    Committer whperson
    timestamp string
}

type whperson struct{
    Name string
    Email string
    Username string
}

type whrepository struct{
    Id int
    Owner whoperator 
    Name string
    Private bool
    Fork bool
    Website string
    CreatedAt string "created_at"
    UpdatedAt string "updated_at"
}

type whoperator struct {
    Id int
    Login string
    FullName string "full_name"
    Email  string
    Avatar string "avatar_url" // In fact we don't really care about this
    Username string
}

type WebHookMessage struct {
    Ref string
    Before string
    After string
    CompareUrl string "compare_url"
    Commits []whcommits
    Repository whrepository
    Pusher whoperator
    Sender whoperator
}
