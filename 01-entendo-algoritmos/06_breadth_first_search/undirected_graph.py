"""Algoritmo de Grafo Não Direcionado"""


class UndirectedGraph:
    """Implementação completa de um grafo não direcionado usando lista de adjacência."""

    def __init__(self):
        """Inicializa um grafo não direcionado vazio."""
        self.adjacency_list = {}

    def add_vertex(self, vertex):
        """Adiciona um novo vértice isolado ao grafo caso ele não exista."""
        if vertex not in self.adjacency_list:
            self.adjacency_list[vertex] = []

    def add_edge(self, vertex1, vertex2):
        """Adiciona uma aresta bidirecional conectando os dois vértices entre si."""
        self.add_vertex(vertex1)
        self.add_vertex(vertex2)
        if vertex2 not in self.adjacency_list[vertex1]:
            self.adjacency_list[vertex1].append(vertex2)
        if vertex1 not in self.adjacency_list[vertex2]:
            self.adjacency_list[vertex2].append(vertex1)

    def remove_edge(self, vertex1, vertex2):
        """Remove a aresta bidirecional existente entre os dois vértices."""
        if vertex1 in self.adjacency_list and vertex2 in self.adjacency_list[vertex1]:
            self.adjacency_list[vertex1].remove(vertex2)
        if vertex2 in self.adjacency_list and vertex1 in self.adjacency_list[vertex2]:
            self.adjacency_list[vertex2].remove(vertex1)

    def remove_vertex(self, vertex):
        """Remove um vértice do grafo e desconecta todas as suas arestas bidirecionais."""
        if vertex in self.adjacency_list:
            for neighbor in list(self.adjacency_list[vertex]):
                self.adjacency_list[neighbor].remove(vertex)
            self.adjacency_list.pop(vertex)

    def has_vertex(self, vertex):
        """Verifica se um determinado vértice existe no grafo."""
        return vertex in self.adjacency_list

    def has_edge(self, vertex1, vertex2):
        """Verifica se existe uma aresta bidirecional entre os dois vértices."""
        return vertex1 in self.adjacency_list and vertex2 in self.adjacency_list[vertex1]

    def get_neighbors(self, vertex):
        """Retorna a lista de todos os vizinhos adjacentes de um vértice."""
        return self.adjacency_list.get(vertex, None)

    def degree(self, vertex):
        """Calcula o grau total (número de conexões) de um vértice específico."""
        if vertex not in self.adjacency_list:
            return None
        return len(self.adjacency_list[vertex])

    def __str__(self):
        """Retorna a representação textual da lista de adjacência do grafo."""
        lines = [f"{vertex} <-> {neighbors}" for vertex, neighbors in self.adjacency_list.items()]
        return "\n".join(lines)


if __name__ == "__main__":
    g = UndirectedGraph()
    g.add_edge("A", "B")
    g.add_edge("A", "C")
    g.add_edge("B", "C")

    print("--- Grafo Inicial ---")
    print(g)

    print("\nGrau do vértice A:", g.degree("A"))
    print("Existe conexão B-C?:", g.has_edge("B", "C"))
    print("Vizinhos de C:", g.get_neighbors("C"))

    g.remove_edge("A", "B")
    print("\n--- Após Remover Aresta A-B ---")
    print(g)

    g.remove_vertex("C")
    print("\n--- Após Remover Vértice C ---")
    print(g)
