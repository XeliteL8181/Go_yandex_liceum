package company

import (
	"fmt"
	"sort"
)

var ValidPositions = []string{
	"директор",
	"зам. директора",
	"начальник цеха",
	"мастер",
	"рабочий",
}

var PositionPriority = map[string]int{
	"директор":       0,
	"зам. директора": 1,
	"начальник цеха": 2,
	"мастер":         3,
	"рабочий":        4,
}

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Worker struct {
	name     string
	position string
	income   uint
}

type Company struct {
	workers []Worker
}

func TotalIncome(salary, experience uint) uint {
	return salary * experience
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	flag := false
	for _, pos := range ValidPositions {
		if pos == position {
			flag = true
			break
		}
	}

	if !flag {
		return fmt.Errorf("некорректная должность %q", position)
	}
	newWorker := Worker{
		name: name, position: position, income: TotalIncome(salary, experience),
	}
	c.workers = append(c.workers, newWorker)

	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	sort.Slice(c.workers, func(i, j int) bool {
		incomeI := c.workers[i].income
		incomeJ := c.workers[j].income
		if incomeI != incomeJ {
			return incomeI > incomeJ
		}

		positionI := PositionPriority[c.workers[i].position]
		positionJ := PositionPriority[c.workers[j].position]
		if positionI != positionJ {
			return positionI < positionJ
		}

		nameI := c.workers[i].name
		nameJ := c.workers[j].name
		return nameI < nameJ
	})

	var result []string
	for _, w := range c.workers {
		result = append(result, fmt.Sprintf(
			"%s — %d — %s",
			w.name, w.income, w.position,
		))
	}

	return result, nil
}
