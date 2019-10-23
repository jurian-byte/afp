package main

// Stack pila lifo
type Stack []rune

// NewStack Create a new stack
func NewStack() Stack {
	return make(Stack, 0)
}

// Copy crea una copia del stack
func (stack Stack) Copy() Stack {
	return []rune(string(stack))
}

// Len Return the number of items in the stack
func (stack Stack) Len() int {
	return len(stack)
}

// Peek View the top item on the stack
func (stack Stack) Peek() (rune, bool) {
	l := len(stack)
	if l == 0 {
		return Ep, false
	}
	return stack[l-1], true
}

// Pop the top item of the stack and return it
func (stack *Stack) Pop() (rune, bool) {
	s := *stack
	l := len(s)
	if l == 0 {
		return Ep, false
	}
	r := s[l-1]
	*stack = s[:l-1]
	return r, true
}

// Push a value onto the top of the stack
func (stack *Stack) Push(r rune) {
	s := *stack
	*stack = append(s, r)
}

// PushString a value onto the top of the stack
func (stack *Stack) PushString(s string) {
	runes := []rune(s)
	st := *stack
	for i := len(runes) - 1; i >= 0; i-- {
		r := rune(runes[i])
		if r == Ep {
			continue
		}
		*stack = append(st, r)
		st = *stack
	}
}
