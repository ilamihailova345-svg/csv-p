package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

//1
/*func main() {
	const filename = "users.csv"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Не удалось открыть файл %s", filename)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		log.Fatal("Ошибка чтения заголовка")
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Ошибка чтения csv файла")
	}

	for _, record := range records {
		fmt.Printf("ID: %s, Name: %s, Email: %s\n", record[0], record[1], record[2])
	}
}*/ 


//2
type Product struct {
	ID    int
	Name  string
	Price float64
}

/*func main() {
	products := []Product{
		{ID: 1, Name: "Серьги", Price: 12900},
		{ID: 2, Name: "Браслет", Price: 9000},
		{ID: 3, Name: "Ободок", Price: 1500},
	}

	fmt.Println("Запись данных в csv файл")

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatal("Не удалось создать csv файл")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Name", "Price"}
	err = writer.Write(header)
	if err != nil {
		log.Fatalf("Ошибка записи заголовка: %v", err)
	}

	for _, product := range products {
		id := strconv.Itoa(product.ID)
		price := strconv.FormatFloat(product.Price, 'f', 2, 64)

		record := []string{id, product.Name, price}
		err := writer.Write(record)
		if err != nil {
			log.Fatalf("Ошибка записи строки: %v", err)
		}
	}

	fmt.Println("Файл products.csv успешно создан.")
}*/


//3
func main() {
	//открываем файл
	file, err := os.Open("products.csv")
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Ошибка чтения данных из csv файла: %v", err)
	}

	loadedProducts := []Product{}

	// пропуск заголовка
	records = records[1:]

	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalf("Ошибка чтения ID %s: %v", record[0], err)
		}

		price, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatalf("Ошибка чтения цены %s: %v", record[2], err)
		}

		product := Product{
			ID:    id,
			Name:  record[1],
			Price: price,
		}
		loadedProducts = append(loadedProducts, product)
	}
	//вывод данных на экран	
	for _, product := range loadedProducts {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", product.ID, product.Name, product.Price)
	}
}