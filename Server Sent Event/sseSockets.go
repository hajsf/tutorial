package main

import "fmt"

type Sockets []string

func (s *Sockets) RemoveElement(e string) error {
	//slice := new([]string)
	i, err := s.FindIndex(e)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	s.RemoveIndex(i)
	return nil
}

func (s *Sockets) FindIndex(e string) (int, error) {
	for i, v := range *s {
		if v == e {
			return i, nil
		}
	}
	return -1, fmt.Errorf("element %v not found", e)
}

func (s *Sockets) RemoveIndex(index int) {
	sockets := *s
	sockets = append(sockets[:index], sockets[index+1:]...)
	*s = sockets
}
