# factro-package-replacer

## Was macht das Programm?
Das Programm ist in der Lage, mehrere "Packages" eines bestimmten "Projects" in Factro gleichzeitig zu aktualisieren. Es ist in der Lage, alle Projects für ein bestimmtes Profjekt herunterzuladen, ein bestimmtes Feld, z.B. den Titel, nach einem bestimmten Wert, z.B. "04.2020", zu durchsuchen und diesen Wert überall durch einen anderen Wert , z.B. "04.2021" zu ersetzen. Anschließend werden die aktualisierten Packages wieder zum Factro Server geschickt, der sie in der Cloud aktualisiert. 

## Wie benutze ich das Programm?
1. factro-package-replacer.exe herunterladen und in ein Verzeichnis speichern.
2. Im selben Verzeichnis einen Ordner mit dem Namen "config" erstellen.
3. Im Ordner "config" eine Textdatei (Endung .txt) mit dem Namen "api_user_token" erstellen.
4. In diese Textdatei den API-Token (JWT) des Nutzers kopieren, für den die Packages aktualisiert werden sollen. Dabei sollte ein Testnutzer verwendet werden, dem nur das Projekt zugeordnet ist, dessen Packages aktualisiert werden sollen. Das API Token lässt sich in Factro generieren.
5. Die Textdatei abspeichern. 
6. "factro-package-replacer.exe" durch Doppelklicken ausführen.
7. Die Anweisungen auf dem Bildschirm befolgen.
8. Eingaben durch die Eingabetaste (Enter) bestätigen.
9. Nach erfolgreicher Ausführung in der Webanwendung überprüfen, ob die Aktualisierung erfolgreich durchgeführt werden konnte.


## Troubleshooting
| Problem                                                                         | Mögliche Ursachen                                                               |
|---------------------------------------------------------------------------------|---------------------------------------------------------------------------------|
| Es gibt keine Fehlermeldung, die Packages werden aber trotzdem nicht aktualisiert. | Das angegebene Administrator-Token besitzt nicht die notwendigen Berechtigungen |
