"""Algoritmo de Dijkstra"""

import heapq
from weighted_graph import WeightedGraph


def dijkstra(graph, start_vertex):
    """Calcula os caminhos mais curtos a partir de um vértice usando fila de prioridade."""
    if not graph.has_vertex(start_vertex):
        return None

    distances = {vertex: float("inf") for vertex in graph.adjacency_list}
    distances[start_vertex] = 0

    priority_queue = [(0, start_vertex)]

    while priority_queue:
        current_distance, current_vertex = heapq.heappop(priority_queue)

        if current_distance > distances[current_vertex]:
            continue

        for neighbor, weight in graph.adjacency_list[current_vertex]:
            distance = current_distance + weight

            if distance < distances[neighbor]:
                distances[neighbor] = distance
                heapq.heappush(priority_queue, (distance, neighbor))

    return distances


if __name__ == "__main__":
    g = WeightedGraph()
    g.add_edge("A", "B", 4)
    g.add_edge("A", "C", 2)
    g.add_edge("B", "C", 5)
    g.add_edge("B", "D", 10)
    g.add_edge("C", "D", 3)
    g.add_edge("D", "E", 1)

    print("--- Executando o Algoritmo de Dijkstra ---")
    shortest_paths = dijkstra(g, "A")

    for dest, dist in shortest_paths.items():
        print(f"Menor distancia de A ate {dest}: {dist}")
