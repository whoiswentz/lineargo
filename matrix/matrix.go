package matrix

type Matrix struct {
	matrix  [][]float64
	rows    int
	columns int
}

// Return a new matrix of given slice of slices
func NewMatrix(m [][]float64) *Matrix {
	return &Matrix{matrix: m, rows: len(m), columns: len(m[0])}
}

// Create a new matrix of zeros
func CreateMatrix(row, columns int) *Matrix {
	matrix := make([][]float64, row)
	for i := range matrix {
		matrix[i] = make([]float64, columns)
	}
	return &Matrix{matrix: matrix, rows: row, columns: columns}
}

func Identity(row, column int) *Matrix {
	return &Matrix{}
}

func (m Matrix) Multiply(o Matrix) *Matrix {
	if m.columns != o.rows {
		panic("The number of rows must be equals to the number of rows of the matrix")
	}

	out := make([][]float64, len(m.matrix))
	for x := range m.matrix {
		for y := range o.matrix {
			out[x] = make([]float64, len(o.matrix))
			out[x][y] += m.matrix[x][y] * o.matrix[y][x]
		}
	}

	return &Matrix{matrix: out, rows: m.rows, columns: m.columns}
}

func (m Matrix) Compare(o Matrix) bool {
	for x := range m.matrix {
		for y := range o.matrix[0] {
			if m.matrix[x][y] != o.matrix[x][y] {
				return false
			}
		}
	}
	return true
}

// Create a matrix with all elements zeros
// TODO test this function
func CreateZeroMatrix(rows, columns int) *Matrix {

	return &Matrix{matrix: [][]float64{}, rows: rows, columns: columns}
}
