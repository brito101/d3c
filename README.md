# D3C - Distributed Command & Control

A distributed command and control system built in Go, featuring a modular agent architecture with separate command implementations.

## 🏗️ Architecture

### Agent Structure
```
agent/
├── agent.go                 # Main agent entry point
├── agent_helpers/
│   └── comand_helper.go     # Command validation and mapping
├── commands/                # Modular command implementations
│   ├── cd.go               # Change directory
│   ├── ls.go               # List files
│   ├── ps.go               # Process list
│   ├── pwd.go              # Print working directory
│   ├── whoami.go           # Current user
│   ├── send.go             # File upload
│   ├── get.go              # File download
│   ├── sleep.go            # Update heartbeat
│   └── default.go          # Shell command execution
└── interfaces/
    └── command_interface.go # Command interface definition
```

### Server Structure
```
server/
├── d3c.go                  # Main server entry point
├── commands/               # Server-side commands
├── helpers/                # Server utilities
└── listeners/
    └── network_listener.go # Network communication
```

## 🚀 Compilation

### Agent
```bash
# Standard compilation
go build -o agent agent/agent.go

# Hide console on Windows
GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o agent.exe agent/agent.go
```

### Server
```bash
go build -o server server/d3c.go
```

## 📋 Available Commands

### Agent Commands
| Command | Description | Usage |
|---------|-------------|-------|
| `cd` | Change directory | `cd <directory>` |
| `ls` | List files | `ls [directory]` |
| `ps` | List processes | `ps` |
| `pwd` | Print working directory | `pwd` |
| `whoami` | Current user | `whoami` |
| `send` | Upload file | `send <file>` |
| `get` | Download file | `get <file>` |
| `sleep` | Update heartbeat | `sleep <seconds>` |
| `default` | Shell execution | Any shell command |

### Server Commands
| Command | Description | Usage |
|---------|-------------|-------|
| `show` | List agents | `show` |
| `select` | Select agent | `select <agent_id>` |
| `send` | Send file to agent | `send <file>` |
| `get` | Get file from agent | `get <file>` |

## 🔧 Features

### Modular Command System
- Each command is implemented as a separate module
- Consistent interface across all commands
- Easy to extend with new commands

### Error Handling
- Comprehensive error handling in all commands
- Graceful connection retry on network failures
- Detailed error messages for debugging

### Cross-Platform Support
- Windows: PowerShell execution
- Linux: Bash execution
- Automatic platform detection

### File Operations
- Secure file upload/download
- File validation before operations
- Error handling for file operations

## 🛠️ Development

### Adding New Commands

1. Create new command file in `agent/commands/`
2. Implement the `Command` interface
3. Add command to `agent_helpers/comand_helper.go`
4. Update mapping in `agent/agent.go`

Example:
```go
// agent/commands/newcommand.go
type NewCommand struct {
    Command string
}

func (instance NewCommand) Exec() (response string, err error) {
    // Implementation here
    return response, nil
}
```

### Command Interface
```go
type Command interface {
    Exec() (response string, err error)
}
```

## 🔒 Security Features

- MD5-based agent ID generation
- Connection validation
- File operation security
- Error logging and monitoring

## 📊 Configuration

### Agent Configuration
- Server address: `127.0.0.1:9090`
- Default heartbeat: 5 seconds
- Configurable via `sleep` command

### Network Configuration
- TCP communication
- Gob encoding for message serialization
- Automatic reconnection on failure

## 🐛 Troubleshooting

### Common Issues

1. **Connection Failed**
   - Check if server is running
   - Verify network connectivity
   - Check firewall settings

2. **Command Not Found**
   - Verify command exists in mapping
   - Check command implementation
   - Review error logs

3. **File Operations Fail**
   - Check file permissions
   - Verify file path
   - Ensure file exists

## 📝 License

This project is part of the D3C distributed command and control system.

## 🤝 Contributing

1. Follow the modular command structure
2. Implement proper error handling
3. Add comprehensive tests
4. Update documentation
