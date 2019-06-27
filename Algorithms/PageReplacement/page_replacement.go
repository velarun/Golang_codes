package main

import "fmt"
import "sort"

func fifo(a []int, n, m int) {

	f := -1
	page_faults := 0

	page := make([]int, m)
	for i:=0; i<m; i++ {
		page[i] = -1
	}

	for i:=0; i<n; i++ {
		flag := 0
		for j:=0; j<m; j++ {
			if page[j] == a[i] {
				flag = 1
				break
			}
		}

		if flag == 0 {
			f = (f+1)%m
			page[f] = a[i]
			page_faults+=1
			for j:=0; j<m; j++ {
				if page[j] != -1 {
					fmt.Print(page[j], " ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Print("-> ", a[i], " Page Fault ", page_faults)
		} else {
			for j:=0; j<m; j++ {
				if page[j] != -1 {
					fmt.Print(page[j], " ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Print("-> ", a[i], " No Page Fault")
		}
		fmt.Println()
	}			
	fmt.Println("Total page faults : ", page_faults)
}

func lru(a []int, n, m int) {
    x := 0
    page_faults := 0
    page := make([]int, m)
	for i:=0; i<m; i++ {
		page[i] = -1
	}

    for i:=0; i<n; i++ {
		flag := 0
		for j:=0; j<m; j++ {
			if page[j] == a[i] {
				flag = 1
				break
			}
		}

		if flag == 0 {
            if page[x] != -1 {
                min := 999
                for k:=0; k<m; k++ {
                    flag := 0
                    j :=  i
                    for j>=0 {
                        j-=1
                        if page[k] == a[j] {
                            flag = 1
							break
						}
					}

                    if flag == 1 && min > j {
                        min = j
						x = k
					}
				}
			}

            page[x] = a[i]
            x=(x+1)%m
            page_faults+=1
			for j:=0; j<m; j++ {
				if page[j] != -1 {
					fmt.Print(page[j], " ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Print("-> ", a[i], " Page Fault ", page_faults)
		} else {
			for j:=0; j<m; j++ {
				if page[j] != -1 {
					fmt.Print(page[j], " ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Print("-> ", a[i], " No Page Fault")
		}
		fmt.Println()
	}			
	fmt.Println("Total page faults : ", page_faults)
}

func min(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}

	return sum
}

func lfu(a []int, n, m int) {
	x := 0
	page_faults := 0
	time := make(map[int]int)
	var b []int
	copy(b, a[:])
	
	page := make([]int, m)
	for i:=0; i<m; i++ {
		page[i] = -1
	}

	for _, v := range a {
		time[v] = 0
	}

	for i:=0; i<n; i++ {
		flag := 0
		for j:=0; j<m; j++ {
			if page[j] == a[i] {
				flag = 1
				time[a[i]]+=1
				break
			}
		}
			
		if flag == 0 {
			rpage := -1
			if page[x] != -1 {
				var t []int
				for _, k := range page {
					t = append(t, time[k])
				}
				
				minis := min(t)
				
				var gpage []int
				for _, k := range page {
					if time[k] == minis {
						gpage = append(gpage, k)	
					}
				}	
						
				maxi := -1
				flag = 0
				
				for _, k := range gpage {
					for n:=0; n<i; n++ {
						if k == b[n] {
							if maxi == -1 {
								maxi = n
								rpage = k
							} else if n < maxi {
								maxi = n 
								rpage = k	
							}
						}
					}
					flag = 1
				}
				
				if flag == 1 {
					b[maxi]=-1
					x = sort.SearchInts(page, rpage)
				}
			}
						
			if rpage != -1 {
				time[rpage]=0
			}
				
			page[x] = a[i]
			x=(x+1)%m
			time[a[i]]+=1
			page_faults+=1
			for j:=0; j<m; j++ {
				if page[j] != -1 {
					fmt.Print(page[j], " ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Print("-> ", a[i], " Page Fault ", page_faults)
		} else {
			for j:=0; j<m; j++ {
				if page[j] != -1 {
					fmt.Print(page[j], " ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Print("-> ", a[i], " No Page Fault")
		}
		fmt.Println()
	}			

	fmt.Println("Total page faults : ", page_faults)
	fmt.Println("Frequencies of pages: ", time)
}

func optimal(a []int, n, m int) {
    page_faults := 0
    FREE := -1
	
	page := make([]int, m)
	for i:=0; i<m; i++ {
		page[i] = FREE
	}

    for i:=0; i<n; i++ {
        flag := 0
        for j:=0; j<m; j++ {
            if page[j] == a[i] {
                flag = 1
				break
			}
		}
            
        if flag == 0 {
            faulted := false
            new_slot := FREE
            for q:=0; q<m; q++ {
                if page[q] == FREE {
                    faulted = true
					new_slot = q
				}
			}
            
            if !faulted {
                max_future := 0
                max_future_q := FREE
                for q:=0; q<m; q++ {
                    if page[q] != FREE {
                        found := false
                        for ii:=i; ii<n; ii++ {
                            if a[ii] == page[q] {
                                found = true
                                if ii > max_future {
                                    max_future = ii
                                    max_future_q = q
								}
								break
							}
						}
                        
                        if !found {
                            max_future_q = q
							break
						}
					}
				}
                faulted = true
				new_slot = max_future_q
			}
            
			page_faults += 1
			page[new_slot] = a[i]
			for j:=m-1; j>-1; j-- {
				if page[j] != FREE {
					fmt.Print(page[j], " ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Print("-> ", a[i], " Page Fault ", page_faults)
		} else {
			for j:=m-1; j>-1; j-- {
				if page[j] != FREE {
					fmt.Print(page[j], " ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Print("-> ", a[i], " No Page Fault")
		}
		fmt.Println()
	}			

	fmt.Println("Total page faults : ", page_faults)
}

func main() {
	pages := []int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2}
	n := len(pages) 
	capacity := 4
	fmt.Println("===============================")
	fmt.Println("FIFO PAGE REPLACEMENT ALGORITHM")
	fmt.Println("===============================")
	fifo(pages, n, capacity) 
	fmt.Println("===============================")
	fmt.Println("LRU PAGE REPLACEMENT ALGORITHM")
	fmt.Println("===============================")
	lru(pages, n, capacity) 
	fmt.Println("===============================")
	fmt.Println("LFU PAGE REPLACEMENT ALGORITHM")
	fmt.Println("===============================")
	lfu(pages, n, capacity)
	fmt.Println("===============================")
	fmt.Println("OPTIMAL PAGE REPLACEMENT ALGORITHM")
	fmt.Println("===============================")
	optimal(pages, n, capacity)
}