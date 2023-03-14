package constant

type Privilege struct {
	Name    string `json:"name" example:"Get Detail Data"`
	SubName string `json:"subName" example:"G.D"`
	Active  bool   `json:"active" example:"true"`
}

func CreateNewPrivileges(name string, subName string) Privilege {
	return Privilege{
		Name:    name,
		SubName: subName,
		Active:  true,
	}
}

func (p Privilege) SetActive(active bool) Privilege {
	p.Active = active
	return p
}

var Privileges = map[int64][]Privilege{}
