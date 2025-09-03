package main

import(
	"fmt"
)

type Employee interface{
	CalculateSalary() float64
}

type FullTime struct{
    MonthlySalary float64
}

func (e FullTime) CalculateSalary() float64 {
	return e.MonthlySalary
}

type Contractor struct{
	MonthlySalary float64
}

func(e Contractor) CalculateSalary() float64{
	return e.MonthlySalary
}

type Freelancer struct{
	HourlyRate, HoursWorked float64
}

func(e Freelancer) CalculateSalary() float64{
	return e.HourlyRate * e.HoursWorked
}

func main(){
	fullTimeEmployee := FullTime{MonthlySalary: 15000}
	contractor := Contractor{MonthlySalary: 3000}
	freelancer := Freelancer{HourlyRate: 100, HoursWorked: 20}

	employee := []Employee{
        fullTimeEmployee,
		contractor,
		freelancer,
	}

	for _, emp := range employee{
		fmt.Printf("Employee's salary: %.2f\n", emp.CalculateSalary())
	}
}
