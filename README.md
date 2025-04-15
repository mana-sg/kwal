# kv-log-store

`kv-log-store` is a lightweight in-memory key-value store with persistence achieved through a write-ahead log (WAL). It supports basic operations like setting, retrieving, and deleting key-value pairs, ensuring data durability by logging all operations to a file.

## Features

- **In-Memory Storage**: Fast access to key-value pairs stored in memory.
- **Write-Ahead Logging**: Persistent storage of operations in a log file (`log.bin`).
- **Basic Operations**:
    - `set <key> <value>`: Add or update a key-value pair.
    - `get <key>`: Retrieve the value for a given key.
    - `del <key>`: Delete a key-value pair.

## Installation

1. Clone the repository:
     ```sh
     git clone https://github.com/mana-sg/kv-log-store.git
     cd kv-log-store
     ```

2. Build the binary:
     ```sh
     make build
     ```

3. Install the binary to your system:
     ```sh
     sudo make install
     ```

4. To uninstall:
     ```sh
     sudo make uninstall
     ```

## Usage

Start the application:
```sh
kv-log-store
```

Use the following commands to interact with the key-value store:

### Set a key-value pair:
```sh
set <key> <value>
```
Example:
```sh
set name Alice
```

### Get the value of a key:
```sh
get <key>
```
Example:
```sh
get name
```

### Delete a key-value pair:
```sh
del <key>
```
Example:
```sh
del name
```

Exit the application by pressing `Ctrl+C`.

## Development

To clean up build artifacts:
```sh
make clean
```