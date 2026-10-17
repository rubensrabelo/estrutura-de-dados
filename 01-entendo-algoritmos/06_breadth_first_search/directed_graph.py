"""Algoritmo de Grafo Direcionado"""


class DirectedGraph:
    """Implementação completa de um grafo direcionado usando lista de adjacência."""

    def __init__(self):
        """Inicializa um grafo direcionado vazio."""
        self.adjacency_list = {}

    def add_vertex(self, vertex):
        """Adiciona um novo vértice isolado ao grafo caso ele não exista."""
        if vertex not in self.adjacency_list:
            self.adjacency_list[vertex] = []

    def add_edge(self, source, destination):
        """Adiciona uma aresta unidirecional ligando a origem ao destino."""
        self.add_vertex(source)
        self.add_vertex(destination)
        if destination not in self.adjacency_list[source]:
            self.adjacency_list[source].append(destination)

    def remove_edge(self, source, destination):
        """Remove a aresta unidirecional existente entre a origem e o destino."""
        if source in self.adjacency_list and destination in self.adjacency_list[source]:
            self.adjacency_list[source].remove(destination)

    def remove_vertex(self, vertex):
        """Remove um vértice do grafo e todas as arestas direcionadas associadas a ele."""
        if vertex in self.adjacency_list:
            self.adjacency_list.pop(vertex)
            for _, neighbors in self.adjacency_list.items():
                if vertex in neighbors:
                    neighbors.remove(vertex)

    def has_vertex(self, vertex):
        """Verifica se um determinado vértice existe no grafo."""
        return vertex in self.adjacency_list

    def has_edge(self, source, destination):
        """Verifica se existe uma aresta direcionada da origem para o destino."""
        return source in self.adjacency_list and destination in self.adjacency_list[source]

    def get_successors(self, vertex):
        """Retorna a lista de sucessores (vizinhos de saída) de um vértice."""
        return self.adjacency_list.get(vertex, None)

    def in_degree(self, vertex):
        """Calcula o grau de entrada (número de arestas que chegam) de um vértice."""
        if vertex not in self.adjacency_list:
            return None
        return sum(1 for source, neighbors in self.adjacency_list.items() if vertex in neighbors)

    def out_degree(self, vertex):
        """Calcula o grau de saída (número de arestas que partem) de um vértice."""
        if vertex not in self.adjacency_list:
            return None
        return len(self.adjacency_list[vertex])

    def __str__(self):
        """Retorna a representação textual da lista de adjacência do grafo."""
        lines = [f"{vertex} -> {neighbors}" for vertex, neighbors in self.adjacency_list.items()]
        return "\n".join(lines)


if __name__ == "__main__":
    g = DirectedGraph()
    g.add_edge("A", "B")
    g.add_edge("A", "C")
    g.add_edge("B", "C")

    print("--- Grafo Inicial ---")
    print(g)

    print("\nGrau de Entrada de C:", g.in_degree("C"))
    print("Grau de Saída de A:", g.out_degree("A"))
    print("Existe aresta A->B?:", g.has_edge("A", "B"))

    g.remove_edge("A", "B")
    print("\n--- Após Remover Aresta A->B ---")
    print(g)

    g.remove_vertex("C")
    print("\n--- Após Remover Vértice C ---")
    print(g)
