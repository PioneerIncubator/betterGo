package slice

func RemoveStringElement(slice []string, key string) []string {
	var index = 0
	for k, v := range slice {
		if v == key {
			index = k
			break
		}
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func RemoveIntElement(slice []int, key int) []int {
	var index = 0
	for k, v := range slice {
		if v == key {
			index = k
			break
		}
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}
