package main

import (
	"fmt"
)

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func reverse(nums []int, l, r int) {
	for i, j := l, r; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

}

func binary_search(nums []int, l, r, target int) int {
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	if nums[l] == target && nums[r] == target {
		return []int{l, r}
	}
	search := binary_search(nums, 0, len(nums)-1, target)
	if search == -1 {
		return []int{-1, -1}
	}
	ans_l, ans_r := l, r
	pix := search
	for l <= pix {
		mid := (l + pix) / 2
		if mid == l {
			if nums[mid] == target {
				ans_l = mid
				break
			} else {
				l = mid + 1
				continue
			}
		}
		if nums[mid] == target && nums[mid-1] != target {
			ans_l = mid
			break
		} else if nums[mid] != target {
			l = mid + 1
		} else {
			pix = mid - 1
		}
	}
	pix = search
	for pix <= r {
		mid := (pix + r) / 2
		if mid == r {
			ans_r = mid
			break
		}
		if nums[mid] == target && nums[mid+1] != target {
			ans_r = mid
			break
		} else if nums[mid] != target {
			r = mid - 1
		} else {
			pix = mid + 1
		}
	}
	return []int{ans_l, ans_r}
}

func Main() {
	var nums []int
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		nums = append(nums, tmp)
	}
	var target int
	fmt.Scan(&target)
	ans := searchRange(nums, target)
	fmt.Println(ans)
}
