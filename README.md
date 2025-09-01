# kv-log-store
`kv-log-store` is a lightweight in-memory key-value store with persistence achieved through a write-ahead log (WAL). It supports basic operations like setting, retrieving, and deleting key-value pairs, ensuring data durability by logging all operations to a file.

## Features
- **In-Memory Storage**: Fast access to key-value pairs stored in memory.
- **Write-Ahead Logging**: Persistent storage of operations in a log file (`log.bin`).
- **HTTP API**: RESTful endpoints for remote access to the key-value store.
- **Background Service**: Runs as a daemon process managed by the `kls` command.
- **Basic Operations**:
    - Set key-value pairs via HTTP API
    - Retrieve values by key via HTTP API  
    - Delete key-value pairs via HTTP API

## Installation
1. Clone the repository:
     ```sh
     git clone https://github.com/mana-sg/kv-log-store.git
     cd kv-log-store
     ```
2. Create a `.env` file in the project root:
     ```sh
     echo "SERVER_LOCATION=$(pwd)" > .env
     ```
3. Build the binary:
     ```sh
     make build
     ```
4. Install the binary to your system:
     ```sh
     sudo make install
     ```
5. To uninstall:
     ```sh
     sudo make uninstall
     ```

## Usage

### Service Management
Use the `kls` command to manage the kv-log-store service running in the background:

#### Start the service:
```sh
kls start
```

#### Stop the service:
```sh
kls stop
```

#### Restart the service:
```sh
kls restart
```

The service runs as a background daemon and exposes HTTP endpoints for interaction.

## HTTP API

When running as a service, kv-log-store exposes a RESTful API on port 8080:

### Set a key-value pair:
```
POST /set
Content-Type: application/json

{
    "key": "name",
    "value": "Alice"
}
```

### Get the value of a key:
```
GET /get?key=name
```

### Delete a key-value pair:
```
POST /delete
Content-Type: application/json

{
    "key": "name"
}
```

### Compact the log file:
```
POST /size/compact
```

### Get storage size information:
```
GET /size/get
```

## Development
To clean up build artifacts:
```sh
make clean
```
