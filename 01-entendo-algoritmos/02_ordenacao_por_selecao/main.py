"""
Módulo de execução principal para testar todas as estruturas de dados
e a ordenação com medição de tempo.
"""

import time
from array_structure import MyArray
from linked_list import LinkedList
from selection_sort import selection_sort


def main():
    """
    Classe que roda o Vetor e a Listagem medindo o tempo de execução.
    """
    # Gerando 100 dados para teste em ordem reversa
    dados = list(range(100, 0, -1))

    print("--- DESEMPENHO ARRAY ESTÁTICO (MyArray) ---")
    arr = MyArray(100)

    # Inserção
    inicio = time.perf_counter_ns()
    for item in dados:
        arr.insert(item)
    fim = time.perf_counter_ns()
    print(f"Tempo de inserção (100 elementos): {fim - inicio} ns")

    # Busca
    inicio = time.perf_counter_ns()
    arr.search(1)
    fim = time.perf_counter_ns()
    print(f"Tempo de busca (pior cenário - elemento 1): {fim - inicio} ns")

    # Ordenação
    arr_elements = [arr.data[i] for i in range(arr.size)]
    inicio = time.perf_counter_ns()
    selection_sort(arr_elements)
    fim = time.perf_counter_ns()
    print(f"Tempo de ordenação (Selection Sort): {fim - inicio} ns")

    # Remoção
    inicio = time.perf_counter_ns()
    arr.remove(100)
    fim = time.perf_counter_ns()
    print(f"Tempo de remoção (primeiro elemento): {fim - inicio} ns\n")


    print("--- DESEMPENHO LISTA ENCADEADA (LinkedList) ---")
    ll = LinkedList()

    # Inserção
    inicio = time.perf_counter_ns()
    for item in dados:
        ll.insert(item)
    fim = time.perf_counter_ns()
    print(f"Tempo de inserção (100 elementos): {fim - inicio} ns")

    # Busca
    inicio = time.perf_counter_ns()
    ll.search(1)
    fim = time.perf_counter_ns()
    print(f"Tempo de busca (pior cenário - elemento 1): {fim - inicio} ns")

    # Ordenação
    ll_elements = []
    current = ll.head
    while current is not None:
        ll_elements.append(current.data)
        current = current.next

    inicio = time.perf_counter_ns()
    selection_sort(ll_elements)
    fim = time.perf_counter_ns()
    print(f"Tempo de ordenação (Selection Sort): {fim - inicio} ns")

    # Remoção
    inicio = time.perf_counter_ns()
    ll.remove(100)
    fim = time.perf_counter_ns()
    print(f"Tempo de remoção (primeiro elemento): {fim - inicio} ns")


if __name__ == "__main__":
    main()
