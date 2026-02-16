package types

type Geo struct {
	Lat string
	Lng string
}

type Address struct {
	Street  string
	Suite   string
	City    string
	Zipcode string
	Geo     Geo
}

type PromiseState int

const (
	Pending PromiseState = iota
	Resolved
	Rejected
)

var stateName = map[PromiseState]string{
	Pending:  "pending",
	Resolved: "resolved",
	Rejected: "rejected",
}

func (ps PromiseState) String() string {
	return stateName[ps]
}


func transition(s PromiseState) PromiseState {
    switch s {
    case Pending:
        return Resolved
    case Resolved:
        return Rejected
	default:
		return Pending
    }

}