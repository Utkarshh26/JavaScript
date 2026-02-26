package binarysearch

func SearchInts(list []int, key int) int {
	st := 0 
    end := len(list) - 1
    
    for st <= end {
        mid := st + (end - st)/2
        if key < list[mid] {
            end = mid - 1
        }else if key > list[mid] {
            st = mid + 1
        }else {
            return mid
        }
    }
    
    return -1
}
