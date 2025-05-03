package vault

type Entry struct {
    Label    string `json:"label"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type Vault struct {
    Entries []Entry `json:"entries"`
}
