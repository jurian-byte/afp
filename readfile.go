package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func generarAFP() *Estado {
	file, err := ioutil.ReadFile("./afp2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\r\n")
	/* for _, line := range lines {
		fmt.Printf("'%s'\n", line)
	} */
	nEstados := strings.Split(lines[0][2:], ",")
	nEstadoI := lines[1][3:]
	nEstadosF := strings.Split(lines[2][2:], ",")
	fmt.Println("Estados:", nEstados)         // []string : nombres de los estados
	fmt.Println("Estado inicial: ", nEstadoI) // string : nombre del estado inicial
	fmt.Println("Estados finales", nEstadosF) // []string : nombres de los estados finales

	estados := make([]*Estado, len(nEstados)) // []*Estado : Los estados que se ocupan en el AFP

	var posInicial int
	for i := range estados {
		estados[i] = &Estado{
			Nombre: nEstados[i],
			Final:  isFinal(nEstados[i], nEstadosF),
		}
		if estados[i].Nombre == nEstadoI {
			posInicial = i
		}
	}
	// Aqui ya tenemos la lista de los estados, cual es el inicial y cuales los finales
	// Falta transiciones.

	for _, nTransicion := range lines[4:] {
		nTransicion := strings.TrimSpace(nTransicion)
		if len(nTransicion) > 2 {
			delta01 := strings.Split(nTransicion, "=") // []string: delta[0]=ORIG,a,b delta[1]=DEST,c1,c2,...,cn
			delta0 := strings.Split(delta01[0], ",")   // ORIG,a,b
			delta1 := strings.Split(delta01[1], ",")   // DEST,c1,c2,...,cn
			setTransicion(estados, delta0[0], delta1[0], delta0[1], delta0[2], delta1[1:])
		}
	}
	return estados[posInicial]
}

func setTransicion(estados []*Estado, origen, dest, a, b string, c []string) {
	var po, pd int //posicion del estado origen, posicion del destino
	for i, estado := range estados {
		if estado.Nombre == origen {
			po = i
		}
		if estado.Nombre == dest {
			pd = i
		}
	}
	t := Transicion{
		Destino: estados[pd],
		A:       getRune(a),
		B:       getRune(b),
		C:       getStringRuned(c),
	}
	fmt.Printf("δ(%s,%s,%s)=(%s,%s)\n", origen, string(t.A), string(t.B), dest, t.C)
	estados[po].Transiciones = append(estados[po].Transiciones, t)
}

func getStringRuned(cs []string) string {
	c := ""
	for _, cn := range cs {
		c += string(getRune(cn))
	}
	return c
}

func getRune(s string) rune {
	switch s {
	case "":
		panic("Cadena vacia no valida")
	case string(Ep):
		return Ep
	case "Z0":
		return Z0
	default:
		r := []rune(s)
		return r[0]
	}
}

func isFinal(nEstado string, nEstadosF []string) bool {
	for _, e := range nEstadosF {
		if e == nEstado {
			return true
		}
	}
	return false
}
