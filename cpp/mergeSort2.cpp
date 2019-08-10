/*
 * @Author: zhanghao 
 * @Date: 2018-11-28 12:38:12 
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-11-28 18:05:59
 */

#include <iostream>
#include<stdlib.h>
#include <algorithm>
#include<time.h>
using namespace std;
#define LEN 1000000
#define VALUE_LIMIT 10000000

class mergeSortType {
    public:
    int value;
    mergeSortType();
    ~mergeSortType();
    void operator=(const mergeSortType &m) { 
         this->value = m.value;
    }
};

mergeSortType::mergeSortType() {
    value = 0;
}

mergeSortType::~mergeSortType() {
    value = 0;
}

bool operator<(mergeSortType const &x, mergeSortType const &y) {
    if (x.value < y.value) {
        return true;
    }
    return false;
}

bool less_mergeSortType(const mergeSortType &m1, const mergeSortType &m2) {
        return m1.value < m2.value;
}

template<typename T>
void merge_my(T array[], int left, int center, int right) {
    T tempArray[right - left +1];
    int p = left;
    int q = center+1;
    int i = 0;
    while(p <= center&&q <= right) {
        if (array[p] < array[q]) {
            tempArray[i] = array[p];
            p++;
            i++;
        }else {
            tempArray[i] = array[q];
            q++;
            i++;
        }
    }
    while (p <= center) {
        tempArray[i] = array[p];
        p++;
        i++;
    }
    while (q <= right) {
        tempArray[i] = array[q];
        q++;
        i++;
    }

    for (i = 0;i < sizeof(tempArray)/sizeof(T);i++) {
        array[left+i] = tempArray[i];
    }
}

template<typename T>
void mergeSort(T array[], int left, int right) {
    if (left < right) {
        int center = (left+right)/2;
        mergeSort(array, left, center);
        mergeSort(array, center+1, right);
        merge_my<mergeSortType>(array, left, center, right);
    }
}

int main() {
    mergeSortType array[LEN];

    // 赋随机值
    for (int i =0;i<LEN;i++) {
        array[i].value = rand()%VALUE_LIMIT;
    }
    // STL排序
    clock_t startTime,endTime;
    startTime = clock();
    sort(array, array+LEN-1, less_mergeSortType);
    endTime = clock();
    cout << "STL Totle Time: " <<(double)(endTime - startTime) / CLOCKS_PER_SEC << "s" << endl;

    // 赋随机值
    for (int i =0;i<LEN;i++) {
        array[i].value = rand()%VALUE_LIMIT;
    }

    // 归并排序
    startTime = clock();
    mergeSort<mergeSortType>(array, 0, LEN-1);
    endTime = clock();
    cout << "归并排序 Totle Time: " <<(double)(endTime - startTime) / CLOCKS_PER_SEC << "s"<<endl;
    cout<<endl;
    cout<<"归并排序后的随机数组:";
    for (int i = 0; i<LEN;i++) {
        if (i%4 == 0) {
            cout<<endl;
        }
        cout<<"array["<<i<<"]{"<<array[i].value<<"}   ";
    }
    return 0;
}
