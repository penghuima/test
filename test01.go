package dog

import (
	"fmt"
)

type Dog struct{
	Name string
}

func (d *Dog) CallMyname() {
	fmt.Printf("The Dog's name is: %q\n", d.Name)
}
