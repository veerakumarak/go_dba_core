package dba

import "errors"

var (
	ErrRepositoryClientNotAbleToConnect = errors.New("dba: unable to connect to client")
	ErrRepositoryEntityNotFound         = errors.New("dba: no entity found")
	ErrRepositoryInternalError          = errors.New("dba: internal error")
)

type Repository[Entity any, Id string | uint64] interface {
	Count() (uint64, error)
	FindById(entity *Entity, id Id) error
	//Find(search Search) ([]Entity, error)
	Save(entity *Entity) error
	SaveAll(entities []Entity) error
	ExistsById(id Id) (bool, error)
	DeleteById(id Id) error
	//Delete(search Search) ([]Entity, error)
}

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
