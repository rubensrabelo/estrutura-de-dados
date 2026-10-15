"""
Módulo para a estrutura de dados de Array Estático.
"""

class MyArray:
    """
    Uma classe que simula um Array de capacidade fixa.
    """

    def __init__(self, capacity: int):
        """Inicializa o array com uma capacidade fixa definida."""
        self.capacity = capacity
        self.data = [None] * capacity  # Iniciliza o array com uma quantidade estipulada de None
        self.size = 0

    def insert(self, element) -> bool:
        """
        Insere um elemento no final do array caso haja capacidade disponível.
        """
        if self.size >= self.capacity:
            print("Erro: O array está cheio!")  # Realizar o tratamento adequado
            return False
        self.data[self.size] = element
        self.size += 1
        return True

    def remove(self, element) -> bool:
        """
        Remove a primeira ocorrência do elemento e desloca os seguintes para a esquerda.
        """
        index = self.search(element)

        if index == -1:
            print(f"Elemento {element} não encontrado para remoção.")
            return False

        for j in range(index, self.size - 1):
            self.data[j] = self.data[j + 1]

        self.data[self.size - 1] = None
        self.size -= 1
        return True

    def search(self, element) -> int:
        """
        Busca por um elemento no array de forma linear.
        """
        for i in range(self.size):
            if self.data[i] == element:
                return i
        return -1

    def __str__(self) -> str:
        """Exibe a estrutura do array atual no console mostrando o preenchimento."""
        valid_elements = [self.data[i] for i in range(self.size)]
        return f"Array (Capacidade: {self.size}/{self.capacity}): {valid_elements}"
