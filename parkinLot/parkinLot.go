package main

import (
	"fmt"
	"time"
)

type Vehicle struct {
	Type string
	Size int
}

func NewVehicle(typ string, size int) *Vehicle {
	return &Vehicle{Type: typ, Size: size}
}

func (v *Vehicle) GetType() string {
	return v.Type
}

func (v *Vehicle) GetSize() int {
	return v.Size
}

func (v *Vehicle) SetType(typ string) {
	v.Type = typ
}

func (v *Vehicle) SetSize(size int) {
	v.Size = size
}

type Driver struct {
	Id        int
	Vehicle   *Vehicle
	DueAmount float64
}

func NewDriver(id int, vehicle *Vehicle) *Driver {
	return &Driver{Id: id, Vehicle: vehicle, DueAmount: 0}
}

func (d *Driver) GetId() int {
	return d.Id
}

func (d *Driver) GetVehicle() *Vehicle {
	return d.Vehicle
}

func (d *Driver) GetDueAmount() float64 {
	return d.DueAmount
}

func (d *Driver) SetId(id int) {
	d.Id = id
}

func (d *Driver) SetVehicle(vehicle *Vehicle) {
	d.Vehicle = vehicle
}

func (d *Driver) SetDueAmount(amount float64) {
	d.DueAmount = amount
}

type ParkingSystem struct {
	Floors                  int
	VehicleSpotsPerFloor    map[string]int
	Spots                   []map[string][]int
	HourlyRate              int
	VehicleParkedTimeRecord map[int]int
	DriverDueAmountRecord   map[int]float64
}

func NewParkingSystem(floors int, vehicleSpotsPerFloor map[string]int, hourlyRate int) *ParkingSystem {
	spots := make([]map[string][]int, floors)
	vehicleParkedRecord := make(map[int]int, 0)
	amountRecord := make(map[int]float64, 0)

	vehicleSpots := make(map[string][]int, 0)
	for v, i := range vehicleSpotsPerFloor {
		vehicleSpots[v] = make([]int, i)
	}

	for i := 0; i < floors; i++ {
		spots[i] = vehicleSpots
	}

	return &ParkingSystem{
		Floors:                  floors,
		VehicleSpotsPerFloor:    vehicleSpotsPerFloor,
		Spots:                   spots,
		HourlyRate:              hourlyRate,
		VehicleParkedTimeRecord: vehicleParkedRecord,
		DriverDueAmountRecord:   amountRecord,
	}
}

func (p *ParkingSystem) GetFloors() int {
	return p.Floors
}

func (p *ParkingSystem) GetVehicleSpotsPerFloor() map[string]int {
	return p.VehicleSpotsPerFloor
}

func (p *ParkingSystem) GetSpots() []map[string][]int {
	return p.Spots
}

func (p *ParkingSystem) GetHourlyRate() int {
	return p.HourlyRate
}

func (p *ParkingSystem) GetVehicleParkedTimeRecord() map[int]int {
	return p.VehicleParkedTimeRecord
}

func (p *ParkingSystem) GetDriverDueAmountRecord() map[int]float64 {
	return p.DriverDueAmountRecord
}

func (p *ParkingSystem) SetFloors(floors int) {
	p.Floors = floors
}

func (p *ParkingSystem) SetVehicleSpotsPerFloor(vehicleSpotsPerFloor map[string]int) {
	p.VehicleSpotsPerFloor = vehicleSpotsPerFloor
}

func (p *ParkingSystem) SetSpots(spots []map[string][]int) {
	p.Spots = spots
}

func (p *ParkingSystem) SetHourlyRate(rate int) {
	p.HourlyRate = rate
}

func (p *ParkingSystem) SetVehicleParkedTimeRecord(record map[int]int) {
	p.VehicleParkedTimeRecord = record
}

func (p *ParkingSystem) SetDriverDueAmountRecord(record map[int]float64) {
	p.DriverDueAmountRecord = record
}

func (p *ParkingSystem) VehicleEntry(driver *Driver) bool {
	spots := p.GetSpots()
	parkedRecord := p.GetVehicleParkedTimeRecord()
	for sIndex, s := range p.Spots {
		for vIndex, i := range s[driver.GetVehicle().GetType()] {
			if i == 0 {
				spots[sIndex][driver.GetVehicle().GetType()][vIndex] = driver.GetId()
				p.SetSpots(spots)

				parkedRecord[driver.GetId()] = int(time.Now().Unix())
				p.SetVehicleParkedTimeRecord(parkedRecord)
				return true
			}
		}
	}

	return false
}

func (p *ParkingSystem) VehicleExit(driver *Driver) bool {
	spots := p.GetSpots()
	parkedRecord := p.GetVehicleParkedTimeRecord()
	amountRecord := p.GetDriverDueAmountRecord()
	driverDueAmount := driver.GetDueAmount()

	totalHours := 0.0
	parkingAmount := 0.0

	for sIndex, s := range p.Spots {
		for vIndex, i := range s[driver.GetVehicle().GetType()] {
			if i == driver.GetId() {
				spots[sIndex][driver.GetVehicle().GetType()][vIndex] = 0
				p.SetSpots(spots)

				totalHours = time.Now().Sub(time.Unix(int64(parkedRecord[driver.GetId()]), 0)).Hours()
				parkingAmount = float64(p.GetHourlyRate()) * totalHours

				driver.SetDueAmount(driverDueAmount + parkingAmount)
				amountRecord[driver.GetId()] = driver.GetDueAmount()
				p.SetDriverDueAmountRecord(amountRecord)

				delete(parkedRecord, driver.GetId())
				p.SetVehicleParkedTimeRecord(parkedRecord)

				return true
			}
		}
	}

	return false
}

func main() {
	car := NewVehicle("Car", 1)
	limo := NewVehicle("Limo", 2)
	truck := NewVehicle("Truck", 3)

	gagan := NewDriver(1, car)
	prerit := NewDriver(2, limo)
	atharva := NewDriver(3, truck)

	vehicleSpotsPerFloor := map[string]int{car.GetType(): 1, limo.GetType(): 1, truck.GetType(): 0}
	system := NewParkingSystem(3, vehicleSpotsPerFloor, 5)

	fmt.Printf("Gagan got entry - %v\n", system.VehicleEntry(gagan))
	fmt.Printf("Athrava got entry - %v\n", system.VehicleEntry(atharva))
	fmt.Printf("Prerit got entry - %v\n", system.VehicleEntry(prerit))

	time.Sleep(60 * time.Second)
	fmt.Printf("Prerit removed vehicle - %v\n", system.VehicleExit(prerit))
	time.Sleep(60 * time.Second)
	fmt.Printf("Athrava removed vehicle - %v\n", system.VehicleExit(atharva))
	time.Sleep(60 * time.Second)
	fmt.Printf("Gagan removed vehicle  - %v\n", system.VehicleExit(gagan))

	fmt.Printf("Prerit due  - %v\n", prerit.GetDueAmount())
	fmt.Printf("Athrava due - %v\n", atharva.GetDueAmount())
	fmt.Printf("Gagan due  - %v\n", gagan.GetDueAmount())

}
