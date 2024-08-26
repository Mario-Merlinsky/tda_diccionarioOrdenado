package diccionario

import (
	"tdas/cola"
	"tdas/pila"
)

const (
	CANT_INICIAL = 0
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iteradorAbb[K comparable, V any] struct {
	pila  pila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   func(a, b K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cantidad: CANT_INICIAL, cmp: funcion_cmp}
}

func crearNodoAbb[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: valor}
}

func (nodo *nodoAbb[K, V]) buscarLugarNodo(clave K, cmp func(K, K) int, padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return nil, padre
	}

	comparacion := cmp(clave, nodo.clave)

	if comparacion == 0 {
		return nodo, padre
	} else if comparacion < 0 {
		return nodo.izquierdo.buscarLugarNodo(clave, cmp, nodo)
	} else {
		return nodo.derecho.buscarLugarNodo(clave, cmp, nodo)
	}

}
func (abb *abb[K, V]) Guardar(clave K, dato V) {
	if abb.cantidad == CANT_INICIAL {
		abb.raiz = crearNodoAbb[K, V](clave, dato)
		abb.cantidad++
		return
	}

	nodo, padre := abb.raiz.buscarLugarNodo(clave, abb.cmp, nil)

	if nodo == nil {
		abb.cantidad++
		nodo = crearNodoAbb[K, V](clave, dato)
		if abb.cmp(nodo.clave, padre.clave) < 0 {
			padre.izquierdo = nodo
		} else {
			padre.derecho = nodo
		}
	}
	nodo.dato = dato
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := abb.raiz.buscarLugarNodo(clave, abb.cmp, nil)
	return nodo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo, _ := abb.raiz.buscarLugarNodo(clave, abb.cmp, nil)
	if nodo == nil {
		panic(ERROR_CLAVE)
	}
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo, padre := abb.raiz.buscarLugarNodo(clave, abb.cmp, nil)

	if nodo == nil {
		panic(ERROR_CLAVE)
	}

	dato := nodo.dato

	if nodo.izquierdo != nil && nodo.derecho != nil {
		clave_reemplazo, dato_reemplazo := buscar_reemplazante[K, V](nodo.izquierdo)
		abb.Borrar(clave_reemplazo)
		nodo.clave = clave_reemplazo
		nodo.dato = dato_reemplazo
		abb.cantidad++

	} else {
		if padre == nil {
			if nodo.izquierdo != nil {
				abb.raiz = nodo.izquierdo
			} else {
				abb.raiz = nodo.derecho
			}

		} else {
			if nodo.izquierdo != nil {
				nodo.izquierdo.herencia(nodo, padre, abb.cmp)
			} else {
				nodo.derecho.herencia(nodo, padre, abb.cmp)
			}
		}
	}
	abb.cantidad--
	return dato
}

func buscar_reemplazante[K comparable, V any](raiz *nodoAbb[K, V]) (K, V) {
	if raiz.derecho == nil {
		return raiz.clave, raiz.dato
	}
	return buscar_reemplazante(raiz.derecho)
}

func (nodo *nodoAbb[K, V]) herencia(padre *nodoAbb[K, V], abuelo *nodoAbb[K, V], cmp func(K, K) int) {
	if cmp(padre.clave, abuelo.clave) < 0 {
		abuelo.izquierdo = nodo
	} else {
		abuelo.derecho = nodo
	}
}

func (abb abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.raiz.iterarInOrder(nil, nil, visitar, abb.cmp)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterarInOrder(desde, hasta, visitar, abb.cmp)
}
func (nodo *nodoAbb[K, V]) iterarInOrder(desde *K, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}

	if desde != nil && cmp(nodo.clave, *desde) < 0 {
		return nodo.derecho.iterarInOrder(desde, hasta, visitar, cmp)
	} else if hasta != nil && cmp(nodo.clave, *hasta) > 0 {
		return nodo.izquierdo.iterarInOrder(desde, hasta, visitar, cmp)
	}

	if !nodo.izquierdo.iterarInOrder(desde, hasta, visitar, cmp) {
		return false
	}
	if !visitar(nodo.clave, nodo.dato) {
		return false
	}

	return nodo.derecho.iterarInOrder(desde, hasta, visitar, cmp)

}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iteradorAbb[K, V]{pila: pila.CrearPilaDinamica[*nodoAbb[K, V]]()}
	iter.apilarHerencia(abb.raiz)
	return iter
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter_rango := &iteradorAbb[K, V]{pila: pila.CrearPilaDinamica[*nodoAbb[K, V]](), desde: desde, hasta: hasta, cmp: abb.cmp}
	iter_rango.apilarHerencia(abb.raiz)
	return iter_rango
}
func (iter *iteradorAbb[K, V]) apilarHerencia(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	if iter.desde != nil && iter.cmp(nodo.clave, *iter.desde) < 0 {
		iter.apilarHerencia(nodo.derecho)
	} else if iter.hasta != nil && iter.cmp(nodo.clave, *iter.hasta) > 0 {
		iter.apilarHerencia(nodo.izquierdo)
	} else {
		iter.pila.Apilar(nodo)
		iter.apilarHerencia(nodo.izquierdo)
	}

}
func (iter iteradorAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter iteradorAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iteradorAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.Desapilar()
	iter.apilarHerencia(nodo.derecho)
}

func (abb *abb[K, V]) IterNivelesInverso(visitar func(clave K, _ V) bool) {
	cola := cola.CrearColaEnlazada[*nodoAbb[K, V]]()
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()

	cola.Encolar(abb.raiz)

	for !cola.EstaVacia() {
		nodo := cola.Desencolar()
		if nodo.izquierdo != nil {
			cola.Encolar(nodo.izquierdo)
		}
		if nodo.derecho != nil {
			cola.Encolar(nodo.derecho)
		}
		pila.Apilar(nodo)
	}

	for !pila.EstaVacia() {
		nodo := pila.Desapilar()
		clave, dato := nodo.clave, nodo.dato

		visitar(clave, dato)
	}
}

func (abb *abb[K, V]) N_Descendientes(n int) int {
	cantidad := 0
	abb.raiz.N_Descendientes_aux(n, &cantidad)
	return cantidad
}

func (nodo *nodoAbb[K, V]) N_Descendientes_aux(n int, cantidad *int) int {
	if nodo == nil {
		return 0
	}
	izq := nodo.izquierdo.N_Descendientes_aux(n, cantidad)
	der := nodo.derecho.N_Descendientes_aux(n, cantidad)

	if izq+der == n {
		*cantidad++
	}
	return izq + der + 1
}
