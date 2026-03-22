<h1 align="center">Simple-Cache</h1>
<p align="center">
  <img src="https://socialify.git.ci/MostafaSensei106/Simple-Cache/image?custom_language=Go&font=KoHo&language=1&logo=https%3A%2F%2Favatars.githubusercontent.com%2Fu%2F138288138%3Fv%3D4&name=1&owner=1&pattern=Floating+Cogs&theme=Light" alt="Simple-Cache Banner">
</p>

<p align="center">
  <strong>A minimalist LRU Cache implementation in Go.</strong>

<p align="center">
  <a href="#about">About</a> •
  <a href="#features">Features</a> •
  <a href="#how-it-works">How It Works</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#technologies">Technologies</a> •
  <a href="#license">License</a>
</p>

---

## About

**Simple-Cache** is a lightweight, thread-safe (in concept) implementation of a **Least Recently Used (LRU)** cache. Built with Go, it demonstrates how to combine a **Hash Map** and a **Doubly Linked List** to achieve $O(1)$ time complexity for both lookups and updates. 

Whether you're looking for a simple caching layer or an educational piece on data structures, Simple-Cache provides a clean and idiomatic starting point.

---

## Features

- **$O(1)$ Performance**: Constant time complexity for `Add`, `Remove`, and `Check` operations.
- **LRU Eviction Policy**: Automatically removes the least recently used items when the maximum capacity is reached.
- **Dual Data Structure**: Utilizes a Go `map` for fast lookups and a custom `Doubly Linked List` for maintaining order.
- **Clean Architecture**: Decoupled models, constants, and cache logic for easy maintainability.

---

## How It Works

Simple-Cache uses a combination of two data structures:

1.  **Hash Map**: Stores keys and pointers to their corresponding nodes in the linked list. This allows for $O(1)$ access to any element.
2.  **Doubly Linked List (Queue)**: Maintains the order of elements based on usage. The "Most Recently Used" (MRU) element is always at the head, while the "Least Recently Used" (LRU) element is at the tail.

### Data Structure Visualization

```mermaid
graph TD
    %% Global Styles
    classDef mapNode fill:#e1f5fe,stroke:#01579b,stroke-width:2px,color:#01579b;
    classDef listNode fill:#fff3e0,stroke:#e65100,stroke-width:2px,color:#e65100;
    classDef sentinel fill:#eceff1,stroke:#455a64,stroke-width:2px,stroke-dasharray: 5 5;

    subgraph Memory [LRU Cache Structure]
        direction TB

        %% Hash Map Section
        subgraph HashMap [Hash Map: Quick Access]
            direction LR
            Key1["Key: 'apple'"] --- Hash1(( )) --> Pointer1
            Key2["Key: 'banana'"] --- Hash2(( )) --> Pointer2
        end

        %% Doubly Linked List Section
        subgraph DLL [Doubly Linked List: Recency Order]
            direction LR
            Head[<b>Head</b><br/>MRU] <--> Node1[<b>Node: 'apple'</b><br/>Value: 100]
            Node1 <--> Node2[<b>Node: 'banana'</b><br/>Value: 200]
            Node2 <--> Tail[<b>Tail</b><br/>LRU]
        end

        %% Connections from Map to List
        Pointer1 -.->|Points to| Node1
        Pointer2 -.->|Points to| Node2
    end

    %% Applying Classes
    class Key1,Key2,Pointer1,Pointer2 mapNode;
    class Node1,Node2 listNode;
    class Head,Tail sentinel;

    %% Annotations
    note1[O-1 Lookup] -.-> HashMap
    note2[O-1 Reorder] -.-> DLL
    ```

### Methods Explanation

-   **`Check(word string)`**: The main entry point. It checks if the word exists in the cache. 
    -   If **found**: It moves the node to the **Front** (Head).
    -   If **not found**: It creates a new node, adds it to the **Front**, and inserts it into the Hash Map.
-   **`Add(node *Node)`**: Inserts a node at the head of the queue. If the `MAX_SIZE` (default 8) is exceeded, it automatically triggers `Remove(Tail.Left)`.
-   **`Remove(node *Node)`**: Detaches a node from its current position in the list and deletes it from the Hash Map.

---

## Installation

Ensure you have **Go 1.18+** installed on your system.

### 🏗️ Build from Source

```bash
# Clone the repository
git clone https://github.com/MostafaSensei106/Simple-Cache.git
cd Simple-Cache

# Build the project
go build -o simple-cache main.go
```

---

## Usage

### 🚀 Running the Demo

The current implementation includes a demo in `cmd/root.go` that iterates through a list of words to demonstrate the LRU eviction policy.

```bash
go run main.go
```

### 📋 Example Output

```text
Simple Cache Start...
add: apple: 1 - [{apple}]
add: banana: 2 - [{banana} <--> {apple}]
add: cherry: 3 - [{cherry} <--> {banana} <--> {apple}]
...
add: kiwi max size reached removeing node: apple: 8 - [{kiwi} <--> {honeydew} <--> {grape} <--> {fig} <--> {elderberry} <--> {date} <--> {cherry} <--> {banana}]
Simple Cache End...
```

---

## Configuration

You can adjust the cache size by modifying the `MAX_SIZE` constant in `internal/constants/app_const.go`:

```go
package constants

const MAX_SIZE = 8 // Change this value to your preferred cache limit
```

---

## Technologies

| Technology          | Description                                                                                                 |
| ------------------- | ----------------------------------------------------------------------------------------------------------- |
| 🧠 **Golang**       | [go.dev](https://go.dev) — The core language powering Simple-Cache: fast and efficient                      |
| 🛠️ **Custom DS**    | Doubly Linked List & Hash Map implementation for optimal LRU performance                                  |
| 📊 **Mermaid**      | Used for architectural visualization in documentation                                                       |

---

## License

This project is licensed under the **MIT License**.
See the [LICENSE](LICENSE) file for full details.

<p align="center">
  Made with ❤️ by <a href="https://github.com/MostafaSensei106">MostafaSensei106</a>
</p>
