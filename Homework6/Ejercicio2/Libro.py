class Libro:
    def __init__(self, codigo, titulo, autor):
        self.codigo = codigo
        self.titulo = titulo
        self.autor = autor
        self.disponible = True

    def __str__(self):
        return f"{self.titulo} por {self.autor} (Codigo: {self.codigo}, Disponible: {self.disponible})"
