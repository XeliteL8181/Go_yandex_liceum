package company

import "testing"

func TestCompany_SortWorkers(t *testing.T) {
    company := &Company{}

    _ = company.AddWorkerInfo("Михаил", "директор", 1000, 12)
    _ = company.AddWorkerInfo("Андрей", "мастер", 1180, 10)
    _ = company.AddWorkerInfo("Игорь", "зам. директора", 1000, 11)
    _ = company.AddWorkerInfo("Олег", "рабочий", 1000, 10)

    result, err := company.SortWorkers()
    if err != nil {
        t.Fatalf("SortWorkers() вернул ошибку: %v", err)
    }

    expected := []string{
        "Михаил — 12000 — директор",
        "Андрей — 11800 — мастер",
        "Игорь — 11000 — зам. директора",
        "Олег — 10000 — рабочий",
    }

    if len(result) != len(expected) {
        t.Fatalf("Ожидалось %d элементов, а получено %d", len(expected), len(result))
    }

    for i := range expected {
        if result[i] != expected[i] {
            t.Errorf("На позиции %d ожидалось: %q, получено: %q", i, expected[i], result[i])
        }
    }
}