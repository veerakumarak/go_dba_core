package dba

import "errors"

var (
	ErrRepositoryClientNotAbleToConnect = errors.New("dba: unable to connect to client")
	ErrRepositoryEntityNotFound         = errors.New("dba: no entity found")
	ErrRepositoryInternalError          = errors.New("dba: internal error")
)

type BaseEntity[ID string | uint64] struct {
	Id        ID
	CreatedAt int64
	UpdatedAt int64
}

type Repository[Entity any, Id string | uint64] interface {
	Count() (uint64, error)
	FindById(*Entity, Id) error
	FindOne(map[string]interface{}) ([]*Entity, error)
	Save(*Entity) error
	SaveAll([]*Entity) error
	ExistsById(Id) (bool, error)
	DeleteById(Id) error
	//Delete(search Search) ([]Entity, error)
}

// Find().Sort().Page()
//type Pageable struct {
//}
//
//type SearchOps struct {
//}
//
//type FilterOps struct {
//}
//
//type SortOps struct {
//}
//
//type Search struct {
//}
