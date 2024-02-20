## Install
go install -v github.com/sendcon/resolvehost@latest

## Usage

1. Place your list of domains in a text file, with one domain per line.
2. Run the program:
    resolvehost domain.txt

Flags:
- `-i4` to display only IPv4 addresses.
- `-i6` to display only IPv6 addresses.

Example: resolvehost -i4 domain.txt
