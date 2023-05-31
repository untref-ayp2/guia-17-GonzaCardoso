package ejercicios

//Reescribir el algortimo de BucketSort, sabiendo que los números que deberá ordenar
//se encuentran entre el 0 y el 9 inclusive, por lo tanto en lugar de guardar los números
// los cuenta y después genera el arreglo final con los elementos ordenados.

func Ej1(arr []int) {
		if len(arr) <= 1 {
		return
	}
	buckets := make([]int, 10, 10)

	// Colocar los elementos en los baldes correspondientes
	for _, num := range arr {
		buckets[num]++
	}

	// Combinar los buckets en el arreglo original
	pos := 0
	for value, i := range buckets {
		for j := 0; j < i; j++ {
			arr[pos] = value
			pos++
		}
	}
}

// Escribir tests unitarios para la función Ej1
