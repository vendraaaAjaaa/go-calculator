# Pac-Pac-GO - CLI Pacman Game

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey?style=for-the-badge)
![Educational Project](https://img.shields.io/badge/Educational-Project-blue?style=for-the-badge)

A classic Pacman-inspired game built entirely for the command-line interface (CLI) using Go. This project was developed as a final assignment for the **Introduction to Algorithms and Programming** course. Navigate through the maze, collect food, and avoid walls in this terminal-based adventure!

## 📚 About This Project

This project was created as a **final assignment** for the **Introduction to Algorithms and Programming** course. It demonstrates the application of fundamental programming concepts learned throughout the semester in a practical, game-based implementation.

### Course Objectives Demonstrated:
- Understanding of basic programming constructs
- Implementation of algorithms for game logic
- Problem-solving through computational thinking
- Code organization and documentation

## 🎮 Features

- **CLI Gameplay**: Play Pacman directly in your terminal
- **Simple Controls**: Intuitive `WASD` movement controls
- **Score System**: Earn points by collecting food dots
- **Maze Navigation**: Find your way through custom-designed walls
- **Win Condition**: Clear all dots to win the game
- **Cross-Platform**: Works on Windows, macOS, and Linux

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher installed
- Terminal/Command Prompt

### Installation & Running

1. **Clone the repository**
   ```bash
   git clone https://github.com/vendraaaAjaaa/Pac-Pac-GO.git
   cd Pac-Pac-GO
   ```

2. **Run the game**
   ```bash
   go run main.go
   ```

## 🕹️ How to Play

### Controls
- **W** - Move Up
- **A** - Move Left  
- **S** - Move Down
- **D** - Move Right
- **Q** - Quit Game

### Game Elements
- **C** - Pacman (You!)
- **#** - Wall (Cannot pass through)
- **.** - Food (Collect for 10 points each)
- **␣** - Empty space (Walkable area)

### Objective
Collect all the food dots (`.`) in the maze to win! Each dot gives you 10 points. The game ends when you've collected all 280 points.

## 📸 Game Preview

```
=== PACMAN GO (CLI VERSION) ===
Skor: 0
Kontrol: w (atas), s (bawah), a (kiri), d (kanan), q (keluar)

# # # # # # # # # # 
# C . . # . . . . # 
# . # . # . # # . # 
# . # . . . . . . # 
# . # # # # # . # # 
# . . . . . . . . # 
# # # # # # # # # # 

Gerak: 
```

## 🏗️ Project Structure

The game implements several programming concepts learned in the course:

### Core Game Components
1. **Maze Representation** - 2D array with different values for walls, food, and empty spaces
2. **Game Loop** - Continuous cycle of drawing, input handling, and game logic
3. **Collision Detection** - Prevents Pacman from passing through walls
4. **Score Tracking** - Real-time score calculation and display

### Programming Concepts Applied
| Concept | Implementation in Code |
|---------|------------------------|
| **Variables & Data Types** | Integer arrays for maze, variables for player position |
| **Control Structures** | If-else for movement logic, for loops for maze rendering |
| **Arrays & 2D Arrays** | Maze represented as 2D integer array |
| **Functions** | `clearScreen()` function for cross-platform terminal clearing |
| **Input/Output** | User input handling and console output |
| **Algorithm Design** | Game loop algorithm and collision detection |

## 📊 Maze Layout

The game features a 7x10 maze with:
- Outer border walls
- Internal maze structure
- Strategic food placement
- Starting position at (1,1)

## 🛠️ Technical Details

### Built With
- **Go** (Golang) - Programming language
- **Standard Library Only** - No external dependencies
- **exec Package** - For cross-platform terminal clearing

### File Structure
```
Pac-Pac-GO/
├── Flowchart-Pac-Pac-GO.png     # Flowchart diagram
├── README.md                    # Project documentation
├── main.go                      # Main game source code
└── pacpacgo.pseudocode          # Pseudocode implementation
```

## 🎯 Learning Objectives

This project demonstrates key concepts from the **Introduction to Algorithms and Programming** course:

### Fundamental Programming Concepts
1. **Problem Analysis** - Breaking down game requirements into computational steps
2. **Algorithm Design** - Designing the game loop and collision logic
3. **Data Structures** - Using 2D arrays to represent game state
4. **Control Flow** - Implementing game logic with conditionals and loops
5. **Modular Programming** - Organizing code into logical functions

### Specific Course Competencies
- Understanding and implementing basic game algorithms
- Applying control structures in a practical project
- Developing problem-solving skills through game development
- Learning cross-platform development considerations

## 🤝 Contributing

Feel free to fork this project and make improvements! Some ideas:
- Add different maze layouts
- Implement ghosts/enemies
- Add difficulty levels
- Create a scoring leaderboard
- Add color to the display

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 👥 Authors

- **Vendra** - [@vendraaaAjaaa](https://github.com/vendraaaAjaaa)
- **Harell** - [@hunchoTX](https://github.com/hunchoTX)
- **Naufal** - [@palapis](https://github.com/palapis)

## 🙏 Acknowledgments

This project was developed as the final assignment for the **Introduction to Algorithms and Programming** course. We would like to express our gratitude to:

- **Our Lecturer** - For guidance and instruction throughout the course
- **Lab Assistants** - For their support during practical sessions
- **Faculty of Informatics** - For providing the learning environment and resources
- **Inspired by the classic Pacman game** - For the original game concept
- **The Go/Golang community** - For excellent documentation and resources
- **Our classmates** - For collaborative learning and peer support

### Special Thanks
This project represents the culmination of our learning in the Introduction to Algorithms and Programming course. It has been a valuable experience in applying theoretical concepts to a practical, functional application.

---

**Enjoy the game!** If you get stuck, remember: the walls are marked with `#` and you need to navigate around them to collect all the dots!

*This project demonstrates the practical application of algorithms and programming fundamentals learned throughout the course.*
