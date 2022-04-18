package interactor

import (
	"fmt"
)

type HomepageInteractor struct{}

func (hp *HomepageInteractor) Index() string {
	fmt.Println("Endpoint Hit: homePage")
	return "Welcome to the HomePage!"

}
