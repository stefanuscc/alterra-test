package alterra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Soal 1 Frasa check
func frasa(kata string) bool {
	// Iterate huruf by huruf
	for i := 0; i < len(kata)/2; i++ {
		// check apakah huruf awal dengan huruf akhir sama
		if kata[i] != kata[len(kata)-i-1] {
			return false
		}
	}

	// if all looks same, return true
	return true
}

func resursiveFrasa(kata string, idx int) bool {
	// Check jika index sudah lewat tengah kalimat
	if idx >= len(kata)/2 {
		return true
	}

	// Check jika kata awal tidak sama dengan kata akhir
	if kata[idx] != kata[len(kata)-idx-1] {
		return false
	}

	// Iterasi rekursif
	return resursiveFrasa(kata, idx+1)
}

func isFrasa(kata string) bool {
	return resursiveFrasa(kata, 0)
}

func TestFrasa(t *testing.T) {
	assert.True(t, frasa("katak"))
	assert.False(t, frasa("basi"))
	assert.True(t, frasa("isi"))
	assert.True(t, frasa("racecar"))
	assert.False(t, frasa("something"))
	assert.True(t, frasa("kodok"))

	// t.Error("Frasa: katak, Result:", frasa("basi"))
}

func TestFrasaRec(t *testing.T) {
	assert.True(t, isFrasa("kodok"))
	assert.False(t, isFrasa("basi"))
	assert.True(t, isFrasa("isi"))
	assert.True(t, isFrasa("racecar"))
	assert.False(t, isFrasa("something"))
	assert.True(t, isFrasa("kodok"))
}

// Soal 2 Penjumlahan diagonal array
func diagonalCalc(matrix [][]int) int {
	// Init total variable
	total := 0
	size := len(matrix) - 1

	// iterasi baris
	for i := 0; i <= size; i++ {
		// Jumlahkan diagonal kiri
		total += matrix[i][i]
		// Jumlahkan diagonal kanan
		total += matrix[i][size-i]
	}

	// Tampilkan total
	return total
}

func TestDiagonalCalc(t *testing.T) {
	matrix := [][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, 12},
	}

	assert.Equal(t, 47, diagonalCalc(matrix))

	matrix2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	assert.Equal(t, 30, diagonalCalc(matrix2))

	matrix3 := [][]int{
		{1, 2, 3, 1},
		{4, 5, 6, 1},
		{7, 8, 9, 1},
		{2, 3, 4, 5},
	}

	assert.Equal(t, 37, diagonalCalc(matrix3))
}
