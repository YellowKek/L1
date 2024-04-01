package main

//var justString string

func someFunc() string {
	v := createHugeString(1 << 10)
	temp := make([]byte, 100)
	copy(temp, v[:100])
	//justString = v[:100]

	// вместо того чтобы хранить в глобальной перемменной подмассив большой строки возвращаем копию
	// благодаря этому v не будет постоянно занимать память
	return string(temp)
}

func main() {
	someFunc()
}
