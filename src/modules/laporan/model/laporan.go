package model

type Laporan struct{
	title			string  `bson:"title"`
	id_laporan 		string  `bson:"id_laporan"`
	deskripsi 		string  `bson:"deskripsi"`
	gambar 			string  `bson:"gambar"`
	kontak_person 	string  `bson:"kontak_person"`
	jml_vote 		int32 	`bson:"jml_vote"`
	email_pelapor 	string  `bson:"email_pelapor"`
	status 			bool	`bson:"status"`
	datetime 		string  `bson:"datetime"`
}

type Laporans []Laporan