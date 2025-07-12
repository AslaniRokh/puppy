package puppy

import (
	"github.com/AslaniRokh/dog"
)

func Bark() string {
	return "waaaagh!"
}

func Barcks() string {
	return "waaaagh waaaagh waaaagh!"
}

func BigBarck() string {
	return dog.BigDog(Bark())
}

func BigBarcks() string {
	return dog.BigDog(Barcks())
}
