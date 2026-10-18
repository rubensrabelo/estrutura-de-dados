"""
Módulo contendo a implementação do algoritmo KNN encapsulado em uma classe estruturada.
"""

import math

class KNeighborsClassifier:
    """Classificador baseado no algoritmo dos K-Vizinhos Próximos (KNN)."""

    def __init__(self, k: int = 3):
        """Inicializa o classificador definindo o número de vizinhos próximos."""
        self.k = k
        self.training_data = []
        self.training_labels = []

    def fit(self, x_train: list, y_train: list) -> None:
        """Armazena os dados e rótulos de treino na memória da classe."""
        self.training_data = x_train
        self.training_labels = y_train

    def _calculate_distance(self, point1: list, point2: list) -> float:
        """Método interno para calcular a distância euclidiana entre dois pontos."""
        squared_sum = sum((p1 - p2) ** 2 for p1, p2 in zip(point1, point2))
        return math.sqrt(squared_sum)

    def predict(self, test_point: list) -> int:
        """Classifica um ponto de teste com base na votação dos K vizinhos mais próximos."""
        distances = []

        # Calcula a distância para todos os pontos armazenados no treino
        for i, train_point in enumerate(self.training_data):
            dist = self._calculate_distance(test_point, train_point)
            distances.append((dist, self.training_labels[i]))

        # Ordena pelas menores distâncias e filtra os K primeiros
        distances.sort(key=lambda x: x[0])
        k_nearest = distances[:self.k]

        # Contabiliza a votação majoritária das classes vizinhas
        votes = {}
        for _, label in k_nearest:
            votes[label] = votes.get(label, 0) + 1

        return max(votes, key=votes.get)


if __name__ == "__main__":
    # Exemplo prático: Classificar carros entre Econômico (0) ou Luxo (1)
    car_features = [
        [1.0, 15.2],  # Econômico
        [1.4, 13.8],  # Econômico
        [1.0, 16.0],  # Econômico
        [3.0, 8.5],   # Luxo
        [4.1, 6.2],   # Luxo
        [2.5, 9.1]    # Luxo
    ]
    car_labels = [0, 0, 0, 1, 1, 1]

    knn = KNeighborsClassifier(k=3)

    knn.fit(car_features, car_labels)

    unknown_car = [2.0, 10.5]

    prediction = knn.predict(unknown_car)

    CAR_TYPE = "Econômico" if prediction == 0 else "Luxo"
    print(
        f"Resultado: O carro com as especificações {unknown_car} foi classificado como {CAR_TYPE}."
    )
