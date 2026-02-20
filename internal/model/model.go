package model

type CommitType string

const (
	TypeAdded  CommitType = "Added"
	TypeFixed  CommitType = "Fixed"
	TypeChanged CommitType = "Changed"
	TypeDocs   CommitType = "Docs"
	TypeChore  CommitType = "Chore"
	TypeTests  CommitType = "Tests"
)

type Commit struct {
	Hash    string
	Subject string
	Body    string

	Type     CommitType
	Area     string
	Breaking bool
}