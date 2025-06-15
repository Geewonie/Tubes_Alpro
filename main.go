package main
import "fmt"

const Max = 100

const (
	Lunas string = "Lunas"
	BelumLunas string = "Belum Lunas"
)

type Peminjam struct {
	Nama string
	Jumlah int
	Tenor int
	BungaTetap bool
	Status string
	Bunga float64
	Cicilan float64
}

type DaftarPeminjam struct {
	Data [Max]Peminjam
	Jumlah int
}

var daftar DaftarPeminjam


func TambahData(p *DaftarPeminjam, data Peminjam) {
	if p.Jumlah < Max {
		data.Bunga = HitungBunga(data.Jumlah, data.Tenor, data.BungaTetap)
		data.Cicilan = HitungCicilan(data.Jumlah, data.Bunga, data.Tenor)
		p.Data[p.Jumlah] = data
		p.Jumlah++
	}
}

func HitungBunga(jumlah int, tenor int, BungaTetap bool) float64 {
	if BungaTetap {
		return float64(jumlah) * 0.1
	} else {
		return float64(jumlah) * 0.05 * float64(tenor)
	}
}

func HitungCicilan(jumlah int, bunga float64, tenor int) float64 {
	return (float64(jumlah) + bunga) / float64(tenor)
}

func SequentialSearch(p *DaftarPeminjam, nama string) int {
	for i := 0; i < p.Jumlah; i++ {
		if p.Data[i].Nama == nama {
			return i
		}
	}
	return -1
}

func BinarySearch(p *DaftarPeminjam, nama string) int {
	PenyusunanNama(p)
	low := 0
	high := p.Jumlah -1
	for low <= high {
		mid := (low + high) / 2
		if p.Data[mid].Nama == nama{
			return mid
		} else if p.Data[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func PenyusunanNama(p *DaftarPeminjam) {
	for i := 1; i < p.Jumlah; i++ {
		key := p.Data[i]
		j := i - 1
		for j >= 0 && p.Data[i].Nama < key.Nama {
			p.Data[j+1] = p.Data[j]
			j--
		}
		p.Data[j+1] = key
	}
}

func SelectionSortJumlah(p *DaftarPeminjam, target bool) {
	for i := 0; i < p.Jumlah; i++ {
		index := i
		for j := i + 1; j < p.Jumlah; j++ {
			if (target && p.Data[j].Jumlah < p.Data[index].Jumlah) || (!target && p.Data[j].Jumlah > p.Data[index].Jumlah) {
				index = j
			}
		} 
		p.Data[i], p.Data[index] = p.Data[index], p.Data[i]
	}
}

func SelectionSortNama(p *DaftarPeminjam, target bool) {
	for i := 0; i < p.Jumlah; i++ {
		index := i
		for j := i + 1; j < p.Jumlah; j++ {
			if (target && p.Data[j].Nama < p.Data[index].Nama) || (!target && p.Data[j].Nama > p.Data[index].Nama) {
				index = j
			}
		}
		p.Data[i], p.Data[index] = p.Data[index], p.Data[i]
	}
}

func Laporan(p *DaftarPeminjam) {
	total := 0
	fmt.Println("LAPORAN PINJAMAN:")
	for i := 0; i < p.Jumlah; i++ {
		fmt.Printf("%d. %s - Pinjaman: %d, Cicilan: %.2f, Status: %s\n",
			i+1, p.Data[i].Nama, p.Data[i].Jumlah,
			p.Data[i].Cicilan, p.Data[i].Status)
		total += p.Data[i].Jumlah
	}
	fmt.Println("Total Pinjaman Diberikan: ", total) 
}

func HapusData (p *DaftarPeminjam, index int) {
	if index >= 0 && index < p.Jumlah {
		for i := index; i < p.Jumlah-1; i++ {
			p.Data[i] = p.Data[i+1]
		}
		p.Jumlah--
	}
}

func UbahData (p *DaftarPeminjam, index int, data Peminjam) {
	if index >= 0 && index < p.Jumlah {
		data.Bunga = HitungBunga(data.Jumlah, data.Tenor, data.BungaTetap)
		data.Cicilan = HitungCicilan(data.Jumlah, data.Bunga, data.Tenor)
		p.Data[index] = data
	}
}




func main() {
	TambahData(&daftar, Peminjam{"Aldi", 10000000, 12, true, BelumLunas, 0, 0})
	TambahData(&daftar, Peminjam{"Budi", 8000000, 10, false, BelumLunas, 0, 0})
	TambahData(&daftar, Peminjam{"Citra", 15000000, 24, true, Lunas, 0, 0})

	Laporan(&daftar)
	fmt.Println()
	




}