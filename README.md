# gorue-common
Implementations of gorue components

## Prompter

### Simple
The simple prompter reads from an io.Reader and writes to an io.Writer.

## State

### fs
`FileStorage` will store and retrieve state from the filesystem.

### mem
`InMemoryStorage` will store and retrieve state from local memory.
State will be lost when `InMemoryStorage` is destroyed.