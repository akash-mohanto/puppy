package puppy

import "github.com/akash-mohanto/dog"

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Barks())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
