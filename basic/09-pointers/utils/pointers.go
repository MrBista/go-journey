package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallPointerarnSection() {
	helper.CalledFunction("CallPointerarnSection")
	basicPointerConcepts()
	pointerDeclaration()
	functionParameters()
	pointerWithStruct()
	pointerWithSliceMap()
	linkedListExample()
	commonMistakes()

	fmt.Println("=== RINGKASAN ===")
	fmt.Println("• & = mengambil alamat memori (address-of operator)")
	fmt.Println("• * = mengakses nilai dari alamat (dereference operator)")
	fmt.Println("• Pointer = variable yang menyimpan alamat memori")
	fmt.Println("• Gunakan pointer untuk efisiensi dan mengubah nilai asli")
	fmt.Println("• Selalu cek nil sebelum dereference pointer")
	fmt.Println("• Slice dan Map sudah reference type secara default")
}

// ====================
// 1. KONSEP DASAR POINTER
// ====================

func basicPointerConcepts() {
	fmt.Println("=== KONSEP DASAR POINTER ===")

	// Variable biasa menyimpan nilai
	x := 42
	fmt.Printf("Nilai x: %d\n", x)
	fmt.Printf("Alamat memori x: %p\n", &x) // &x mengambil alamat memori

	// Pointer menyimpan alamat memori
	var ptr *int // Deklarasi pointer ke int
	ptr = &x     // ptr sekarang menunjuk ke alamat x

	fmt.Printf("Nilai ptr (alamat yang ditunjuk): %p\n", ptr)
	fmt.Printf("Nilai yang ditunjuk ptr: %d\n", *ptr) // *ptr mengakses nilai

	// Mengubah nilai melalui pointer
	*ptr = 100
	fmt.Printf("Nilai x setelah diubah via pointer: %d\n", x)
	fmt.Println()
}

// ====================
// 2. DEKLARASI POINTER
// ====================

func pointerDeclaration() {
	fmt.Println("=== DEKLARASI POINTER ===")

	// Cara 1: Deklarasi eksplisit
	var ptr1 *int
	fmt.Printf("Pointer kosong: %v\n", ptr1) // nil

	// Cara 2: Langsung assign
	num := 25
	ptr2 := &num
	fmt.Printf("ptr2 menunjuk ke: %d\n", *ptr2)

	// Cara 3: Menggunakan new()
	ptr3 := new(int) // Membuat pointer ke int dengan nilai zero (0)
	fmt.Printf("Nilai default dari new(int): %d\n", *ptr3)
	*ptr3 = 50
	fmt.Printf("Setelah diubah: %d\n", *ptr3)

	// Pointer ke berbagai tipe data
	str := "Hello"
	ptrStr := &str
	fmt.Printf("String via pointer: %s\n", *ptrStr)

	isTrue := true
	ptrBool := &isTrue
	fmt.Printf("Boolean via pointer: %t\n", *ptrBool)
	fmt.Println()
}

// ====================
// 3. POINTER SEBAGAI PARAMETER FUNGSI
// ====================

// Pass by value (copy nilai)
func changeValueCopy(x int) {
	x = 999
	fmt.Printf("Dalam fungsi changeValueCopy: %d\n", x)
}

// Pass by pointer (reference ke alamat memori)
func changeValuePointer(x *int) {
	*x = 999
	fmt.Printf("Dalam fungsi changeValuePointer: %d\n", *x)
}

func functionParameters() {
	fmt.Println("=== POINTER SEBAGAI PARAMETER ===")

	num := 42
	fmt.Printf("Nilai awal: %d\n", num)

	// Pass by value - tidak mengubah nilai asli
	changeValueCopy(num)
	fmt.Printf("Setelah changeValueCopy: %d\n", num)

	// Pass by pointer - mengubah nilai asli
	changeValuePointer(&num)
	fmt.Printf("Setelah changeValuePointer: %d\n", num)
	fmt.Println()
}

// ====================
// 4. POINTER DAN STRUCT
// ====================

type Person struct {
	Name string
	Age  int
}

// Method dengan receiver pointer
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// Method dengan receiver value
func (p Person) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func pointerWithStruct() {
	fmt.Println("=== POINTER DAN STRUCT ===")

	// Membuat struct
	person1 := Person{Name: "Alice", Age: 25}
	fmt.Printf("person1: %+v\n", person1)

	// Pointer ke struct
	personPtr := &person1
	fmt.Printf("Akses via pointer: %+v\n", *personPtr)

	// Akses field via pointer (Go otomatis dereference)
	fmt.Printf("Name via pointer: %s\n", personPtr.Name)
	personPtr.Age = 30
	fmt.Printf("Setelah ubah age: %+v\n", person1)

	// Method dengan pointer receiver
	person1.UpdateAge(35)
	fmt.Printf("Setelah UpdateAge: %s\n", person1.GetInfo())
	fmt.Println()
}

// ====================
// 5. POINTER DAN SLICE/MAP
// ====================

func modifySlice(s []int) {
	s[0] = 999 // Slice adalah reference type, jadi ini mengubah aslinya
}

func modifySlicePointer(s *[]int) {
	*s = append(*s, 100) // Mengubah slice itu sendiri
}

func pointerWithSliceMap() {
	fmt.Println("=== POINTER DAN SLICE/MAP ===")

	// Slice (reference type)
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice awal: %v\n", numbers)

	modifySlice(numbers)
	fmt.Printf("Setelah modifySlice: %v\n", numbers)

	modifySlicePointer(&numbers)
	fmt.Printf("Setelah modifySlicePointer: %v\n", numbers)

	// Map (juga reference type)
	ages := map[string]int{"Alice": 25, "Bob": 30}
	fmt.Printf("Map awal: %v\n", ages)

	// Map secara default sudah reference
	updateMap(ages)
	fmt.Printf("Setelah updateMap: %v\n", ages)
	fmt.Println()
}

func updateMap(m map[string]int) {
	m["Charlie"] = 35
}

// ====================
// 6. CONTOH PRAKTIS: LINKED LIST
// ====================

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Add(data int) {
	newNode := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (ll *LinkedList) Display() {
	current := ll.Head
	fmt.Print("LinkedList: ")
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func linkedListExample() {
	fmt.Println("=== CONTOH PRAKTIS: LINKED LIST ===")

	list := &LinkedList{}
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Display()
	fmt.Println()
}

// ====================
// 7. COMMON MISTAKES DAN BEST PRACTICES
// ====================

func commonMistakes() {
	fmt.Println("=== KESALAHAN UMUM ===")

	// Mistake 1: Nil pointer dereference
	var ptr *int
	// fmt.Println(*ptr) // PANIC! Cannot dereference nil pointer

	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("Pointer is nil")
	}

	// Mistake 2: Pointer ke local variable
	// Ini sebenarnya OK di Go karena ada escape analysis
	ptr = createPointer()
	fmt.Printf("Pointer dari fungsi: %d\n", *ptr)

	// Best Practice: Selalu cek nil sebelum dereference
	safePointerAccess(ptr)
	safePointerAccess(nil)
	fmt.Println()
}

func createPointer() *int {
	x := 42
	return &x // OK di Go, variable akan di-allocate di heap
}

func safePointerAccess(ptr *int) {
	if ptr != nil {
		fmt.Printf("Safe access: %d\n", *ptr)
	} else {
		fmt.Println("Cannot access nil pointer")
	}
}
