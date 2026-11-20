package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("wrong data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("bad steps value: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps must be > 0")
	}

	trainingType := parts[1]
	if trainingType == "" {
		return errors.New("empty training type")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("bad duration value: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration must be > 0")
	}

	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 {
		return "", errors.New("steps must be > 0")
	}
	if t.Duration <= 0 {
		return "", errors.New("duration must be > 0")
	}
	if t.Weight <= 0 {
		return "", errors.New("weight must be > 0")
	}
	if t.Height <= 0 {
		return "", errors.New("height must be > 0")
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("unknown training type")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s Длительность: %.2f ч. Дистанция: %.2f км. Скорость: %.2f км/ч Сожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}