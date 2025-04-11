package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Employee struct to hold employee data
type Employee struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	HourlyRate    float64 `json:"hourly_rate"`
	HoursWorked   float64 `json:"hours_worked"`
	MonthlySalary float64 `json:"monthly_salary"`
}

// Global slice to hold all employees
var employees []Employee

// Function to initialize the system
func initializeSystem() {
	employees = []Employee{}

	if _, err := os.Stat("employees.json"); !os.IsNotExist(err) {
		if err := loadDataFromFile(); err != nil {
			fmt.Println("Error loading data:", err)
		}
	}

	fmt.Println("System initialized.")
}

// Function to add a new employee
func addEmployee(id int, name string, hourlyRate float64) {
	employee := Employee{
		ID:         id,
		Name:       name,
		HourlyRate: hourlyRate,
	}
	employees = append(employees, employee)
	fmt.Printf("Added employee: ID=%d, Name=%s, HourlyRate=%.2f\n", id, name, hourlyRate)
}

// Function to enter hours worked for an employee
func enterHoursWorked(id int, hours float64) error {
	for i := range employees {
		if employees[i].ID == id {
			employees[i].HoursWorked = hours
			employees[i].MonthlySalary = employees[i].HourlyRate * hours
			fmt.Printf("Updated hours for employee ID=%d: HoursWorked=%.2f\n", id, hours)
			return nil
		}
	}
	return fmt.Errorf("employee with ID=%d not found", id)
}

// Function to calculate salaries for all employees
func calculateSalaries() {
	for i := range employees {
		employees[i].MonthlySalary = employees[i].HourlyRate * employees[i].HoursWorked
	}
	fmt.Println("Salaries calculated.")
}

// Function to generate a payroll report
func generatePayrollReport() {
	fmt.Printf("%-10s %-20s %-15s %-15s\n", "ID", "Name", "Hours Worked", "Monthly Salary")
	fmt.Println("---------------------------------------------------------")
	for _, emp := range employees {
		fmt.Printf("%-10d %-20s %-15.2f $%-14.2f\n", emp.ID, emp.Name, emp.HoursWorked, emp.MonthlySalary)
	}
}

// Function to save data to a file
func saveDataToFile() error {
	file, err := os.Create("employees.json")
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(employees); err != nil {
		return fmt.Errorf("error encoding data to file: %v", err)
	}

	return nil
}

// Function to load data from a file
func loadDataFromFile() error {
	file, err := os.Open("employees.json")
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&employees); err != nil {
		return fmt.Errorf("error decoding data from file: %v", err)
	}

	return nil
}

// Function to handle user input and execute corresponding actions
func handleUserInput(scanner *bufio.Scanner) {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Enter Hours Worked")
		fmt.Println("3. Calculate Salaries")
		fmt.Println("4. Generate Payroll Report")
		fmt.Println("5. Exit")

		fmt.Print("Enter choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter Employee ID: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid ID. Please enter a numeric value.")
				continue
			}

			fmt.Print("Enter Employee Name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter Hourly Rate: ")
			scanner.Scan()
			rate, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Invalid Hourly Rate. Please enter a numeric value.")
				continue
			}

			addEmployee(id, name, rate)

		case "2":
			fmt.Print("Enter Employee ID: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid ID. Please enter a numeric value.")
				continue
			}

			fmt.Print("Enter Hours Worked: ")
			scanner.Scan()
			hours, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Invalid Hours Worked. Please enter a numeric value.")
				continue
			}

			if err := enterHoursWorked(id, hours); err != nil {
				fmt.Println(err)
			}

		case "3":
			calculateSalaries()

		case "4":
			generatePayrollReport()

		case "5":
			if err := saveDataToFile(); err != nil {
				fmt.Println("Error saving data:", err)
			}
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func main() {
	initializeSystem()

	scanner := bufio.NewScanner(os.Stdin)
	handleUserInput(scanner)
}
