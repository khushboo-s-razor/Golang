package dayone

import "fmt"

// Employee interface
type Employee interface {
	CalculateSalary() float64
}

// FullTime struct
type FullTime struct {
	dailyRate  float64
	daysWorked int
}

// Contractor struct
type Contractor struct {
	dailyRate  float64
	daysWorked int
}

// Freelancer struct
type Freelancer struct {
	hourlyRate  float64
	hoursWorked int
}

// CalculateSalary method for FullTime
func (f FullTime) CalculateSalary() float64 {
	return f.dailyRate * float64(f.daysWorked)
}

// CalculateSalary method for Contractor
func (c Contractor) CalculateSalary() float64 {
	return c.dailyRate * float64(c.daysWorked)
}

// CalculateSalary method for Freelancer
func (fl Freelancer) CalculateSalary() float64 {
	return fl.hourlyRate * float64(fl.hoursWorked)
}

func Qc() { //main function
	fullTimeEmployee := FullTime{dailyRate: 500, daysWorked: 30}
	contractor := Contractor{dailyRate: 100, daysWorked: 30}
	freelancer := Freelancer{hourlyRate: 100, hoursWorked: 20}

	employees := []Employee{fullTimeEmployee, contractor, freelancer}

	for _, employee := range employees {
		fmt.Printf("Salary: %.2f\n", employee.CalculateSalary())
	}
}
