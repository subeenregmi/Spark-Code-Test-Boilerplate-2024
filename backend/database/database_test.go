package database

import ( 
    "reflect"
	"testing"
)

func TestAdd(t *testing.T) {
    db := CreateNewDB()

    tm := TodoModel{Title: "Going for a run", Description:"With George."}
    db.Add(tm);

    if v := db.TodoList[1]; v != tm {
        t.Fatalf(`db.Add(tm) has failed. Expected %v. Got %v.`, tm, v) 
    }

    N := 100
    for i := 0; i < N-1; i++ {
        db.Add(tm);
    }
    
    for i := 2; i < N-1; i++ {
        if v := db.TodoList[i]; v != tm {
            t.Fatalf(
                `Adding %d elements has failed. At i = %d, expected %v. Got %v.`,
                N, i, tm, v,
            ) 
        }
    }
}

func TestGetAll(t *testing.T) {
    db := CreateNewDB()
    
    tm := TodoModel{Title: "Going for a run", Description:"With George."}
    db.Add(tm);

    sample := []TodoModel{tm}

    if v := db.GetAll(); !reflect.DeepEqual(sample, v) {
        t.Fatalf(`db.GetAll() has failed. Expected %v. Got %v.`, sample, v) 
    }

    tm = TodoModel{Title: "Buy groceries", Description: "Milk, eggs, bread, and butter."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Complete homework", Description: "Finish math and science assignments."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Schedule dentist appointment", Description: "Call the clinic and book a slot for Friday."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Clean the garage", Description: "Organize tools and throw away old items."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Prepare presentation", Description: "Create slides for Monday's meeting."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Watch tutorial videos", Description: "Go through Go programming basics on YouTube."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Plan weekend trip", Description: "Research places and book accommodations."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Water the plants", Description: "Focus on the indoor plants in the living room."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Fix bike", Description: "Check brakes and oil the chain."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Organize bookshelf", Description: "Sort books by genre and donate unused ones."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Prepare dinner", Description: "Try the new pasta recipe with garlic bread."}
    sample = append(sample, tm)
    db.Add(tm);

    tm = TodoModel{Title: "Attend yoga class", Description: "6 PM at the community center."}
    sample = append(sample, tm)
    db.Add(tm);

    if db.Count != 13 {
        t.Fatalf(`Adding to database has a counting error. Expected %d. Got %d.`, db.Count, 13)
    }
}
