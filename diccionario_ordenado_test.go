package diccionario_test

import (
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmp(a int, b int) int {
	return a - b
}

func TestDiccionarioAbbVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioAbbClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](cmp)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioAbbGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatoAbb(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioAbbBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorradosAbb(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericasAbb(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](cmp)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestClaveVaciaAbb(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNuloAbb(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestVolumenAbb(t *testing.T) {
	t.Log("Prueba guardar muschos elementos, luego se asegura que los elementos pertenezcan al diccionario y luego los borra" +
		"alfina se asegura que el diccionario termine vacio")
	abb := TDADiccionario.CrearABB[int, int](cmp)
	dict := TDADiccionario.CrearHash[int, int]()

	for i := 0; i < 1000000; i++ {
		numero := rand.Intn(1000000)
		dict.Guardar(numero, numero)
		abb.Guardar(numero, numero)
	}

	for iter := dict.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		require.True(t, abb.Pertenece(clave))
		require.EqualValues(t, valor, abb.Obtener(clave))
		require.EqualValues(t, valor, abb.Borrar(clave))
	}
	require.False(t, abb.Pertenece(0))

}

func TestIteradorInternoInOrderClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoInOrderValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoInOrderValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIterarDiccionarioAbbVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterDiccionarioAbb(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorAbbNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestPruebaIterarAbbTrasBorrados(t *testing.T) {
	t.Log("Prueba que tras borrar elementos se pueda iterar el Diccionario con normalidad")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestVolumenIterarAbbCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](cmp)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIterarRangoCompletoEnOrden(t *testing.T) {
	t.Log("Prueba que se pueda usar el iterador internopor rango de principio a fin con cotas fuera del rango" +
		"que posee el arbol")
	abb := TDADiccionario.CrearABB[int, int](cmp)
	slice_orden := []int{1, 5, 7, 10, 15, 20, 25}
	slice_insertar := []int{10, 20, 25, 15, 7, 5, 1}
	for _, valor := range slice_insertar {
		abb.Guardar(valor, valor)
	}
	desde := 0
	hasta := 30
	b := 0
	i := &b
	abb.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		require.EqualValues(t, slice_orden[*i], clave)
		require.EqualValues(t, slice_orden[*i], dato)
		*i++
		return true
	})
}

func TestIterarRangoIncompleto(t *testing.T) {
	t.Log("Prueba que se pueda iterar solo en un rango")
	abb := TDADiccionario.CrearABB[int, int](cmp)
	abb.Guardar(20, 20)
	abb.Guardar(15, 15)
	abb.Guardar(30, 30)
	abb.Guardar(18, 18)
	abb.Guardar(2, 2)
	abb.Guardar(35, 35)
	abb.Guardar(45, 45)

	desde := 4
	hasta := 25
	suma := 0
	ptrsuma := &suma

	abb.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		*ptrsuma += clave
		return true

	})

	require.EqualValues(t, 53, suma)
}

func TestIterarAcotadoSuperior(t *testing.T) {
	t.Log("Test para probar si se puede iterar correctamente sin una cota inferior del rango pero con una superior")
	abb := TDADiccionario.CrearABB[int, int](cmp)

	abb.Guardar(500, 500)
	abb.Guardar(300, 300)
	abb.Guardar(700, 700)
	abb.Guardar(250, 250)
	abb.Guardar(400, 400)

	arr := []int{250, 300, 400}
	i := 0
	ptri := &i
	suma := 0
	ptrsuma := &suma
	hasta := 450
	abb.IterarRango(nil, &hasta, func(clave, dato int) bool {
		require.EqualValues(t, arr[*ptri], clave)
		*ptrsuma += clave
		*ptri++
		return true
	})
	require.EqualValues(t, 950, suma)
}

func TestIterarAcotadoInferior(t *testing.T) {
	t.Log("Test que prueba que se pueda iterar por rango sin una cota superior pero con una inferior")
	abb := TDADiccionario.CrearABB[int, int](cmp)

	abb.Guardar(50, 50)
	abb.Guardar(-2, -2)
	abb.Guardar(40, 40)
	abb.Guardar(70, 70)
	abb.Guardar(85, 85)

	arr := []int{50, 70, 85}
	i := 0
	ptri := &i
	suma := 0
	ptrsuma := &suma
	desde := 50
	abb.IterarRango(&desde, nil, func(clave, dato int) bool {
		require.EqualValues(t, arr[*ptri], clave)
		*ptrsuma += clave
		*ptri++
		return true
	})
	require.EqualValues(t, 205, suma)
}

func TestIteradorExtConRango(t *testing.T) {
	t.Log("Se prueba que el iterador externo funcione en un rango acotado")
	abb := TDADiccionario.CrearABB[int, int](cmp)

	abb.Guardar(10, 10)
	abb.Guardar(20, 20)
	abb.Guardar(5, 5)
	abb.Guardar(25, 25)
	abb.Guardar(15, 15)
	abb.Guardar(8, 8)
	abb.Guardar(3, 3)
	slice_ord := []int{5, 8, 10, 15, 20}

	desde := 4
	hasta := 21
	suma := 0
	ptrsuma := &suma
	b := 0
	i := &b
	for iter := abb.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		require.EqualValues(t, slice_ord[*i], clave)
		*i++
		*ptrsuma += clave
	}
	require.Equal(t, 58, suma)
}

