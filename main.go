package rackrock

func main() {
	r := Setup()

	// routing paths here

	r.Run(":3001")
}
