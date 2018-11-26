/*
 * @Author: zhanghao 
 * @Date: 2018-11-20 15:29:07 
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-11-27 02:15:51
 */

#include <iostream>
#include<stdlib.h>
using namespace std;

void mergeSort(int array[], int left, int right);

void merge(int array[], int left, int center, int right);

int main() {
    int array[1000];
    for (int i =0;i<1000;i++) {
        array[i] = rand()%10000;
    }

    int len = sizeof(array)/sizeof(int);
    mergeSort(array, 0, len-1);
    cout<<"最终输出："<<endl;
    for (int i = 0; i<len;i++) {
        std::cout<<"index "<<i<<":"<<array[i]<<std::endl;
    }
    cout<<"--------------------------"<<endl;
    return 0;
}

void mergeSort(int array[], int left, int right) {
    if (left < right) {
        int center = (left+right)/2;
        mergeSort(array, left, center);
        mergeSort(array, center+1, right);
        merge(array, left, center, right);
    }
}

void merge(int array[], int left, int center, int right) {
    int tempArray[right - left +1];
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

    for (i = 0;i < sizeof(tempArray)/sizeof(int);i++) {
        array[left+i] = tempArray[i];
    }
}
