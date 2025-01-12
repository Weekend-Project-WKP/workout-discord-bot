package helpers

import (
	"workoutbot/internal/db"
	"workoutbot/internal/models"
)

func CalculatePoints(categoryPoint int, categoryPointInterval int, workoutLength int) float64 {
	return float64((workoutLength / categoryPointInterval) * categoryPoint)
}

func CreateWorkouts(workouts []models.Workout) error {
	_, err := db.WorkoutsInsertMany(workouts)
	return err
}
