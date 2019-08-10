/*
 * @Author: zhanghao 
 * @Date: 2018-11-21 17:38:28 
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-11-24 00:55:30
 */

#include<iostream>
#include<string>
#include<sstream>
using namespace std;

string getBracketFormat(int* array, int numMatrix);
string int2str(int &i);
void printInOtherFormat(int* array, int numMatrix);

int main() {
    // 矩阵A(n)的行为array[n-1],列为array[n],其中n>=1
    int* array = new int[1];
    cout<<"输入A1的行数(<1结束输入):";
    cin>>array[0];
    int i = 0;
    while (array[i] > 0) {
        i++;
        cout<<"输入A"<<i<<"的列数(<1结束输入):";
        cin>>array[i];
    }
    string result = getBracketFormat(array, i-1);
    cout<<result<<endl;
    cout<<"other format:"<<endl;
    printInOtherFormat(array, i-1);
    return 0;
}

void MatrixChain(int* p, int n, int** m, int** s);
string Traceback(int i, int j, int** s);

string getBracketFormat(int* array, int numMatrix){
    int** minMulTimes = new int*[numMatrix+1];
    for (int i=0;i<=numMatrix;i++) {
        minMulTimes[i] = new int[numMatrix+1];
    }
    for (int i=1;i<=numMatrix;i++) {
        for (int j=i;j<=numMatrix;j++) {
            minMulTimes[i][j] = 0;
        }
    }
    int** bestSplit = new int*[numMatrix+1];
    for (int i=0;i<=numMatrix;i++) {
        bestSplit[i] = new int[numMatrix+1];
    }    
    for (int i=1;i<numMatrix;i++) {
        for (int j=i;j<numMatrix;j++) {
            bestSplit[i][j] = 0;
        }
    }
    MatrixChain(array, numMatrix, minMulTimes, bestSplit);
    cout<<"最佳计算次序需要"<<minMulTimes[1][numMatrix]<<"次乘法运算"<<endl;
    bool* bracket = new bool[numMatrix+1];
    for (int i=1; i < numMatrix+1; i++) {
        bracket[i] = false;
    }
    return Traceback(1, numMatrix, bestSplit);
}

void Traceback2(int i, int j, int** s);

void printInOtherFormat(int* array, int numMatrix) {
    int** minMulTimes = new int*[numMatrix+1];
    for (int i=0;i<=numMatrix;i++) {
        minMulTimes[i] = new int[numMatrix+1];
    }
    for (int i=1;i<numMatrix;i++) {
        for (int j=i;j<numMatrix;j++) {
            minMulTimes[i][j] = 0;
        }
    }
    int** bestSplit = new int*[numMatrix+1];
    for (int i=0;i<=numMatrix;i++) {
        bestSplit[i] = new int[numMatrix+1];
    }
    for (int i=1;i<numMatrix;i++) {
        for (int j=i;j<numMatrix;j++) {
            bestSplit[i][j] = 0;
        }
    }
    MatrixChain(array, numMatrix, minMulTimes, bestSplit);
    cout<<"最佳计算次序需要"<<minMulTimes[1][numMatrix]<<"次乘法运算"<<endl;
    Traceback2(1,numMatrix,bestSplit);
}

void MatrixChain(int* p, int n, int** m, int** s) {
    for (int i =1; i<=n;i++) {
        m[i][i] = 0;
    }
    for (int r =2; r<=n;r++) {
        for (int i =1;i<=n-r+1;i++){
            int j = i+r-1;
            m[i][j] = m[i+1][j]+p[i-1]*p[i]*p[j];
            s[i][j] = i;
            for (int k = i+1;k<j;k++) {
                int t = m[i][k] + m[k+1][j]+p[i-1]*p[k]*p[j];
                if (t< m[i][j]) {
                    m[i][j] = t;
                    s[i][j] = k;
                }
            }
        }
    }
}

string Traceback(int i, int j, int** s) {
    if (i == j) {
        return "A"+int2str(i);
    }
    if (i > j) {
        cout<<"错误: i>j"<<endl;
    }
    return "("+Traceback(i, s[i][j], s)+Traceback(s[i][j]+1,j,s)+")";
}

// string getBracketString(int numMatrix, bool* bracket){
//     string result = "(";
//     for (int i=1;i<=numMatrix;i++) {
//         result+= "A"+int2str(i);
//         if (bracket[i] == true) {
//             result+=")(";
//         }
//     }
//     return result+")";
// }

void Traceback2(int i, int j, int** s) {
    if (i == j) return;
    if (i > j) {
        cout<<"错误: i>j"<<endl;
    }
    Traceback2(i, s[i][j], s);
    Traceback2(s[i][j]+1,j,s);
    cout<<"Multiply [A"<<i<<":"<<s[i][j]<<"]";
    cout<<"and [A"<<(s[i][j]+1)<<":"<<j<<"]"<<endl;
}

string int2str(int &i) {
    string s;
    stringstream ss(s);
    ss << i;
    return ss.str();
}