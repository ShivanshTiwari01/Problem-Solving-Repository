// Program: Find the Longest subarray with given sum K.
#include <iostream>
#include <vector>
#include <map>
#include <algorithm>
using namespace std;

// Approach one: Using Hash Map this approach works for positive and negatives both. Best approach for arrays containing positve, zeroes and negatives.
int getLongestSubarray(vector<int> &a, long long k) {
    int n = a.size(); // size of the array.
    map<long long, int> preSumMap;
    long long sum = 0;
    int maxLen = 0;

    for (int i = 0; i < n; i++) {
        //calculate the prefix sum till index i:
        sum += a[i];

        // if the sum = k, update the maxLen:
        if (sum == k) {
            maxLen = max(maxLen, i + 1);
        }

        // calculate the sum of remaining part i.e. x-k:
        long long rem = sum - k;

        //Calculate the length and update maxLen:
        if (preSumMap.find(rem) != preSumMap.end()) {
            int len = i - preSumMap[rem];
            maxLen = max(maxLen, len);
        }

        //Finally, update the map checking the conditions:
        if (preSumMap.find(sum) == preSumMap.end()) {
            preSumMap[sum] = i;
        }
    }
    return maxLen;
}

// Approach two: Using Two pointer, this approach works only for positives including zeroes, but is more optimised approach.
// Best approach for arrays containing only positives and zeroes.
int getLongestSubarray2(vector<int> &a, long long k) {
    int n = a.size();        // size of the array.
    int left = 0, right = 0; // 2 pointers
    long long sum = a[0];
    int maxLen = 0;

    while (right < n) {
        // if sum > k, reduce the subarray from left
        // until sum becomes less or equal to k:
        while (left <= right && sum > k) {
            sum -= a[left];
            left++;
        }

        // if sum = k, update the maxLen i.e. answer:
        if (sum == k) {
            maxLen = max(maxLen, right - left + 1);
        }

        // Move forward thw right pointer:
        right++;
        if (right < n)
            sum += a[right];
    }
    return maxLen;
}

int main() {

    vector<int> a = {2, 3, 5, 1, 9};
    long long k = 10;
    int len = getLongestSubarray(a, k);
    cout << "The length of the longest subarray is: " << len << "\n";

    return 0;
}