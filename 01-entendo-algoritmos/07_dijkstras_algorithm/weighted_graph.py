"""Algoritmo de Grafo Ponderado"""


class WeightedGraph:
    """Implementação de um grafo direcionado e ponderado usando lista de adjacência."""

    def __init__(self):
        """Inicializa um grafo ponderado vazio."""
        self.adjacency_list = {}

    def add_vertex(self, vertex):
        """Adiciona um novo vértice isolado ao grafo caso ele não exista."""
        if vertex not in self.adjacency_list:
            self.adjacency_list[vertex] = []

    def add_edge(self, source, destination, weight):
        """Adiciona uma aresta direcionada com peso associado entre os vértices."""
        self.add_vertex(source)
        self.add_vertex(destination)
        self.adjacency_list[source].append((destination, weight))

    def has_vertex(self, vertex):
        """Verifica se um determinado vértice existe no grafo."""
        return vertex in self.adjacency_list

    def __str__(self):
        """Retorna a representação textual da lista de adjacência com os pesos."""
        lines = [f"{vertex} -> {neighbors}" for vertex, neighbors in self.adjacency_list.items()]
        return "\n".join(lines)
