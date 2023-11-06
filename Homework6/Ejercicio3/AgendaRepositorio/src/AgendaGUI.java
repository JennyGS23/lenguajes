/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */

import Data.Agenda;
import Data.ObjAgenda;
import GUI.AgendaMain;
/*
En Agenda se encuentra la implementacion y explicación de Singleton
En Agregar una fecha cambié el formato de fecha de agrega como 12-04-23
ObjAgendaFactory y ObjAgendaFactoryImple cree esos archivos para implementar el patrón de Factory
 */
/**
 *
 * @author oviquez
 */
public class AgendaGUI {
    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        
        
        Agenda miAgenda = new Agenda();
        
        AgendaMain guiMain = new AgendaMain(miAgenda);
        guiMain.setVisible(true);
    }
    
}
