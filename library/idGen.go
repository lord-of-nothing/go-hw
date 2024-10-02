package library

func GenerateAllId() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

func GenerateEvenId() func() int {
	id := 0
	return func() int {
		id += 2
		return id
	}
}
