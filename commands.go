package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	commandList := map[string]cliCommand{
		"exit": {
			name:        "Exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "Help",
			description: "See available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "Map",
			description: "See a list of 20 locations. Use the command again to see the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "Mapb",
			description: "See the previous page from the list of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "Explore",
			description: "Explore an area and see the list of pokemon that can be found. Syntax: explore <area>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "Catch",
			description: "Try to catch a pokemon. Syntax: catch <pokemon>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "Inspect",
			description: "Inspect the states of a pokemon you have caught. Syntax: inspect <pokemon>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "Pokedex",
			description: "Show the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
	return commandList
}
