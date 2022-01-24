// Copyright (c) 2022 Golubov Andrey
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package queue

type entity struct {
	Name string
}

func New(name string) entity {
	return entity{Name: name}
}
