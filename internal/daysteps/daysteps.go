package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("wrong data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("bad steps value: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps must be > 0")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("bad duration value: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration must be > 0")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("steps must be > 0")
	}
	if ds.Duration <= 0 {
		return "", errors.New("duration must be > 0")
	}
	if ds.Weight <= 0 {
		return "", errors.New("weight must be > 0")
	}
	if ds.Height <= 0 {
		return "", errors.New("height must be > 0")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d. Дистанция составила %.2f км. Вы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
