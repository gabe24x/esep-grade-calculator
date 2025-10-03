package esepunittests

type GradeCalculator struct {
	grades []Grade
	mode   string // "letter" (default) or "passfail"
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{grades: make([]Grade, 0), mode: "letter"}
}

func NewGradeCalculatorWithMode(mode string) *GradeCalculator {
	gc := NewGradeCalculator()
	if mode == "passfail" {
		gc.mode = "passfail"
	} else {
		gc.mode = "letter"
	}
	return gc
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{Name: name, Grade: grade, Type: gradeType})
}

func (gc *GradeCalculator) GetFinalGrade() string {
	num := gc.calculateNumericalGrade()

	if gc.mode == "passfail" {
		if num >= 70 {
			return "P"
		}
		return "F"
	}

	if num >= 90 {
		return "A"
	} else if num >= 80 {
		return "B"
	} else if num >= 70 {
		return "C"
	} else if num >= 60 {
		return "D"
	}
	return "F"
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	a := computeAverage(filterByType(gc.grades, Assignment))
	e := computeAverage(filterByType(gc.grades, Exam))
	s := computeAverage(filterByType(gc.grades, Essay))
	weighted := float64(a)*0.5 + float64(e)*0.35 + float64(s)*0.15
	return int(weighted)
}

func filterByType(all []Grade, typ GradeType) []Grade {
	out := make([]Grade, 0, len(all))
	for _, g := range all {
		if g.Type == typ {
			out = append(out, g)
		}
	}
	return out
}

func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for _, g := range grades {
		sum += g.Grade
	}
	return sum / len(grades)
}
