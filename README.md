# 🧢 Go Pokedex REPL
Fully functional command-line Pokedex built in Go! This REPL (Read-Eval-Print-Loop) uses the PokeAPI to fetch real-time Pokémon data, complete with HTTP networking, JSON decoding, and in-memory caching.

## 📝 Description
This project is an interactive CLI application that lets users explore Pokémon data from their terminals. Currently supports fetching Pokémon info, listing areas, and more - all while caching responses for optimzied performance.

## 💡 Motivation
The goal of the project was to use Go and APIs to make something interesting and functional. The Pokedex served as an awesome idea to put backend development skills to fun use.

## 🚀 Quickstart

### Prerequisites
- Go 1.19 or higher installed.
- Internet connection for PokeAPI.

### Installation
Clone the repo:
```
git clone https://github.com/josequiceno2000/pokedex
cd pokedex
```
Build and run the project:
```
go build -o pokedex
./pokedex
```

## 🛠️ Usage
Once launched, you will be greeted with an interactive REPL interface.
Here are some of the supported commands:
| **Command** | **Description** | 
| -------- | -------- |
| `help` | Shows list of available commands | 
| `exit` | Exits the Pokedex CLI |
| `map` | Shows paginated list of location areas | 
| `explore <location_name>` | Lists Pokémon found in given area | 
| `catch <pokemon_name>` | Tries to catch the Pokémon by name | 
| `inspect <pokemon_name>` | Presents details of caught Pokémon | 
| `pokedex` | Lists all caught Pokémon |  

## 🤝 Contributing
Contributions are so welcome and appreciated! If you wanna help improve this project or add features, get started by:
1. Forking the repo
2. Creating a new branch: `git checkout -b feature-name`
3. Making your changes and adding tests
4. Commiting and pushing: `git commit -m "Add feature"` then `git push origin feature-name`
5. Open a pull request!

### Current Contribution Possibilities:
- Add support for previous command history
- Simulate Pokémon battles
- Save caught Pokémon between sessions
- Add support for evolution, experience, and leveling up
- Improve code structure and test coverage