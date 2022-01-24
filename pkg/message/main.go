// Copyright (c) 2022 Golubov Andrey
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package message

type entity struct {
	Body interface{}
}

func New(body interface{}) entity {
	return entity{Body: body}
}
