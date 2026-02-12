package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	carsNum := float64(productionRate) * (successRate / 100)
	return carsNum
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(cars int) uint {
	groupsOfTen := cars / 10
	remainingCars := cars % 10

	costForGroups := groupsOfTen * 95000
	costForRemaining := remainingCars * 10000

	totalCost := costForGroups + costForRemaining
	return uint(totalCost)
}
