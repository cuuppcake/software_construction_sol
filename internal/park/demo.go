package park

import "fmt"

func showResult(name string, err error) {
	if err != nil {
		fmt.Println(name, "— ошибка:", err)
	} else {
		fmt.Println(name, "— успешно добавлено")
	}
}

func RunDemo(p *Park) {
	fmt.Println("ВышКат — учёт парка")
	fmt.Println("Добавление транспорта и техосмотр:")
	showResult("Самокат", p.AddTransport(NewScooter("Ninebot", 101, 0.5, 8)))
	showResult("Велосипед", p.AddTransport(NewBicycle("Forward", 102, 9)))
	showResult("Электровелосипед", p.AddTransport(NewElectricBicycle("Xiaomi", 103, 0.8, 7)))
	showResult("Неисправный самокат", p.AddTransport(NewScooter("Broken", 104, -0.4, 7)))
	showResult("Повторный номер", p.AddTransport(NewBicycle("Duplicate", 101, 7)))

	fmt.Println("\nДобавление вещей:")
	showResult("Шлем", p.AddItem(NewHelmet("Шлем", 201)))
	showResult("Зарядная станция", p.AddItem(NewChargingStation("Зарядная станция", 202, 1.2)))

	fmt.Println("\nКоличество транспорта:", p.TransportCount())
	fmt.Println("Количество вещей:", p.ItemCount())
	fmt.Printf("Суточное энергопотребление: %.2f кВт·ч\n", p.TotalEnergy())

	fmt.Println("\nПодходит новичкам:")
	for _, t := range p.ForBeginners() {
		fmt.Println(t.GetName(), "—", t.GetNumber())
	}

	fmt.Println("\nВесь инвентарь:")
	for _, item := range p.Inventory() {
		fmt.Println(item.GetName(), "—", item.GetNumber())
	}
}
