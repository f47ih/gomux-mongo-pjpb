package repository

import(
	"github.com/f47ih/gomux-mongo-pjpb/src/modules/laporan/model"
)
type laporanRepository interface {
	Save(*model.Laporan) error
	Update(string, *model.Laporan) error
	Delete(string) error
	findByID(string) (*model.Laporan)
	findAll() (model.Laporan, error)
}