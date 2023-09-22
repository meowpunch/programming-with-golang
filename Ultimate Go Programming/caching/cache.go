package caching

import "fmt"

const (
	rows = 2 * 1024
	cols = 2 * 1024
)

var matrix [rows][cols]byte

type data struct {
	v byte
	p *data
}

var h *data

func init() {
	var c = h

	// Create a link list with the same number of elements.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			// Create a new node and link it in.
			var d data
			if c == nil {
				c = &d
			}

			c = c.p

			// Add a value to all even elements.
			if row%2 == 0 {
				matrix[row][col] = 0xFF
				d.v = 0xFF
			}
		}
	}

	// Count the number of elements in the link list.
	var ctr int
	d := h
	for d != nil {
		ctr++
		d = d.p
	}

	fmt.Println("Elements in the link list", ctr)
	fmt.Println("Elements in the matrix", rows*cols)
}

// LinkedListTraverse traverses the linked list linearly.
func LinkedListTraverse() int {
	var ctr int

	d := h
	for d != nil {
		if d.v == 0xFF {
			ctr++
		}

		d = d.p
	}

	return ctr
}

// ColumnTraverse traverses the matrix linearly down each column.
func ColumnTraverse() int {
	var ctr int

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

// RowTraverse traverses the matrix linearly down each row.
func RowTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}
