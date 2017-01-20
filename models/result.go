package models

// Result as a return json template.
type RegisterResult struct {
	Result int64
	Err    string
}

type LoginResult struct {
	Result int64
	Err    string
}
