/*
 * @Author: zhanghao 
 * @Date: 2018-11-25 00:50:39 
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-11-25 23:44:17
 */

#include<iostream>
#include<vector>
#include<iomanip>
#include<sstream>
using namespace std;
#define moreThanSolutionSum 1000000 // 不知道解的个数, 但这个数字不能小于解的个数
#define FROM 1
#define TO 12 // 14还能算，15要很久

string int2str(int &i) {
    string s;
    stringstream ss(s);
    ss << i;
    return ss.str();
}

class CQueen {
    public:
    CQueen(int n);
    ~CQueen();
    bool place(int k);
    void backtrace(int t);
    string getFirstSolution();
    void printAllSolution();
    int queenSum;
    int* tempSolution;
    int** allSolution;
    int solutionSum;
};

CQueen::CQueen(int n) {
    queenSum = n;
    solutionSum = 0;
    tempSolution = new int[n+1];
    allSolution = new int*[moreThanSolutionSum]; 
}

CQueen::~CQueen(void) {
    delete tempSolution;
    for (int i=1;i<=solutionSum;i++) {
        delete allSolution[i];
    }
    delete allSolution;
    queenSum = 0;
    solutionSum = 0;
}

bool CQueen::place(int k) {
    for (int i=1;i<k;i++) {
        if (abs(k-i) == abs(tempSolution[k]-tempSolution[i])||tempSolution[k] == tempSolution[i]) {
            return false;
        }
    }
    return true;
}

void CQueen::backtrace(int t) {
    if (t > queenSum) {
        solutionSum++;
        allSolution[solutionSum] = new int[queenSum+1];
        for (int i=1;i<=queenSum;i++) {
            allSolution[solutionSum][i] = tempSolution[i];
        }
    }else {
        for (int i=1;i<=queenSum;i++) {
            tempSolution[t] = i;
            if (place(t)) {
                backtrace(t+1);
            }
        }
    }
}

string CQueen::getFirstSolution(void) {
    if (solutionSum == 0) {
        return "null";
    }
    string s = "("+int2str(allSolution[1][1]);
    for (int i=2;i<=queenSum;i++) {
        s += ","+int2str(allSolution[1][i]);
    }
    return s += ")";
}

void CQueen::printAllSolution(void) {
    cout<<queenSum<<"皇后问题共有"<<solutionSum<<"种解答:"<<endl;
    for (int i=1;i<=solutionSum;i++) {
        for (int j=1;j<=queenSum;j++) {
            cout<<"  "<<allSolution[i][j];
        }
        cout<<endl;
    }
}

void Draw_line(vector<int> max,int columns) {  //画行线
	for (int i = 0; i < columns; i++) {
		cout << "+-";
		for (int j = 0; j <= max[i]; j++) {
			cout << '-';
		}
	}
	cout << '+' << endl;
}


void Draw_Datas(vector<int> max, vector<vector<string> > Str,vector<string> D,int columns,int row) { //显示构造过程，状态转换矩阵
	Draw_line(max,columns);
	for (int i = 0; i < D.size(); i++) {
		cout << "| " << setw(max[i]) << setiosflags(ios::left) << setfill(' ') << D[i] << ' ';
	}
	cout << '|' << endl;
	Draw_line(max, columns);
	for (int i = 0; i < row; i++) {
		for (int j = 0; j < columns; j++) {
			cout << "| " << setw(max[j]) << setiosflags(ios::left) << setfill(' ');
			cout << Str[i][j] << ' ';
		}
		cout << '|'<<endl;
	}
	Draw_line(max, columns);
}

int main(int argc, char const *argv[])
{
    // 打印n个皇后的所有解答
    CQueen q(5);
    q.backtrace(1);
    q.printAllSolution();

    // 打印m～n个皇后的解答个数与第一个解答
    vector<string>tableHeader;
    tableHeader.push_back("n queen");
    tableHeader.push_back("solution count");
    tableHeader.push_back("first solution");
    vector<int>colWidth;
    colWidth.push_back(10);
    colWidth.push_back(16);
    colWidth.push_back(40);
    int colsCount = 3;
    int rowsCount = TO - FROM +1;
    vector<vector<string> >content;
    for (int i=FROM;i<=TO;i++) {
        CQueen q(i);
        q.backtrace(1);
        vector<string>row;
        row.push_back(int2str(i));
        row.push_back(int2str(q.solutionSum));
        row.push_back(q.getFirstSolution());
        content.push_back(row);
    }
    Draw_Datas(colWidth,content,tableHeader,colsCount,rowsCount);
    return 0;
}