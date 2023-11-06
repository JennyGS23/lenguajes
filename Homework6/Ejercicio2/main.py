from Biblioteca import Biblioteca
from Libro import Libro
from Socio import Socio

# Crear una instancia de la biblioteca
biblioteca = Biblioteca()

# Crear libros
libro1 = Libro("L001", "Libro 1", "Autor 1")
libro2 = Libro("L002", "Libro 2", "Autor 2")
libro3 = Libro("L003", "Libro 3", "Autor 1")
libro4 = Libro("L004", "Libro 4", "Autor 3")
libro5 = Libro("L005", "Libro 5", "Autor 4")
libro6 = Libro("L006", "Libro 6", "Autor 5")

# Crear socios
socio1 = Socio(1, "Juan Perez", "123 Calle Principal")
socio2 = Socio(2, "Maria Gonzalez", "456 Avenida Secundaria")
socio3 = Socio(3, "Carlos Lopez", "789 Calle Peatonal")

# Agregar libros y socios a la biblioteca
biblioteca.agregar_libro(libro1)
biblioteca.agregar_libro(libro2)
biblioteca.agregar_libro(libro3)
biblioteca.agregar_libro(libro4)
biblioteca.agregar_libro(libro5)
biblioteca.agregar_libro(libro6)
biblioteca.agregar_socio(socio1)
biblioteca.agregar_socio(socio2)
biblioteca.agregar_socio(socio3)

# Prestar libros a los socios
biblioteca.prestar_libro(libro1, socio1, "2023-11-03")
biblioteca.prestar_libro(libro2, socio2, "2023-11-05")
biblioteca.prestar_libro(libro3, socio2, "2023-11-06")
biblioteca.prestar_libro(libro4, socio2, "2023-11-06")
biblioteca.prestar_libro(libro5, socio2, "2023-11-06")

# Obtener socios con más de 3 libros prestados
socios_con_mas_de_3_libros = biblioteca.obtener_socios_con_mas_de_3_libros()

# Imprimir información de libros y socios
print("Libros en la biblioteca:")
for libro in biblioteca.libros:
    print(libro)

print("\nSocios en la biblioteca:")
for socio in biblioteca.socios:
    print(socio)

# Imprimir socios con más de 3 libros prestados
print("\nSocios con mas de 3 libros prestados:")
for socio in socios_con_mas_de_3_libros:
    print(socio)
