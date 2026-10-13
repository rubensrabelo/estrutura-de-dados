"""Módulo para o algoritmo de busca simples."""

def simple_search(arr: list | range, item: any) -> int | None:
    """Busca um item elemento por elemento na lista e retorna seu índice."""
    for i in range(len(arr)):
        if arr[i] == item:
            return i
    return None
