package main

import "fmt"

// ngebuat iinterface, yaitu "mahasiswa"nya
type Perkuliahan interface {
    Info() string
}

// ini untuk menentukan struct yang jadi dasar object" 
type Mahasiswa struct {
    Name  string
    Grade int
}

// untuk engambil informasi
func (m *Mahasiswa) Info() string {
    return fmt.Sprintf("Nama: %s, Nilai: %d", m.Name, m.Grade)
}

// ini struct matakulihanya
type Matkul struct {
    Name  string
    Units int
}

// ambil informasi matkul
func (M *Matkul) Info() string {
    return fmt.Sprintf("Mata Kuliah: %s, SKS: %d", M.Name, M.Units)
}

// Factory
func CreateObject(objectType string, name string, gradeOrUnits int) Perkuliahan {
    if objectType == "mahasiswa" {
        return &Mahasiswa{Name: name, Grade: gradeOrUnits}
    } else if objectType == "mata kuliah" {
        return &Matkul{Name: name, Units: gradeOrUnits}
    }
    return nil
}

func main() {
	//hasilnya
    Mahasiswa := CreateObject("mahasiswa", "John", 85)
    Matkul := CreateObject("mata kuliah", "Pemrograman", 3)
    fmt.Println(Mahasiswa.Info())
    fmt.Println(Matkul.Info())
}