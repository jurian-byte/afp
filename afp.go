package main

import "fmt"

const (
	// Z0 indica el fondo de la pila
	Z0 = '$'
	// Ep epsilon caracter vacio
	Ep = 'ɛ'
)

// Estado del automata
type Estado struct {
	Nombre       string
	Final        bool
	Transiciones []Transicion
}

// Transicion representa el cambio de estado
type Transicion struct {
	Destino *Estado
	A       rune
	B       rune
	C       string
}

// Siguientes Regresa las transiciones que se realizan con el caracter dado
func (e Estado) Siguientes(r rune) []Transicion {
	var transiciones []Transicion
	for _, t := range e.Transiciones {
		if t.A == r || t.A == Ep {
			transiciones = append(transiciones, t)
		}
	}
	return transiciones
}

// ValidarCadena verifica si la cadena pertenece al lenguaje dado por el AFP
func (e *Estado) ValidarCadena(s string) bool {
	stack := NewStack()
	stack.Push(Z0)
	return e.analizar(s, stack)
}

// analizar a
func (e *Estado) analizar(s string, stack Stack) bool {
	sOriginal := s
	var r rune
	if e.Final && len(s) == 0 {
		imprimeEstado(e, sOriginal, stack, true)
		return true
	}
	if len(s) == 0 {
		r = Ep
	} else {
		r = rune(s[0])
		s = s[1:]
	}
	transiciones := e.Siguientes(r)
	for _, t := range transiciones {
		stack2 := stack.Copy()
		if t.B != Ep {
			top, ok := stack2.Pop()
			if !ok { // la pila no tiene elementos
				continue
			}
			if top != t.B {
				continue
			}
			stack2.PushString(t.C)
			if t.A == Ep {
				ok = t.Destino.analizar(sOriginal, stack2)
			} else {
				ok = t.Destino.analizar(s, stack2)
			}
			if ok {
				imprimeEstado(e, sOriginal, stack, true)
				return true
			}
			imprimeEstado(e, sOriginal, stack, false)
		}
	}
	imprimeEstado(e, sOriginal, stack, false)
	return false
}

func imprimeEstado(e *Estado, s string, stack Stack, ok bool) {
	if e.Final {
		fmt.Printf("(|%s|,", e.Nombre)
	} else {
		fmt.Printf("(%s,", e.Nombre)
	}
	if ok {
		fmt.Printf("%s,%s) ok\n", s, string(stack))
	} else {
		fmt.Printf("%s,%s) x\n", s, string(stack))
	}
}
