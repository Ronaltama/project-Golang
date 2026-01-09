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
	nilai :=hitungSkor (m.NilaiTugas , m.NilaiUAS)
	status := tentukanStatus(nilai)
	fmt.Printf ("Nama: %s | Skor:%.2f |Status : %s \n" , m.Nama, nilai, status )	
	}
}


func hitungSkor(tugas int, uas int)float64 {
	skorAkhir := (float64(tugas) * 0.3) + (float64 (uas) * 0.7)
	return skorAkhir
}

func tentukanStatus( skor float64) string {
if skor>= 70 {
	return "LULUS"
	} else {
	return " TIDAK LULUS"
	}
}


