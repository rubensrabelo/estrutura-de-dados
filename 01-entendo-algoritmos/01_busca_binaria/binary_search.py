"""Módulo para o algoritmo de busca binária."""

def binary_search(arr: list | range, item: any) -> int | None:
    """Busca um item em uma lista ordenada e retorna seu índice."""
    low = 0
    high = len(arr) - 1

    while low <= high:
        middle = (low + high) // 2
        guess = arr[middle]

        if guess == item:
            return middle
        if guess > item:
            high = middle - 1
        else:
            low = middle + 1

    return None
