"""Algoritmo de Busca em Largura"""

from collections import deque
from directed_graph import DirectedGraph
from undirected_graph import UndirectedGraph


def bfs(graph, start_vertex):
    """
    Realiza uma busca em largura a partir de um vértice inicial retornando o caminho.
    """
    if not graph.has_vertex(start_vertex):
        return []

    visited = set([start_vertex])
    queue = deque([start_vertex])
    order = []

    while queue:
        vertex = queue.popleft()
        order.append(vertex)

        # Abstração geral para capturar vizinhos de ambos os tipos de grafos
        if hasattr(graph, "get_successors"):
            neighbors = graph.get_successors(vertex)
        else:
            neighbors = graph.get_neighbors(vertex)

        if neighbors:
            for neighbor in neighbors:
                if neighbor not in visited:
                    visited.add(neighbor)
                    queue.append(neighbor)

    return order


if __name__ == "__main__":
    print("--- TESTANDO BFS NO GRAFO DIRECIONADO ---")
    d_graph = DirectedGraph()
    d_graph.add_edge("A", "B")
    d_graph.add_edge("A", "C")
    d_graph.add_edge("B", "D")
    d_graph.add_edge("C", "E")
    d_graph.add_edge("D", "F")
    d_graph.add_edge("E", "F")

    print("Estrutura do Grafo:")
    print(d_graph)
    print("Ordem do BFS partindo de A:", bfs(d_graph, "A"))

    print("\n--- TESTANDO BFS NO GRAFO NÃO DIRECIONADO ---")
    u_graph = UndirectedGraph()
    u_graph.add_edge("A", "B")
    u_graph.add_edge("A", "C")
    u_graph.add_edge("B", "D")
    u_graph.add_edge("C", "E")
    u_graph.add_edge("D", "F")
    u_graph.add_edge("E", "F")

    print("Estrutura do Grafo:")
    print(u_graph)
    print("Ordem do BFS partindo de A:", bfs(u_graph, "A"))
