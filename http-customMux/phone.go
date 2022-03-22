package main

var phones = []*Phone{}

type Phone struct {
	Id     string
	Name   string
	Vendor string
	Year   int32
}

func init() {
	phones = append(phones, &Phone{Id: "P001", Name: "Redmi Note 9 Pro", Vendor: "Redmi", Year: 2020})
	phones = append(phones, &Phone{Id: "P002", Name: "Realme 6 Pro", Vendor: "Realme", Year: 2020})
	phones = append(phones, &Phone{Id: "P003", Name: "Samsung S22 Ultra", Vendor: "Samsung", Year: 2022})
	phones = append(phones, &Phone{Id: "P004", Name: "iPhone 13 Pro", Vendor: "Apple", Year: 2022})
}

func GetPhones() []*Phone {
	return phones
}

func SelectPhone(id string) *Phone {
	for _, each := range phones {
		if each.Id == id {
			return each
		}
	}

	return nil
}
