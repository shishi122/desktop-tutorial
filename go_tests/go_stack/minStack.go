package gostack

import (
	"fmt"
)

// Структура стека
type Stack[t any] struct {
	objs []t
}

// Метод Push для вставки элементов
func (s *Stack[t]) Push(obj t) {
	s.objs = append(s.objs, obj)
}

// Метод Pop для доставания последнего элемента стека
func (s *Stack[t]) Pop() (res t, err error) {
	if len(s.objs) != 0 {
		res = s.objs[len(s.objs)-1]
		s.objs = s.objs[:len(s.objs)-1]
	} else {
		err = fmt.Errorf("Stack is empty")
	}
	return
}

// Метод Peek возвращает последний элемент но не удвляет его
func (s *Stack[t]) Peek() (res t, err error) {
	if len(s.objs) != 0 {
		res = s.objs[len(s.objs)-1]
	} else {
		err = fmt.Errorf("Stack is empty")
	}
	return
}

// Метод IsEmpty проверяет пуст ли стек
func (s *Stack[t]) IsEmpty() bool {
	if len(s.objs) == 0 {
		return true
	}
	return false
}
