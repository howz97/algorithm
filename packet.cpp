/*
 * @Author: zhanghao 
 * @Date: 2018-11-25 22:43:27 
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-11-26 15:38:27
 */

#include<iostream>
#include<vector>
#include <algorithm>
#include<iomanip>
#include<fstream>
#include<sstream>
using namespace std;
#define pktCap 500

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

float str2float(string num){
    float res;
    stringstream stream(num);
    stream>>res;
    return res;
}


class Ware {
    public:
    int no;
    int weight;
    int value;
    float vPerW;
    bool loaded;
    Ware();
    Ware(int n, int w,int v);
    ~Ware();
    void calcultVPerW();
    void load();
};

Ware::Ware() {
    no = 0;
    weight = 0;
    value = 0;
    vPerW = 0;
    loaded = false;
}

Ware::Ware(int n, int w, int v) {
    no = n;
    weight = w;
    value = v;
    vPerW = 0;
    loaded = false;
}

Ware::~Ware() {
    no = 0;
    weight = 0;
    value = 0;
    vPerW = 0;
    loaded = false;
}

void Ware::calcultVPerW() {
    vPerW = float(value)/float(weight);
}

void Ware::load() {
    loaded = true;
}

vector<Ware> readWares(string filename) {
    ifstream in(filename);
    if (!in) {
        cout<<"错误：文件找不到"<<endl;
        vector<Ware> wares;
        return wares;
    }
    string temp;
    if (!getline(in, temp)) {
        cout<<"错误：空文件"<<endl;
        vector<Ware> wares;
        return wares;
    }

    // 计算除去第一行，有多少行（当有多余换行符此结果将偏大）
    int wareCount = 0;
    for (; getline(in, temp);wareCount++) {}
    if (wareCount == 0) {
        cout<<"错误：无物品"<<endl;
        vector<Ware> wares;
        return wares;
    }

    // 重新读取
    vector<Ware> wares;
    in.close();
    ifstream in2(filename);
    getline(in2, temp);
    int weightWidth = temp.find("value");
    for (int i=0; getline(in2, temp);i++) {
        if (temp.empty()) {
            continue;
        }
        Ware w;
        w.no = i;
        w.weight = str2float(temp.substr(0,weightWidth));
        w.value = str2float(temp.substr(weightWidth));
        wares.push_back(w);
    }
    return wares;
}

void displayAllWare(vector<Ware> &wares) {
    int numWare = wares.size();
    if (numWare < 1) {
        return;
    }

    vector<string>tableHeader;
    tableHeader.push_back("ware id");
    tableHeader.push_back("weight");
    tableHeader.push_back("value");
    vector<int>colWidth;
    colWidth.push_back(10);
    colWidth.push_back(10);
    colWidth.push_back(10);
    int colsCount = 3;
    int rowsCount = numWare;
    vector<vector<string> >content;
    for (int i=0;i<numWare;i++) {
        vector<string>row;
        row.push_back(to_string(wares.at(i).no));
        row.push_back(to_string(wares.at(i).weight));
        row.push_back(to_string(wares.at(i).value));
        content.push_back(row);
    }
    Draw_Datas(colWidth,content,tableHeader,colsCount,rowsCount);
}

void displayLoadedWare(vector<Ware> &wares) {
    int numWare = wares.size();
    if (numWare < 1) {
        return;
    }

    vector<string>tableHeader;
    tableHeader.push_back("ware id");
    tableHeader.push_back("weight");
    tableHeader.push_back("value");
    vector<int>colWidth;
    colWidth.push_back(10);
    colWidth.push_back(10);
    colWidth.push_back(10);
    int colsCount = 3;
    int rowsCount = 0;
    int totalWeight = 0;
    int totalValue = 0;
    vector<vector<string> >content;
    for (int i=0;i<numWare;i++) {
        if (wares.at(i).loaded) {
            vector<string>row;
            row.push_back(to_string(wares.at(i).no));
            row.push_back(to_string(wares.at(i).weight));
            row.push_back(to_string(wares.at(i).value));
            content.push_back(row);
            totalWeight += wares.at(i).weight;
            totalValue += wares.at(i).value;
            rowsCount++;
        }
    }
    Draw_Datas(colWidth,content,tableHeader,colsCount,rowsCount);
    cout<<"total: "<<rowsCount<<" wares.  weight: "<<totalWeight<<" value: "<<totalValue<<endl;
}

void calculateVPerW(vector<Ware> &wares) {
    int numWare = wares.size();
    if (numWare < 1) {
        return;
    }

    for (int i=0;i<numWare;i++) {
        wares[i].calcultVPerW();
    }
}

bool greater_ware(const Ware &w1, const Ware &w2) {
        return w1.vPerW > w2.vPerW;
}

void load(vector<Ware> &wares, float loadLimit) {
    sort(wares.begin(), wares.end(),greater_ware);
    int numWare = wares.size();

    for (int i=0;i<numWare;i++) {
        if (wares.at(i).weight <= loadLimit) {
            wares.at(i).load();
            loadLimit -= wares.at(i).weight;
        }
    }
}

int main() {
    vector<Ware> wares = readWares("./packet.txt");
    if (wares.empty()) {
        return 1;
    }
    cout<<"全部物品:"<<endl;
    displayAllWare(wares);
    calculateVPerW(wares);
    load(wares, pktCap);
    cout<<endl;
    cout<<"背包内物品:"<<endl;
    displayLoadedWare(wares);
}