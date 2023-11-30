package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
	"path/filepath"
	"github.com/alexeyco/simpletable"
)

type item struct {
	Task      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todos []item

const (
	todoFileName = ".godo.json"
)

func getAbsolutePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, todoFileName), nil
}



func (t *Todos) Add(task string) {
	todo := item{
		Task:      task,
		Status:    "Ongoing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) UpdateStatus(taskId int, status string) error {
	if taskId <= 0 || taskId > len(*t) {
		return errors.New("invalid index")
	}
	list := *t
	list[taskId-1].UpdatedAt = time.Now()
	list[taskId-1].Status = status
	return nil
}

func (t *Todos) Delete(taskId int) error {
	if taskId <= 0 || taskId > len(*t) {
		return errors.New("invalid index")
	}
	list := *t
	*t = append(list[:(taskId-1)], list[taskId:]...)
	return nil
}

func (t *Todos) Load() error {
	fileName, err := getAbsolutePath()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store() error {
	fileName, err := getAbsolutePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func (t *Todos) PrintTodos() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Todo"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Created"},
			{Align: simpletable.AlignCenter, Text: "Updated"},
		},
	}
	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		status := blue(item.Status)
		if item.Status == "Done" {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			status = green(item.Status)
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: status},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.UpdatedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos.", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) PrintTodo(id int) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Todo"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Created"},
			{Align: simpletable.AlignCenter, Text: "Updated"},
		},
	}
	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		if idx == id {
			task := blue(item.Task)
			status := blue(item.Status)
			if item.Status == "Done" {
				task = green(fmt.Sprintf("\u2705 %s", item.Task))
				status = green(item.Status)
			}
			cells = append(cells, []*simpletable.Cell{
				{Text: fmt.Sprintf("%d", idx)},
				{Text: task},
				{Text: status},
				{Text: item.CreatedAt.Format(time.RFC822)},
				{Text: item.UpdatedAt.Format(time.RFC822)},
			})
		}
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	count := 0
	for _, item := range *t {
		if item.Status != "Done" {
			count++
		}
	}
	return count
}
