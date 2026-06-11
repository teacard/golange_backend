package dto

// MapSlice 將 []T 透過 mapper 轉換成 []U，取代手寫 for 迴圈
func MapSlice[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}
