package main

import "fmt"


type Mahasiswa struct {
	Nama string
	NilaiTugas int
	NilaiUAS int
}

func main (){
	kelas:=[] Mahasiswa{
		{Nama:"ronal", NilaiTugas:90 , NilaiUAS: 100},
		{Nama: "iqbal", NilaiTugas: 70, NilaiUAS:80},
		{Nama: "bagas", NilaiTugas:60, NilaiUAS:50},
	}

	for _, m := range kelas{
	skorAkhir := (float64 (m.NilaiTugas)*0.3) + (float64(m.NilaiUAS)*0.7)
		if skorAkhir>= 70 {
		fmt.Printf ("Nama: %s | Skor:%2f |Status : LULUS\n " , m.Nama, skorAkhir )
	}else {
		fmt.Printf ("Nama: %s | Skor:%2f |Status : BELUM LULUS\n " , m.Nama, skorAkhir )
	}

	}
}




