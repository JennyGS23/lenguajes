package Data;

import java.util.LinkedList;

/*
Diferencias:

Eager Singleton: Este enfoque crea la instancia de la clase de inmediato al cargar la clase en memoria.
Esto garantiza que la instancia esté disponible de inmediato y es útil si tienes certeza de que la instancia
siempre será necesaria al inicio de la aplicación. Es más sencillo y seguro en entornos multi-hilo,
ya que la instancia se inicializa antes de que pueda haber problemas de concurrencia.

Lazy Singleton: En este enfoque, la instancia se crea solo cuando se solicita por primera vez.
Es útil si deseas retrasar la creación de la instancia hasta que realmente se necesite para ahorrar recursos.
Sin embargo, debes tener en cuenta problemas de concurrencia si varias partes de tu aplicación intentan acceder a
la instancia al mismo tiempo. Puedes utilizar mecanismos de sincronización para garantizar la seguridad en entornos multi-hilo.

En este código de Agenda:
La instancia única de la clase Agenda se crea como un campo estático private static final Agenda instance = new Agenda();.
Esto asegura que la instancia se cree de inmediato cuando se carga la clase en la memoria.
El constructor de Agenda es privado, lo que evita que se pueda crear otra instancia de la clase desde fuera de la misma.
El método estático getInstance() proporciona acceso a la instancia única y ya creada de Agenda.
El método getInstance() siempre devuelve la misma instancia de Agenda, garantizando que solo haya una instancia en toda la aplicación.
En resumen, la implementación actual en la clase Agenda es un ejemplo de un Eager Singleton, ya que la instancia se crea de inmediato
al cargar la clase en la memoria.
 */

public class Agenda {
    private static final Agenda instance = new Agenda();

    private LinkedList<Object> listaObjetos;
    private LinkedList<String> contactos;
    private LinkedList<String> eventos;

    public Agenda() {
        this.listaObjetos = new LinkedList<Object>();
        this.contactos = new LinkedList<String>();
        this.eventos = new LinkedList<String>();
    }

    public static Agenda getInstance() {
        return instance;
    }

    public LinkedList<String> getContactos() {
        return contactos;
    }

    public LinkedList<String> getEventos() {
        return eventos;
    }

    public LinkedList<Object> getListaObjetos() {
        return listaObjetos;
    }
}
