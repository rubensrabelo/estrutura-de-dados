"""Sistema de Votação Eletrônica com Métricas de Infraestrutura"""

import time
from hash_table import HashTable


def register_voters(hash_table):
    """Cadastra eleitores com IDs de votação únicos na tabela hash medindo o tempo."""
    start_time = time.perf_counter()
    for i in range(1, 1001):
        voter_id = f"VOTE-{i:04d}"
        voter_name = f"Eleitor Numero {i}"
        hash_table.insert(voter_id, [voter_name, False])
    end_time = time.perf_counter()
    return (end_time - start_time) * 1000


def cast_vote(hash_table, voter_id):
    """Verifica as credenciais do eleitor e computa o voto se permitido."""
    voter_data = hash_table.search(voter_id)

    if voter_data is None:
        print(f" Alerta: O ID {voter_id} nao esta cadastrado no sistema.")
        return

    name, has_voted = voter_data

    if has_voted:
        print(f" Negado: {name} (ID: {voter_id}) ja realizou o voto anteriormente!")
    else:
        voter_data = True
        hash_table.insert(voter_id, voter_data)
        print(f" Sucesso: Voto confirmado para {name} (ID: {voter_id}).")


def display_infrastructure_metrics(hash_table, build_time):
    """Calcula e exibe as métricas de performance da tabela hash de votação."""
    total_collisions = 0
    total_registered = 0

    for bucket in hash_table.table:
        if len(bucket) > 1:
            total_collisions += len(bucket) - 1
        for _ in bucket:
            total_registered += 1

    search_start = time.perf_counter()
    for i in range(1, 1001):
        hash_table.search(f"VOTE-{i:04d}")
    search_end = time.perf_counter()
    total_search_time = (search_end - search_start) * 1000
    avg_search_time = total_search_time / 1000

    print("\n" + "=" * 45)
    print("    MÉTRICAS DE INFRAESTRUTURA DA TABELA    ")
    print("=" * 45)
    print(f"  Total de eleitores cadastrados : {total_registered}")
    print(f"  Tamanho final da tabela hash   : {hash_table.size}")
    print(f"  Tempo de construção/cadastro   : {build_time:.2f} ms")
    print(f"  Total de colisões reais        : {total_collisions}")
    print(f"  Tempo total de pesquisa (1000) : {total_search_time:.4f} ms")
    print(f"  Tempo médio por pesquisa       : {avg_search_time:.4f} ms")
    print("=" * 45)


if __name__ == "__main__":
    # Inicializa com tamanho 50 para forçar o crescimento/remontagem dinâmico até 1000
    voting_booth = HashTable(size=50)
    
    print("--- INICIANDO CADASTRO E SISTEMA ---")
    time_taken_to_build = register_voters(voting_booth)

    print("\n--- SIMULANDO VOTOS DE EXEMPLO ---")
    cast_vote(voting_booth, "VOTE-0025")
    cast_vote(voting_booth, "VOTE-0025")
    cast_vote(voting_booth, "VOTE-0999")
    cast_vote(voting_booth, "VOTE-1500")

    display_infrastructure_metrics(voting_booth, time_taken_to_build)
