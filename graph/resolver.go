package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"graphqlExample/graph/model"
)

type Resolver struct {
	CharacterStore map[string]model.Character
}
