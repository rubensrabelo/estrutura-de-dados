"""
Módulo para o algoritmo Selection Sort que retorna uma nova estrutura.
"""

def selection_sort(arr):
    """
    Ordena uma estrutura de dados utilizando o algoritmo Selection Sort
    e retorna uma nova cópia ordenada.
    """
    new_arr = list(arr)
    n = len(new_arr)
    for i in range(n):
        min_index = i
        for j in range(i + 1, n):
            if new_arr[j] < new_arr[min_index]:
                min_index = j
        if min_index != i:
            new_arr[i], new_arr[min_index] = new_arr[min_index], new_arr[i]
    return new_arr
