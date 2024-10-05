package main

type Comp struct {
	DisplayName string
	Instance    string
}

type category string

const (
	Open      category = "Open"
	Mixed     category = "Mixed"
	WomenPlus category = "WomenPlus"
	Other     category = "Other"
)

type parentOrg string

type Team struct {
	DisplayName string
	ParentOrg   parentOrg
	Category    category
}

type Ref struct {
	DisplayName string
	parentOrg   parentOrg
}

type CompDay struct {
	Comp   Comp
	Day    string
	Courts int
	Teams  []Team
	Refs   []Ref
}

type Match struct {
	CompDay CompDay
	Teams   []Team
	CourtID int
	Ref     Ref
}