func TestIteradorExternoAcotadoSuperior(t *testing.T) {
	t.Log("Prueba que esl iterador funcione correctamente si solo esta acotado superiormente")
	abb := TDADiccionario.CrearABB[int, int](cmp)
	abb.Guardar(11, 11)
	abb.Guardar(10, 10)
	abb.Guardar(7, 7)
	abb.Guardar(1, 1)
	abb.Guardar(8, 8)
	abb.Guardar(12, 12)
	abb.Guardar(13, 13)
	arr := []int{1, 7, 8, 10, 11}
	hasta := 11
	pntrhasta := &hasta
	iter := abb.IteradorRango(nil, pntrhasta)
	for i := 0; iter.HaySiguiente(); i++ {
		clave, _ := iter.VerActual()
		require.EqualValues(t, arr[i], clave)
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
}

func TestIteradorExternoAcotadoInferior(t *testing.T) {
	t.Log("Prueba que el iterador funcione correctamente sin ser acotado superiormente")
	abb := TDADiccionario.CrearABB[int, int](cmp)
	abb.Guardar(20, 20)
	abb.Guardar(15, 15)
	abb.Guardar(10, 10)
	abb.Guardar(25, 25)
	abb.Guardar(22, 22)
	abb.Guardar(30, 30)
	arr := []int{20, 22, 25, 30}
	desde := 20
	pntrdesde := &desde
	iter := abb.IteradorRango(pntrdesde, nil)
	for i := 0; iter.HaySiguiente(); i++ {
		clave, _ := iter.VerActual()
		require.EqualValues(t, arr[i], clave)
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())

}

func TestVolumenIteradorExterno(t *testing.T) {
	t.Log("Prueba iterar con el iterador esterno muchos elementos y ademas se asegura que los elemenos vayan en inorder")
	abb := TDADiccionario.CrearABB[int, int](cmp)
	dict := TDADiccionario.CrearHash[int, int]()
	n := -1 //variable para saber si se esta iterando en orden segun la funcion de comparacion
	for i := 0; i < 100000; i++ {
		numero := rand.Intn(100000)
		dict.Guardar(numero, numero)
		abb.Guardar(numero, numero)
	}
	for iter_abb := abb.Iterador(); iter_abb.HaySiguiente(); iter_abb.Siguiente() {
		clave, dato := iter_abb.VerActual()
		require.True(t, dict.Pertenece(clave))
		require.EqualValues(t, dict.Obtener(clave), dato)
		require.True(t, clave > n && dato > n)
		n = clave
	}

}

func TestVolumenIteradorExternoRango(t *testing.T) {
	t.Log("Prueba iterar muchos elementos con el Iterador Externo un rango acotado y se haga en inorder, y que se mantenga en el rango")
	abb := TDADiccionario.CrearABB[int, int](cmp)
	dict := TDADiccionario.CrearHash[int, int]()
	n := -1 //variable para saber si se esta iterando en orden segun la funcion de comparacion
	for i := 0; i < 100000; i++ {
		numero := rand.Intn(100000)
		dict.Guardar(numero, numero)
		abb.Guardar(numero, numero)
	}
	desde := 100
	hasta := 20000
	for iter_abb := abb.IteradorRango(&desde, &hasta); iter_abb.HaySiguiente(); iter_abb.Siguiente() {
		clave, dato := iter_abb.VerActual()
		require.True(t, dict.Pertenece(clave))
		require.EqualValues(t, dict.Obtener(clave), dato)
		require.True(t, clave > n && dato > n)
		require.True(t, clave >= desde || clave <= hasta)
		n = clave
	}

}
func TestIterExternoCasoBordeDesde(t *testing.T) {
	t.Log("Prueba que en un preorder [10,5,7,9] al iterar el primer elemento sea 9 y posteriormente el 10")
	abb := TDADiccionario.CrearABB[int, int](cmp)
	abb.Guardar(10, 10)
	abb.Guardar(5, 5)
	abb.Guardar(7, 7)
	abb.Guardar(9, 9)
	orden := []int{9, 10}
	desde := 8
	iter := abb.IteradorRango(&desde, nil)
	for i := 0; i < len(orden); i++ {
		require.True(t, iter.HaySiguiente())
		clave, _ := iter.VerActual()
		require.EqualValues(t, orden[i], clave)
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())

}
func TestIterExternoCasoBordeHasta(t *testing.T) {
	t.Log("Prueba que en un preorder [5,12,11,9] al iterar el primer elemento sea 5 y posteriormente el 9")
	abb := TDADiccionario.CrearABB[int, int](cmp)
	abb.Guardar(5, 5)
	abb.Guardar(12, 12)
	abb.Guardar(11, 11)
	abb.Guardar(9, 9)
	orden := []int{5, 9}
	hasta := 10
	iter := abb.IteradorRango(nil, &hasta)
	for i := 0; i < len(orden); i++ {
		require.True(t, iter.HaySiguiente())
		clave, _ := iter.VerActual()
		require.EqualValues(t, orden[i], clave)
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())

}
