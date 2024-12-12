// Program: Find the number of maximum consecutive ones that appear in given array;

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int max_consecutive_ones(vector<int> arr) {
    int count = 0;
    int max_count = 0;
    for (int i : arr) {
        if (i == 1)
            count++;
        else
            count = 0;
        max_count = max(max_count, count);
    }
    return max_count;
}

int main() {

    int n;
    cin >> n;
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    cout << max_consecutive_ones(arr) << endl;

    return 0;
}