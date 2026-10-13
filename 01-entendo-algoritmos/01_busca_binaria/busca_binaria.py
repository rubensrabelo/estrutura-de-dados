"""
Código referente a pesquisa binária
"""


def binary_search(arr, item):
    """Função interativa da pesquisa binária"""
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


if __name__ == "__main__":
    my_array = [1, 3, 5, 7, 9]
    print(binary_search(my_array, 3))
    print(binary_search(my_array, -1))
