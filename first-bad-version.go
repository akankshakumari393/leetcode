/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
 	lo, hi := 1, n
	for lo < hi {
        if mid := (lo+hi)/2; isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid+1
		}
	}
	
	return lo   
}
