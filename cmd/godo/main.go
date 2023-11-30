package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/ritikdhasmana/godo"
)

func main() {
	add := flag.Bool("add", false, "Add a new todo (Type the task as flag arguement or in new line)")
	finish := flag.Int("finish", 0, "Mark todo as finished whose id matches with the id passed")
	id := flag.Int("id", 0, "List the todo whose id matches with the id passed")
	setStatus := flag.Int("set-status", 0, "Set a custom status for todo with matching id (Pass the new status as flag arguement after passing id). Example: 'godo --set-status 1 apple'")
	delete := flag.Int("delete", 0, "Deletes the todo with the passed id")
	list := flag.Bool("list", false, "List all todos")

	flag.Parse()

	todos := &todo.Todos{}

	//initial setup
	printErrorAndExit(todos.Load())

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		printErrorAndExit(err)

		todos.Add(task)
		err = todos.Store()
		printErrorAndExit(err)
		printResult(todos, "Added!")

	case *setStatus > 0:
		status, err := getStatus(flag.Args()...)
		printErrorAndExit(err)

		printErrorAndExit(todos.UpdateStatus(*setStatus, status))
		printErrorAndExit(todos.Store())
		printResult(todos, "Updated!")

	case *finish > 0:
		printErrorAndExit(todos.UpdateStatus(*finish, "Done"))
		printErrorAndExit(todos.Store())
		printResult(todos, "Updated!")

	case *id > 0:
		todos.PrintTodo(*id)

	case *list:
		todos.PrintTodos()

	case *delete > 0:
		printErrorAndExit(todos.Delete(*delete))
		printErrorAndExit(todos.Store())
		printResult(todos, "Deleted!")

	default:
		fmt.Fprintln(os.Stdout, "Invalid command! Type `godo --help` to see all available commands.")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil

}

func getStatus(args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	} else {
		return "", errors.New("empty status is not allowed, pass status along with task id")
	}
}

func printErrorAndExit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func printResult(todos *todo.Todos, operation string) {
	fmt.Fprintln(os.Stdout, operation)
	fmt.Fprintln(os.Stdout, "")

	todos.PrintTodos()
}
