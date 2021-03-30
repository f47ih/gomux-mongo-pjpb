package repository

import (
	"time"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/f47ih/gomux-mongo-pjpb/src/modules/laporan/model"
)

type laporanRepositoryMongo struct{
	db *mgo.Database
	collection string
}

func NewLaporanRepositoryMongo(db *mgo.Database, collection string) *laporanRepositoryMongo{
	return &laporanRepositoryMongo{
		db : db,
		collection: collection,
	}
}

func (r *laporanRepositoryMongo) Save(laporan *model.Laporan) error{
	
	return nil
}

func (r *laporanRepositoryMongo) Update(id string, laporan *model.Laporan) error{
	return nil
}

func (r *laporanRepositoryMongo) Delete(id string) error{
	return nil
}

func (r *laporanRepositoryMongo) findByID(id string) (*model.Laporan, error){
	return nil,nil
}

func (r *laporanRepositoryMongo) findAll() (model.Laporans, error){
	return nil,nil
}