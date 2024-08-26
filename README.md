Simple Pot Generator CLI
==================

This Go script scans a specified directory for files with specific extensions and generates a POT file with the strings found in those files.

Compilation
-----------

### On Linux and macOS

    go build -o <executable-name>

### On Windows

    go build -o <executable-name>.exe

Usage
-----

Once compiled, you can run the script using the following syntax:

    ./<executable-name> -i /example/path -o output-file.pot -p extensions

Where:

*   `-i`: Path to the directory to scan. If not specified, the current directory will be used.
*   `-o`: Name of the output POT file. The default is `to-translate.pot`.
*   `-p`: File extensions to search for, separated by slashes. The default value is `php/js`.
*   `-h`: Displays the script help.

### Example Usage

To scan a directory named `src` and generate a POT file with strings from files with `php` and `js` extensions:

    ./<executable-name> -i /project-path -o translation.pot

Contributions
-------------

Contributions are welcome. Please fork the repository and submit a pull request with your changes.

License
-------

This project is licensed under the MIT License.