"""
Módulo para comparar a performance e consumo de hardware usando abordagens de programação dinâmica.
"""

import time
import tracemalloc

def fibonacci_recursive(n: int) -> int:
    """Calcula o enésimo número de Fibonacci usando recursão simples padrão."""
    if n <= 0:
        return 0
    if n == 1:
        return 1
    return fibonacci_recursive(n - 1) + fibonacci_recursive(n - 2)

def fibonacci_memorization(n: int, memo: dict = None) -> int:
    """Calcula o enésimo número de Fibonacci usando a abordagem dinâmica Top-Down (Memorization)."""
    if memo is None:
        memo = {}
    if n <= 0:
        return 0
    if n == 1:
        return 1
    if n in memo:
        return memo[n]
    memo[n] = fibonacci_memorization(n - 1, memo) + fibonacci_memorization(n - 2, memo)
    return memo[n]

def fibonacci_tabulation(n: int) -> int:
    """Calcula o enésimo número de Fibonacci usando a abordagem dinâmica Bottom-Up (Tabulation)."""
    if n <= 0:
        return 0
    if n == 1:
        return 1

    fib_table = [0] * (n + 1)
    fib_table[1] = 1
    for i in range(2, n + 1):
        fib_table[i] = fib_table[i - 1] + fib_table[i - 2]
    return fib_table[n]

def run_performance_test():
    """Executa e compara o tempo de execução e o consumo de memória RAM das abordagens."""
    # N=40 é seguro para todos os métodos e roda rápido
    number = 40

    print(f"--- Iniciando Teste de Performance e Consumo do PC [N={number}] ---\n")

    # 1. Medição da Recursão Normal
    tracemalloc.start()
    start_time = time.perf_counter()
    res_rec = fibonacci_recursive(number)
    end_time = time.perf_counter()
    _, peak_rec = tracemalloc.get_traced_memory()
    tracemalloc.stop()

    print(" Normal (Recursão Simples):")
    print(f"    - Resultado: {res_rec}")
    print(f"    - Tempo de Execução: {end_time - start_time:.6f} segundos")
    print(f"    - Pico de Memória RAM: {peak_rec / 1024:.2f} KB")

    # 2. Medição de Memorization
    tracemalloc.start()
    start_time = time.perf_counter()
    res_memo = fibonacci_memorization(number)
    end_time = time.perf_counter()
    _, peak_memo = tracemalloc.get_traced_memory()
    tracemalloc.stop()

    print("\n DP Top-Down (Memorization):")
    print(f"    - Resultado: {res_memo}")
    print(f"    - Tempo de Execução: {end_time - start_time:.6f} segundos")
    print(f"    - Pico de Memória RAM: {peak_memo / 1024:.2f} KB")

    # 3. Medição de Tabulation
    tracemalloc.start()
    start_time = time.perf_counter()
    res_tab = fibonacci_tabulation(number)
    end_time = time.perf_counter()
    _, peak_tab = tracemalloc.get_traced_memory()
    tracemalloc.stop()

    print("\n DP Bottom-Up (Tabulation):")
    print(f"    - Resultado: {res_tab}")
    print(f"    - Tempo de Execução: {end_time - start_time:.6f} segundos")
    print(f"    - Pico de Memória RAM: {peak_tab / 1024:.2f} KB")

if __name__ == "__main__":
    run_performance_test()
