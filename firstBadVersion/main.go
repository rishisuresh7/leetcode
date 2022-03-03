package main
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

 // This is just a mock implementaion of the function, should not be used for testing.
 func isBadVersion(version int) bool {
	 return false
 }

func firstBadVersion(n int) int {
    return binarySearch(n, 0, n)
}

func binarySearch(maxNum, left, right int) int {
    mid := left + (right-left)/2;
    if left < right {
        if isBadVersion(mid) {
            return binarySearch(mid, left, mid)
        } else {
            return binarySearch(mid, mid+1, right)
        }
    }

    return mid
}
