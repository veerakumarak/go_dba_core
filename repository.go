package dba

import "errors"

var (
	ErrRepositoryClientNotAbleToConnect = errors.New("dba: unable to connect to client")
	ErrRepositoryEntityNotFound         = errors.New("dba: no entity found")
	ErrRepositoryInternalError          = errors.New("dba: internal error")
)

type Repository[Entity any] interface {
	Count() (int64, error)
	FindById(entity *Entity, id string) error
	//Find(search Search) ([]Entity, error)
	Save(entity *Entity) error
	SaveAll(entities []Entity) error
	ExistsById(id string) (bool, error)
	DeleteById(id string) error
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
