"""Algoritmo de Tabela Hash"""

import math
import time


def default_hash_function(key, size):
    """Gera um índice numérico para a chave usando a operação de módulo."""
    return hash(key) % size


def hash_sum_chars(key, size):
    """Gera um índice somando o valor de cada caractere na string."""
    return sum(ord(char) for char in str(key)) % size


def hash_multiplication(key, size):
    """Gera um índice utilizando o método da multiplicação com a constante A."""
    a = (math.sqrt(5) - 1) / 2
    val = sum(ord(char) for char in str(key))
    return math.floor(size * ((val * a) % 1))


class HashTable:
    """Implementação de uma tabela hash usando encadeamento para tratar colisões."""

    def __init__(self, size, hash_func=default_hash_function):
        """Inicializa a tabela hash com o tamanho e a função de hash especificados."""
        self.size = size
        self.hash_func = hash_func
        self.table = [[] for _ in range(size)]

    def insert(self, key, value):
        """Insere ou atualiza um par chave-valor na tabela hash."""
        index = self.hash_func(key, self.size)

        for item in self.table[index]:
            if item[0] == key:
                item[1] = value
                return

        self.table[index].append([key, value])

    def search(self, key):
        """Busca o valor associado a uma chave na tabela hash."""
        index = self.hash_func(key, self.size)

        for item in self.table[index]:
            if item[0] == key:
                return item[1]

        return None

    def delete(self, key):
        """Remove um par chave-valor da tabela hash."""
        index = self.hash_func(key, self.size)

        for i, item in enumerate(self.table[index]):
            if item[0] == key:
                self.table[index].pop(i)
                return True

        return False


if __name__ == "__main__":
    my_table = HashTable(5)

    my_table.insert("name", "Alice")
    my_table.insert("age", 25)
    my_table.insert("city", "New York")
    print("Estado inicial da tabela:", my_table.table)

    print("Busca 'name':", my_table.search("name"))
    print("Busca 'age':", my_table.search("age"))

    my_table.insert("age", 26)
    print("Busca 'age' atualizada:", my_table.search("age"))

    print("Deletou 'city'?:", my_table.delete("city"))
    print("Busca 'city' pós-remoção:", my_table.search("city"))
    print("Estado final da tabela:", my_table.table)
    print("\n" + "=" * 50 + "\n")

    DATASET = [f"user_id_{i}_{i*7}" for i in range(10000)]
    TABLE_SIZE = 2000

    FUNCTIONS = [
        ("Nativa / Default", default_hash_function),
        ("Soma de Caracteres", hash_sum_chars),
        ("Método Multiplicação", hash_multiplication),
    ]

    print(f"--- BENCHMARK: 10.000 inserções em tabelas de tamanho {TABLE_SIZE} ---")

    for name, func in FUNCTIONS:
        ht = HashTable(TABLE_SIZE, hash_func=func)

        start_time = time.perf_counter()
        for data_key in DATASET:
            ht.insert(data_key, "data")
        end_time = time.perf_counter()

        elapsed_time = (end_time - start_time) * 1000

        total_collisions = 0
        for bucket in ht.table:
            if len(bucket) > 1:
                total_collisions += len(bucket) - 1

        lengths = [len(bucket) for bucket in ht.table]
        mean = sum(lengths) / TABLE_SIZE
        variance = sum((x - mean) ** 2 for x in lengths) / TABLE_SIZE
        std_dev = math.sqrt(variance)

        print(f"\n[{name}]")
        print(f"  Tempo de montagem: {elapsed_time:.2f} ms")
        print(f"  Total de colisões reais: {total_collisions}")
        print(f"  Desvio Padrão (Distribuição): {std_dev:.2f}")
x