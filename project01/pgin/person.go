package pgin

type Person struct {
	Id   int    `json:"id" binding:"required"`
	Name SName `json:"name"`
	Age  int    `json:"age"`
}

type SName struct {
	Name1 string `json:"name1" binding:"required"`
	Name2 string `json:"name2" binding:"required"`
}