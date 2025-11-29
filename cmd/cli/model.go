package main

type GitHubEvent struct {
    ID        string      `json:"id"`
    Type      string      `json:"type"`
    Actor     Actor       `json:"actor"`
    Repo      Repo        `json:"repo"`
    Payload   Payload     `json:"payload"`
    Public    bool        `json:"public"`
    CreatedAt string      `json:"created_at"`
}

type Actor struct {
    ID          int    `json:"id"`
    Login       string `json:"login"`
    DisplayLogin string `json:"display_login"`
    GravatarID  string `json:"gravatar_id"`
    URL         string `json:"url"`
    AvatarURL   string `json:"avatar_url"`
}

type Repo struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    URL  string `json:"url"`
}

type Payload struct {
    // Para WatchEvent
    Action string `json:"action,omitempty"`

    // Para PushEvent
    RepositoryID int    `json:"repository_id,omitempty"`
    PushID       int64  `json:"push_id,omitempty"`
    Ref          string `json:"ref,omitempty"`
    Head         string `json:"head,omitempty"`
    Before       string `json:"before,omitempty"`
}
