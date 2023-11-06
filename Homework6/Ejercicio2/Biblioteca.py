class Biblioteca:
    def __init__(self):
        self.socios = []
        self.libros = []
        self.prestamos = []

    def agregar_socio(self, socio):
        self.socios.append(socio)

    def agregar_libro(self, libro):
        self.libros.append(libro)

    def prestar_libro(self, libro, socio, fecha):
        if libro.disponible:
            libro.disponible = False
            prestamo = Prestamo(libro, socio, fecha)
            self.prestamos.append(prestamo)
            socio.libros_prestados.append(prestamo)

    def buscar_libros_por_autor(self, autor):
        return [libro for libro in self.libros if libro.autor == autor]

    def obtener_libros_vencidos(self):
        # Implementa la lógica para encontrar libros vencidos
        pass

    def obtener_socios_con_mas_de_3_libros(self):
        return list(filter(lambda socio: len(socio.libros_prestados) > 3, self.socios))


class Prestamo:
    def __init__(self, libro, socio, fecha):
        self.libro = libro
        self.socio = socio
        self.fecha = fecha


biblioteca = Biblioteca()
