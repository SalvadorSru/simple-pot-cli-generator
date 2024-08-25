PO Generator
============

Version: 1.0.0

Description
-----------

This tool scans a specified folder for files containing specific translation functions and generates a Portable Object Template (POT) file.

Usage
-----

The program can be run with the following command-line arguments:

        -i, --input <path>     Path to the folder to scan (default is current directory)
        -o, --output <file>    Name of the output POT file (default is 'to-translate.pot')
        -h, --help             Display help information
    

Examples
--------

Run the program with a specific input directory and output file:

        ./po-generator -i ./src -o output.pot
    

Run the program with default settings:

        ./po-generator
    

Requirements
------------

Go programming language (version 1.16 or higher) must be installed.

Installation
------------

Clone the repository and navigate into the project directory:

        git clone https://github.com/salvadorsru/po-generator.git
        cd po-generator
    

Tutorial: Making the Program Executable
---------------------------------------

### For Linux and macOS

1\. Open a terminal.

2\. Navigate to the directory where the \`po-generator\` code is located.

3\. Run the following command to build the program:

        go build -o po-generator
    

4\. Make the program executable by running:

        chmod +x po-generator
    

5\. You can now run the program with:

        ./po-generator
    

### For Windows

1\. Open Command Prompt or PowerShell.

2\. Navigate to the directory where the \`po-generator\` code is located.

3\. Build the program by running:

        go build -o po-generator.exe
    

4\. You can now run the program with:

        po-generator.exe
    

License
-------

This project is licensed under the MIT License.

* * *

Generador de PO
===============

Versión: 1.0.0

Descripción
-----------

Esta herramienta escanea una carpeta especificada en busca de archivos que contienen funciones de traducción específicas y genera un archivo de plantilla de objeto portátil (POT).

Uso
---

El programa se puede ejecutar con los siguientes argumentos de línea de comandos:

        -i, --input <path>     Ruta a la carpeta a escanear (el valor predeterminado es el directorio actual)
        -o, --output <file>    Nombre del archivo POT de salida (el valor predeterminado es 'to-translate.pot')
        -h, --help             Mostrar información de ayuda
    

Ejemplos
--------

Ejecuta el programa con un directorio de entrada y un archivo de salida específicos:

        ./po-generator -i ./src -o output.pot
    

Ejecuta el programa con la configuración predeterminada:

        ./po-generator
    

Requisitos
----------

Se debe instalar el lenguaje de programación Go (versión 1.16 o superior).

Instalación
-----------

Clona el repositorio y navega al directorio del proyecto:

        git clone https://github.com/salvadorsru/po-generator.git
        cd po-generator
    

Tutorial: Haciendo el Programa Ejecutable
-----------------------------------------

### Para Linux y macOS

1\. Abre una terminal.

2\. Navega al directorio donde se encuentra el código de \`po-generator\`.

3\. Ejecuta el siguiente comando para construir el programa:

        go build -o po-generator
    

4\. Haz el programa ejecutable ejecutando:

        chmod +x po-generator
    

5\. Ahora puedes ejecutar el programa con:

        ./po-generator
    

### Para Windows

1\. Abre Command Prompt o PowerShell.

2\. Navega al directorio donde se encuentra el código de \`po-generator\`.

3\. Construye el programa ejecutando:

        go build -o po-generator.exe
    

4\. Ahora puedes ejecutar el programa con:

        po-generator.exe
    

Licencia
--------

Este proyecto está bajo la Licencia MIT.