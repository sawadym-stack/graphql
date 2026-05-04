package graph

import (
	"sync"
	"gqlgen-app/graph/model"
)

type Resolver struct{
	mu            sync.Mutex
	postObservers map[string]chan *model.Post
}