// Program: Find the missing number in array.

#include <iostream>
#include <vector>
using namespace std;

int missingNumber(vector<int> &nums) {
    int s1 = 0;
    for (int i : nums) {
        s1 += i;
    }
    int n = nums.size();
    int s2 = (n * (n + 1)) / 2;
    return s2 - s1;
}

int main() {

    int n;
    cin >> n;
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    cout << missingNumber(arr);

    return 0;
}