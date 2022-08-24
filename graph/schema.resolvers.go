package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"outerbanks-api/graph/generated"
	"outerbanks-api/graph/model"
	"strconv"
)

// UpsertCharacter is the resolver for the upsertCharacter field.
func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	id := input.ID
	var character model.Character
	character.Name = input.Name
	character.CliqueType = input.CliqueType

	n := len(r.Resolver.CharacterStore)
	if n == 0 {
		r.Resolver.CharacterStore = make(map[string]model.Character)
	}

	if id != nil {
		cs, ok := r.Resolver.CharacterStore[*id]
		if !ok {
			return nil, fmt.Errorf("character not found")
		}
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		} else {
			character.IsHero = cs.IsHero
		}
		character.ID = *id
		r.Resolver.CharacterStore[*id] = character
	} else {
		gid := strconv.Itoa(n + 1)
		character.ID = gid
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		}
		r.Resolver.CharacterStore[gid] = character
	}

	return &character, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	charector, ok := r.Resolver.CharacterStore[id]
	if !ok {
		return nil, fmt.Errorf("character not found")
	}

	return &charector, nil
}

// Characters is the resolver for the characters field.
func (r *queryResolver) Characters(ctx context.Context, cliqueType model.CliqueType) ([]*model.Character, error) {
	// filter data from r.Resolver.CharacterStore
	characters := make([]*model.Character, 0)
	for _, character := range r.Resolver.CharacterStore {
		if character.CliqueType == cliqueType {
			characters = append(characters, &character)
		}
	}

	return characters, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Pogues(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented: Pogues - pogues"))
}
func (r *queryResolver) Kooks(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented: Kooks - kooks"))
}
