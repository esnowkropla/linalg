package linalg

import (
	"errors"
	"strconv"
)

type Matrix struct {
	row, col int
	Mij      []complex128
}

func (M *Matrix) String() string {
	ret := ""
	for i := 0; i < M.row; i++ {
		for j := 0; j < M.col; j++ {
			ret += strconv.FormatFloat(real(M.Elem(i, j)), 'g', -1, 64) + " "
			ret += strconv.FormatFloat(imag(M.Elem(i, j)), 'g', -1, 64) + "j"
			if j < M.col-1 {
				ret += ", "
			}
		}
		ret += "\n"
	}
	return ret
}

/* Access Functions */
func (M *Matrix) Elem(row, col int) complex128 {
	return M.Mij[row+col*M.row]
}

func (M *Matrix) Set_elem(row, col int, val complex128) {
	M.Mij[row+col*M.row] = val
}

func (M *Matrix) Set_int(row, col, val int) {
	M.Set_elem(row, col, complex(float64(val), 0))
}

func (M *Matrix) Set_float(row, col int, val float64) {
	M.Set_elem(row, col, complex(val, 0))
}

/* Creation Functions */
func Ident(dim int) *Matrix {
	var mat = new(Matrix)
	mat.row = dim
	mat.col = dim
	mat.Mij = make([]complex128, dim*dim)
	for i := 0; i < dim; i++ {
		mat.Set_elem(i, i, 1)
	}

	return mat
}

func Zero(row, col int) *Matrix {
	var M = new(Matrix)
	M.row = row
	M.col = col
	M.Mij = make([]complex128, row*col)
	return M
}

func (M *Matrix) Copy() *Matrix {
	var N = Zero(M.col, M.row)
	copy(N.Mij, M.Mij)

	return N
}

func Init(row, col int, data []complex128) *Matrix {
	var M = Zero(row, col)
	M.Mij = data
	return M
}

/* Math Functions */
func (M *Matrix) Eq(N *Matrix) bool {
	if (M.row != N.row) || (M.col != N.col) {
		return false
	}
	for i := 0; i < M.col; i++ {
		for j := 0; j < M.row; j++ {
			if M.Elem(i, j) != N.Elem(i, j) {
				return false
			}
		}
	}
	return true
}

func Mul(A, B, C *Matrix) error {
	if A.col != B.row || A.row != C.row || B.col != C.col {
		return errors.New("Matrix Multiplication size mismatch")
	}

	for i := 0; i < C.row; i++ {
		for j := 0; j < C.col; j++ {
			var accum complex128
			for k := 0; k < A.col; k++ {
				accum += A.Elem(i, k) * B.Elem(k, j)
			}
			C.Set_elem(i, j, accum)
		}
	}
	return nil
}

func Add(A, B, C *Matrix) error {
	if A.col != B.col || A.col != C.col || A.row != B.row || A.row != C.row {
		return errors.New("Matrix Addition size mismatch")
	}

	for i := 0; i < A.col; i++ {
		for j := 0; j < A.row; j++ {
			C.Set_elem(i, j, A.Elem(i, j)+B.Elem(i, j))
		}
	}
	return nil
}

func (M *Matrix) Scale(scalar complex128) {
	for i := 0; i < M.col*M.row; i++ {
		M.Mij[i] *= scalar
	}
}
