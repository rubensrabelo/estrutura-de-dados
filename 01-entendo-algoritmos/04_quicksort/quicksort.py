"""Algoritmo do Quicksort"""


def partition(arr, pivot):
    """Divide a lista em três sublistas baseadas no pivô central."""
    left = [x for x in arr if x < pivot]
    right = [x for x in arr if x > pivot]
    return left, right


def quicksort(arr):
    """Ordena uma lista de elementos em ordem crescente de forma recursiva."""
    if len(arr) <= 1:
        return arr

    pivot = arr[len(arr) // 2]
    left, right = partition(arr, pivot)

    return quicksort(left) + [pivot] + quicksort(right)


if __name__ == "__main__":
    standard_list = [3, 6, 8, 10, 1, 2, 1, -5, 0]
    print(quicksort(standard_list))

    sorted_list = [1, 2, 3, 4, 5]
    print(quicksort(sorted_list))

    reversed_list = [5, 4, 3, 2, 1]
    print(quicksort(reversed_list))

    empty_list = []
    print(quicksort(empty_list))

    single_element_list = [7]
    print(quicksort(single_element_list))

    repeated_elements_list = [4, 4, 2, 2, 8, 4, 1]
    print(quicksort(repeated_elements_list))
