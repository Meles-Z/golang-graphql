package interfaces

import (
	"github.com/meles-z/golang-graphql/app/domain/repository"
)

type Resolver struct {
	MovieRepo repository.MovieRepository
	UserRepo  repository.UserRepository
}
