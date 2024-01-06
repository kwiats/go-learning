package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QuizGames interface {
	StartQuiz([]Quiz) int64
}
type Question struct {
	A, B     int64
	Operator string
}

type Answer struct {
	Value int64
}

type Quiz struct {
	Question      Question
	CorrectAnswer Answer
	YourAnswer    Answer
}

var Operators string = "+-*/"

func (q *Quiz) CheckAnswer() bool {
	return q.CorrectAnswer.Value == q.YourAnswer.Value
}

func (q *Quiz) CalculateCorrectAnswer() {
	switch q.Question.Operator {
	case "+":
		q.CorrectAnswer.Value = q.Question.A + q.Question.B
	case "-":
		q.CorrectAnswer.Value = q.Question.A - q.Question.B
	}
}

type ScoreTracker struct {
	Mistakes int64
}

func (c *ScoreTracker) IncrementMistakes() {
	c.Mistakes++
}

func (q *Quiz) CreateQuiz(a, b int64, operator string) {
	question := Question{
		A:        a,
		Operator: operator,
		B:        b,
	}
	q.Question = question
	q.CalculateCorrectAnswer()
}

func readFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	return records
}

func findOperator(record string) string {
	operators := strings.Split(Operators, "")
	for _, operator := range operators {
		if strings.Contains(record, operator) {
			return operator
		}
	}
	return ""
}

func CreateListOfQuizes(fileName string) []Quiz {
	records := readFile(fileName)

	quizes := make([]Quiz, 0, len(records))
	for _, record := range records {
		operator := findOperator(record[0])
		if operator == "" {
			continue
		}
		equation := strings.Split(record[0], operator)

		if len(equation) != 2 {
			fmt.Printf("Invalid equation format: %s\n", record[0])
			continue
		}

		a, err := strconv.ParseInt(strings.TrimSpace(equation[0]), 10, 64)
		if err != nil {
			fmt.Printf("Error parsing number: %s\n", equation[0])
			continue
		}

		b, err := strconv.ParseInt(strings.TrimSpace(equation[1]), 10, 64)
		if err != nil {
			fmt.Printf("Error parsing number: %s\n", equation[1])
			continue
		}

		quiz := Quiz{}

		quiz.CreateQuiz(a, b, operator)
		quizes = append(quizes, quiz)

	}

	return quizes
}

func (tracker *ScoreTracker) StartQuiz(quizes []Quiz) int64 {
	var userAnswer int64
	for i := 0; i < len(quizes); i++ {
		fmt.Printf("Question %d: What is %v %v %v? ", i+1, quizes[i].Question.A, quizes[i].Question.Operator, quizes[i].Question.B)

		_, err := fmt.Scan(&userAnswer)
		if err != nil {
			fmt.Println("Inncorrect answer, try again.")
			i--
			continue
		}
		quizes[i].YourAnswer.Value = userAnswer
		if !quizes[i].CheckAnswer() {
			tracker.IncrementMistakes()
		}
		fmt.Printf("Your answer: %v. Correct: %v.\n", quizes[i].YourAnswer.Value, quizes[i].CorrectAnswer.Value)

	}
	return tracker.Mistakes
}

func main() {
	var quizGames QuizGames = &ScoreTracker{}
	quizes := CreateListOfQuizes("problems.csv")
	mistakes := quizGames.StartQuiz(quizes)

	correctAnswers := int64(len(quizes)) - mistakes

	fmt.Printf("Your score: %v/%v! \n", correctAnswers, len(quizes))

}
