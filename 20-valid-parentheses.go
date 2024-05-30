package main

type Stack struct {
	items []string
}

func (s *Stack) length() int {
	return len(s.items)
}

func (s *Stack) Push(el string) {
	s.items = append(s.items, el)
}

func (s *Stack) Pop() string {
	lastIdx := len(s.items) - 1
	lastItem := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return lastItem
}

func (s *Stack) Top() string {
	return s.items[s.length()-1]
}

func isValid(s string) bool {

	hashmap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	myStack := Stack{}

	for _, char := range s {

		if _, ok := hashmap[string(char)]; ok {

			if myStack.length() > 0 && (myStack.Top() == hashmap[string(char)]) {
				myStack.Pop()
			} else {
				return false
			}
		} else {
			myStack.Push(string(char))
		}

	}

	return myStack.length() == 0
}
