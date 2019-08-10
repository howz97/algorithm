/*
 * @Author: zhanghao 
 * @Date: 2018-11-23 16:11:06 
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-11-24 00:55:34
 */

#include<iostream>
#include<fstream>
#include<string>
using namespace std;

void generateAuxiliaryTable(char* seq1, char* seq2, int seq1Num, int seq2Num, int** &lcsLen, int** &dir);
string  LCSstring(char* seq1,int i, int j,int** dir);

int main(int argc, char const *argv[])
{
    // 此程序仅仅支持ASICC编码
    ifstream in("./lcs.txt");
    if (!in) {
        cout<<"错误：文件找不到"<<endl;
    }
    string seqString[2];
    char** seq = new char*[2];
    int seqNum[2];
    for (int i=0;i<2;i++) {
        if (!getline(in, seqString[i])) {
            cout<<"错误：读取序列"<<i+1<<endl; // seqString[0]代表序列1
        }
        seq[i] = new char[1];
        seqString[i] = " "+seqString[i];
        strcpy(seq[i], seqString[i].c_str());
        seqNum[i] = seqString[i].length()-1; // 第一个字符的索引为0,此程序跳过它，所以字符数量减1
    }
    int** lcsLen;
    int** dir;
    generateAuxiliaryTable(seq[0],seq[1],seqNum[0],seqNum[1],lcsLen,dir);
    string lcs = LCSstring(seq[0], seqNum[0], seqNum[1], dir);
    cout<<"LCS: "<<lcs<<"   长度:"<<lcsLen[seqNum[0]][seqNum[1]]<<endl;
    return 0;
}

void generateAuxiliaryTable(char* seq1, char* seq2, int seq1Num, int seq2Num, int** &lcsLen, int** &dir) {
    lcsLen = new int*[seq1Num+1];
    for (int i=0;i<=seq1Num;i++) {
        lcsLen[i] = new int[seq2Num+1];
    }
    dir = new int*[seq1Num+1];
    for (int i=0;i<=seq1Num;i++) {
        dir[i] = new int[seq2Num+1];
    }
    for (int i=0;i<=seq1Num;i++) {
        lcsLen[i][0] = 0;
    }
    for (int j=0;j<=seq2Num;j++) {
        lcsLen[0][j] = 0;
    }
    for (int i=1;i<=seq1Num;i++) {
        for (int j=1;j<=seq2Num;j++) {
            if (seq1[i] == seq2[j]) {
                lcsLen[i][j] = lcsLen[i-1][j-1] +1;
                dir[i][j] = 1;
            }else if (lcsLen[i][j-1]>lcsLen[i-1][j]) {
                lcsLen[i][j] = lcsLen[i][j-1];
                dir[i][j] = 0;
            }else {
                lcsLen[i][j] = lcsLen[i-1][j];
                dir[i][j] = 2;
            }
        }
    }   
}

string  LCSstring(char* seq1,int i, int j,int** dir) {
    if (i==0||j==0) {
        return "";
    }
    if (dir[i][j] == 1) {
        return LCSstring(seq1,i-1,j-1,dir)+seq1[i];
    }else if (dir[i][j] == 0) {
        return LCSstring(seq1, i, j-1, dir);
    }else {
        return LCSstring(seq1, i-1,j, dir);
    }
}