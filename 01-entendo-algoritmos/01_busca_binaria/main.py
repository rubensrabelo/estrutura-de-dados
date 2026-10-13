"""Módulo principal para execução e comparação dos testes de busca."""

import time
from binary_search import binary_search
from simple_search import simple_search

def format_time(ns: int) -> str:
    """Converte o tempo em nanossegundos para a unidade mais adequada."""
    if ns < 1000:
        return f"{ns} ns"
    elif ns < 1000000:
        return f"{ns / 1000:.2f} µs"
    elif ns < 1000000000:
        return f"{ns / 1000000:.2f} ms"
    else:
        return f"{ns / 1000000000:.2f} s"

def run_test(size: int) -> None:
    """Executa e compara os testes de busca binária e simples no pior caso."""
    data_list = range(size)
    worst_case = -1

    # Teste da Busca Binária
    start_bin = time.perf_counter_ns()
    binary_search(data_list, worst_case)
    end_bin = time.perf_counter_ns()
    time_bin = format_time(end_bin - start_bin)

    # Teste da Busca Simples
    start_sim = time.perf_counter_ns()
    simple_search(data_list, worst_case)
    end_sim = time.perf_counter_ns()
    time_sim = format_time(end_sim - start_sim)

    print(f"Tamanho: {size:,} | Binária: {time_bin} | Simples: {time_sim}")

if __name__ == '__main__':
    run_test(10)
    run_test(100)
    run_test(1000)
    run_test(10000)
    run_test(100000)
    run_test(1000000)
