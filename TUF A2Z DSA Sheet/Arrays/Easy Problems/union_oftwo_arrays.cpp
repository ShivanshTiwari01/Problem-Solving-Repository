// Program: Find union of two arrays.

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

vector<int> arrays_union(vector<int> a, vector<int> b, int n, int m) {
    vector<int> Union;
    int i = 0, j = 0;
    while (i < n and j < m) {
        if (a[i] <= b[j]) {
            if (Union.size() == 0 or Union.back() != a[i]) {
                Union.push_back(a[i]);
            }
            i++;
        } else {
            if (Union.size() == 0 or Union.back() != b[j]) {
                Union.push_back(b[j]);
            }
            j++;
        }
    }
    while (i < n) {
        if (Union.back() != a[i])
            Union.push_back(a[i]);
        i++;
    }
    while (j < m) {
        if (Union.back() != b[j])
            Union.push_back(b[j]);
        j++;
    }
    return Union;
}

int main() {

    int n, m;
    cin >> n >> m;
    vector<int> a(n);
    vector<int> b(m);

    cout << "enter elements of first array: ";
    for (int i = 0; i < n; i++) {
        cin >> a[i];
    }
    cout << "enter elements of second array: ";
    for (int i = 0; i < m; i++) {
        cin >> b[i];
    }

    vector<int> Union = arrays_union(a, b, n, m);

    for (int i : Union) {
        cout << i << " ";
    }

    return 0;
}
