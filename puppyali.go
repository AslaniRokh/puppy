package puppy

import (
	"github.com/AslaniRokh/dog"
)

func Bark() string {
	return "woof!"
}

func Barcks() string {
	return "woof woof woof!"
}

func BigBarck() string {
	return dog.BigDog(Bark())
}

func BigBarcks() string {
	return dog.BigDog(Barcks())
}
