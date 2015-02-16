package domain

type MonthSummaryRepository interface {
	Store(monthSummary MonthSummary) (MonthSummary, error)
  Delete(monthSummary MonthSummary) error
  FindMonth(year int, month int) (MonthSummary, error)
}

type MonthSummary struct {
	ID string `datastore:"-"`
  Year int
	Month int
  Day1 string `datastore:",noindex"`
  Day2 string `datastore:",noindex"`
  Day3 string `datastore:",noindex"`
  Day4 string `datastore:",noindex"`
  Day5 string `datastore:",noindex"`
  Day6 string `datastore:",noindex"`
  Day7 string `datastore:",noindex"`
  Day8 string `datastore:",noindex"`
  Day9 string `datastore:",noindex"`
  Day10 string `datastore:",noindex"`
  Day11 string `datastore:",noindex"`
  Day12 string `datastore:",noindex"`
  Day13 string `datastore:",noindex"`
  Day14 string `datastore:",noindex"`
  Day15 string `datastore:",noindex"`
  Day16 string `datastore:",noindex"`
  Day17 string `datastore:",noindex"`
  Day18 string `datastore:",noindex"`
  Day19 string `datastore:",noindex"`
  Day20 string `datastore:",noindex"`
  Day21 string `datastore:",noindex"`
  Day22 string `datastore:",noindex"`
  Day23 string `datastore:",noindex"`
  Day24 string `datastore:",noindex"`
  Day25 string `datastore:",noindex"`
  Day26 string `datastore:",noindex"`
  Day27 string `datastore:",noindex"`
  Day28 string `datastore:",noindex"`
  Day29 string `datastore:",noindex"`
  Day30 string `datastore:",noindex"`
  Day31 string `datastore:",noindex"`
  Duration1 int `datastore:",noindex"` // future enhancement to track time spent
  Duration2 int `datastore:",noindex"` // future enhancement to track time spent
  Duration3 int `datastore:",noindex"` // future enhancement to track time spent
  Duration4 int `datastore:",noindex"` // future enhancement to track time spent
  Duration5 int `datastore:",noindex"` // future enhancement to track time spent
  Duration6 int `datastore:",noindex"` // future enhancement to track time spent
  Duration7 int `datastore:",noindex"` // future enhancement to track time spent
  Duration8 int `datastore:",noindex"` // future enhancement to track time spent
  Duration9 int `datastore:",noindex"` // future enhancement to track time spent
  Duration10 int `datastore:",noindex"` // future enhancement to track time spent
  Duration11 int `datastore:",noindex"` // future enhancement to track time spent
  Duration12 int `datastore:",noindex"` // future enhancement to track time spent
  Duration13 int `datastore:",noindex"` // future enhancement to track time spent
  Duration14 int `datastore:",noindex"` // future enhancement to track time spent
  Duration15 int `datastore:",noindex"` // future enhancement to track time spent
  Duration16 int `datastore:",noindex"` // future enhancement to track time spent
  Duration17 int `datastore:",noindex"` // future enhancement to track time spent
  Duration18 int `datastore:",noindex"` // future enhancement to track time spent
  Duration19 int `datastore:",noindex"` // future enhancement to track time spent
  Duration20 int `datastore:",noindex"` // future enhancement to track time spent
  Duration21 int `datastore:",noindex"` // future enhancement to track time spent
  Duration22 int `datastore:",noindex"` // future enhancement to track time spent
  Duration23 int `datastore:",noindex"` // future enhancement to track time spent
  Duration24 int `datastore:",noindex"` // future enhancement to track time spent
  Duration25 int `datastore:",noindex"` // future enhancement to track time spent
  Duration26 int `datastore:",noindex"` // future enhancement to track time spent
  Duration27 int `datastore:",noindex"` // future enhancement to track time spent
  Duration28 int `datastore:",noindex"` // future enhancement to track time spent
  Duration29 int `datastore:",noindex"` // future enhancement to track time spent
  Duration30 int `datastore:",noindex"` // future enhancement to track time spent
  Duration31 int `datastore:",noindex"` // future enhancement to track time spent
}