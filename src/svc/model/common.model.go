package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/imamponco/v-gin-boilerplate/src/svc/dto"
	"time"
)

var EmptyObjectJSON = json.RawMessage("{}")

type BaseField struct {
	CreatedAt  time.Time       `db:"createdAt"`
	UpdatedAt  time.Time       `db:"updatedAt"`
	ModifiedBy *Modifier       `db:"modifiedBy"`
	Version    int64           `db:"version"`
	Metadata   json.RawMessage `db:"metadata"`
}

type Modifier struct {
	ID       string `json:"id"`
	Role     string `json:"role"`
	FullName string `json:"fullName"`
}

func (m *Modifier) Scan(src interface{}) error {
	return ScanJSON(src, m)
}

func (m *Modifier) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// ScanJSON is a generic scanner function to parse json from row data
func ScanJSON(src interface{}, target interface{}) error {
	// If source is nil, set target to nil
	if src == nil {
		return nil
	}
	// Assert source to byte
	source, ok := src.([]byte)
	if !ok {
		return errors.New("vsqlx: type assertion to byte failed")
	}
	// Unmarshal to target
	err := json.Unmarshal(source, target)
	if err != nil {
		return err
	}
	return nil
}

func NewBaseField(modifiedBy *Modifier) BaseField {
	if modifiedBy == nil {
		modifiedBy = &Modifier{ID: "", Role: "", FullName: ""}
	}
	// Init timestamp
	t := time.Now()

	return BaseField{
		CreatedAt:  t,
		UpdatedAt:  t,
		ModifiedBy: modifiedBy,
		Version:    1,
		Metadata:   EmptyObjectJSON,
	}
}

func ToBaseFieldDTO(m *BaseField) *dto.BaseField {
	return &dto.BaseField{
		UpdatedAt:  m.UpdatedAt.Unix(),
		CreatedAt:  m.CreatedAt.Unix(),
		ModifiedBy: ToModifierDTO(m.ModifiedBy),
		Version:    m.Version,
	}
}

func ToModifierDTO(model *Modifier) *dto.Modifier {
	return &dto.Modifier{
		ID:       model.ID,
		Role:     model.Role,
		FullName: model.FullName,
	}
}

func ToModifier(subject *dto.Subject) *Modifier {
	return &Modifier{
		ID:       subject.ID,
		Role:     "User",
		FullName: subject.FullName,
	}
}
