package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	// --- IMPORTANTE: Ajuste estes caminhos para seu módulo Go ---
	"example.com/data"
	"example.com/sort"
	// -----------------------------------------------------------
)

// bToMb converte bytes para Megabytes.
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// printMemStats exibe as estatísticas de memória.
func printMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s:\n", label)
	fmt.Printf("\tHeap Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotal Alloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys Mem = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

// SorterFunc é um tipo para representar qualquer uma das suas funções de sort
// adaptadas para receber o array e as métricas.
type SorterFunc func([]int, *quicksort.Metricas)

// runTest executa um teste para um dado algoritmo e array base.
func runTest(testName string, sorter SorterFunc, baseArray []int) {
	fmt.Printf("\n--- Testando %s ---\n", testName)

	// Cria uma cópia para não alterar o array base
	arr := make([]int, len(baseArray))
	copy(arr, baseArray)

	// Inicializa as métricas
	metrics := quicksort.Metricas{}

	// Medição de Memória (Antes)
	runtime.GC()
	printMemStats("Memória Antes")

	// Medição de Tempo (Início) e Execução
	startTime := time.Now()
	sorter(arr, &metrics) // Chama a função de ordenação adaptada
	duration := time.Since(startTime)

	// Medição de Memória (Depois)
	runtime.GC()
	printMemStats("Memória Depois")

	// Exibe Resultados
	fmt.Printf("Tempo: %s | Comparações: %d | Trocas: %d\n", duration, metrics.Comparisons, metrics.Swaps)

	// Verificação (Opcional)
	isSorted := true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			isSorted = false
			fmt.Println("*** ATENÇÃO: Array NÃO está ordenado corretamente! ***")
			break
		}
	}
	fmt.Println("Array está ordenado?", isSorted)
	fmt.Println("------------------------------")
}

func main() {
	// Semeia o gerador de números aleatórios (importante se GenerateData usa rand)
	rand.Seed(time.Now().UnixNano())

	// --- Adaptadores para Funções de Sort ---
	mSortAdapter := func(arr []int, m *quicksort.Metricas) {
		quicksort.Quicksort(arr, 0, len(arr)-1, m)
	}

	hybridAdapter := func(arr []int, m *quicksort.Metrics) {
		_ = quicksort.Qsort_HybridIterative(arr, m)
	}

	nonInPlaceAdapter := func(arr []int, m *quicksort.Metrics) {
		sorted := quicksort.Qsort_NonInPlace(arr, m)
		copy(arr, sorted)
	}

	// Agrupa os algoritmos
	sortersToTest := map[string]SorterFunc{
		"Mquicksort":             mSortAdapter,
		"HybridIterativeSort":    hybridAdapter,
		"Quicksort (NonInPlace)": nonInPlaceAdapter,
	}

	// --- Execução dos Testes (3 Vezes, Cada Vez com Novo Dataset Decrescente) ---
	fmt.Println("\n>>> INICIANDO TESTES COM 3 DATASETS DECRESCENTES DIFERENTES <<<")

	for i := 1; i <= 3; i++ {
		fmt.Printf("\n================= DATASET %d/3 =================\n", i)
		// Chama GenerateData() A CADA ITERAÇÃO para obter um novo dataset
		currentDescArray := data.GenerateData()
		fmt.Printf("Gerado Dataset %d (Tamanho: %d)\n", i, len(currentDescArray))

		// Roda cada algoritmo com este novo dataset decrescente
		for name, sorter := range sortersToTest {
			runTest(fmt.Sprintf("%s (Dataset %d)", name, i), sorter, currentDescArray)
		}
	}

	fmt.Println("\n--- Testes Concluídos ---")
}