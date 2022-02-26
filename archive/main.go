package main

const (
	host   = "localhost"
	port   = 5432
	user   = "nithin"
	dbname = "unsploosh"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	gorm_main()
}
