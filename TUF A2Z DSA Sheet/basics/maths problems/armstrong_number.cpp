// Program: Check if the given number is Armstrong or not. 
// Armstrong numbers are the numbers whose sum of cubes of each digit is equal to the number itself eg 153 and 1634.

#include <iostream>
#include <vector>
using namespace std;

int power(int x, unsigned int y)
{
    if (y == 0)
        return 1;
    if (y % 2 == 0)
        return power(x, y / 2) * power(x, y / 2);
    return x * power(x, y / 2) * power(x, y / 2);
}

string check_armstrong(int n) {
    
    int digits = 0;
    
    int temp = n;
    while(temp > 0) {
        digits += 1;
        temp /= 10;
    }
    temp = n;
    int arm = 0;
    
    while(temp > 0) {
        int num = temp%10;
        arm += power(num, digits);
        temp /= 10;
    }
    return arm == n ? "true" : "false";
}

int main() {
    int n;
    cin >> n; 
    cout << "Given no is armstrong: " << check_armstrong(n);
    return 0;
}
