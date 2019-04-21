/*
 * @lc app=leetcode id=278 lang=java
 *
 * [278] First Bad Version
 *
 * https://leetcode.com/problems/first-bad-version/description/
 *
 * algorithms
 * Easy (28.57%)
 * Total Accepted:    214.3K
 * Total Submissions: 726.3K
 * Testcase Example:  '5\n4'
 *
 * You are a product manager and currently leading a team to develop a new
 * product. Unfortunately, the latest version of your product fails the quality
 * check. Since each version is developed based on the previous version, all
 * the versions after a bad version are also bad.
 * 
 * Suppose you have n versions [1, 2, ..., n] and you want to find out the
 * first bad one, which causes all the following ones to be bad.
 * 
 * You are given an API bool isBadVersion(version) which will return whether
 * version is bad. Implement a function to find the first bad version. You
 * should minimize the number of calls to the API.
 * 
 * Example:
 * 
 * 
 * Given n = 5, and version = 4 is the first bad version.
 * 
 * call isBadVersion(3) -> false
 * call isBadVersion(5) -> true
 * call isBadVersion(4) -> true
 * 
 * Then 4 is the first bad version. 
 * 
 */
/* The isBadVersion API is defined in the parent class VersionControl.
      boolean isBadVersion(int version); */
// References:
//	http://bit.ly/2GhVloe
//	http://bit.ly/2VdtEqk
//	http://bit.ly/2Ixh5jF
//	http://bit.ly/2IwVojI
//	http://bit.ly/2Xzn7Uv
//	http://bit.ly/2XtNOcS
public class Solution extends VersionControl {
    public int firstBadVersion(int n) {
        // Left included / Right not included will overflow when:
        // n = 2^31 - 1.
        // Thus, use Left included / Right included instead.
        int l = 1;
        int r = n;

        while (l <= r) {
            int m = l + (r - l) / 2;

            // Upper bound binary search:
            //  Try to find first index: m, which:
            //  versions[m] -> versions[n-1] are all bad versions.
		    //  i.e. m is the first index of bad versions.
            if (isBadVersion(m)) {
			    // Search left part.
                r = m - 1;
            } else {
                // Search right part.
                l = m + 1;
            }
        }

        // 'l' is the first index,
        //  which versions[l] -> versions[n-1] are all bad versions,
        return l;
    }
}
