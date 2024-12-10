#include <iostream>
#include <vector>
#include <limits.h>
using namespace std;

int second_largest_element(vector<int> &arr, int n) {
    int largest = INT_MIN;
    int slargest = INT_MIN;
    for (auto i : arr) {
        if (i > largest) {
            slargest = largest;
            largest = i;
        } else if (i > slargest and i != largest) {
            slargest = i;
        }
    }
    return slargest;
}

int main() {

    int n;
    cin >> n;
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    cout << second_largest_element(arr, n);

    return 0;
}
