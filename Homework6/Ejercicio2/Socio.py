from Biblioteca import Biblioteca
from Libro import Libro


class Socio:
    def __init__(self, numero_socio, nombre, direccion):
        self.numero_socio = numero_socio
        self.nombre = nombre
        self.direccion = direccion
        self.libros_prestados = []

    def __str__(self):
        return f"Socio: {self.nombre} (Numero de Socio: {self.numero_socio})"

    def __repr__(self):
        return self.__str__()


