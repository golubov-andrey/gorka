package queue

type entity struct {
	Name string
}

func New(name string) entity {
	return entity{Name: name}
}

