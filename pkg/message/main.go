package message

type entity struct {
	Body interface{}
}

func New(body interface{}) entity {
	return entity{Body: body}
}
