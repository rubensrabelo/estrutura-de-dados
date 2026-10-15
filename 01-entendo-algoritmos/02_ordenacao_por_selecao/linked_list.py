"""
Módulo para a estrutura de dados de lista encadeada simples.
"""

class Node:
    """
    Representa um nó individual na lista encadeada.
    """

    def __init__(self, data):
        """Inicializa o nó com os dados fornecidos e o próximo nó como None."""
        self.data = data
        self.next = None


class LinkedList:
    """
    Uma classe que simula uma Lista Encadeada Simples.
    """

    def __init__(self):
        """Inicializa uma lista encadeada vazia."""
        self.head = None
        self.size = 0

    def insert(self, element) -> bool:
        """
        Insere um elemento no final da lista encadeada.
        """
        new_node = Node(element)

        if self.head is None:
            self.head = new_node
        else:
            current = self.head
            while current.next is not None:
                current = current.next
            current.next = new_node

        self.size += 1
        return True

    def remove(self, element) -> bool:
        """
        Remove a primeira ocorrência do elemento utilizando o método search para busca.
        """
        if self.head is None:
            print("Erro: A lista está vazia!")
            return False

        # Caso o elemento esteja logo na cabeça (head) da lista
        if self.head.data == element:
            self.head = self.head.next
            self.size -= 1
            return True

        # Busca o nó anterior ao elemento que queremos remover
        previous_node = self._search_previous(element)

        if previous_node is None:
            print(f"Elemento {element} não encontrado para remoção.")
            return False

        # Remove o elemento pulando o nó que o contém
        previous_node.next = previous_node.next.next
        self.size -= 1
        return True

    def search(self, element) -> int:
        """
        Busca por um elemento na lista de forma linear.
        Retorna o índice (posição) se encontrado, ou -1 caso contrário.
        """
        current = self.head
        index = 0

        while current is not None:
            if current.data == element:
                return index
            current = current.next
            index += 1

        return -1

    def _search_previous(self, element) -> Node:
        """
        Método auxiliar para buscar o nó ANTERIOR à primeira ocorrência do elemento.
        Retorna o nó anterior se encontrado, ou None caso contrário.
        """
        current = self.head
        while current.next is not None:
            if current.next.data == element:
                return current
            current = current.next
        return None

    def __str__(self) -> str:
        """Exibe a estrutura da lista atual no console mostrando o encadeamento."""
        elements = []
        current = self.head

        while current is not None:
            elements.append(str(current.data))
            current = current.next

        list_structure = " -> ".join(elements) if elements else "Vazia"
        return f"LinkedList (Tamanho: {self.size}): {list_structure}"
