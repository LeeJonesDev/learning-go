package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// IsInRange tests if an index is in a slices length
func IsInRange(index int, length int) bool {
	return 0 <= index && length > index
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if IsInRange(index, len(slice)) {
		return slice[index]
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if IsInRange(index, len(slice)) {
		slice[index] = value
		return slice
	}
	return append(slice, value)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var start, end []int
	length := len(slice)

	//check index is in range
	if index < 0 || index > length-1 {
		return slice
	}

	start = slice[:index]
	if IsInRange(index+1, length) {
		end = slice[index+1:]
	}

	return append(start, end...)
}
