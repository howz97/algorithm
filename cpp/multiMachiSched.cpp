/*
 * @Author: zhanghao 
 * @Date: 2018-11-26 13:47:41 
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-11-26 17:58:04
 */

#include<iostream>
#include<vector>
#include<algorithm>
#include<fstream>
#include<sstream>
#include<string>
#include<iomanip>
using namespace std;
#define NUM_MACHINE 5

float str2int(string num){
    int res;
    stringstream stream(num);
    stream>>res;
    return res;
}

class Job {
    public:
    int no;
    int time;
    Job();
    Job(int n, int t);
    ~Job();
};

Job::Job() {
    no = 0;
    time = 0;
}

Job::Job(int n, int t) {
    no = n;
    time = t;
}

Job::~Job() {
    no = 0;
    time = 0;
}

class Machine {
    public:
    int no;
    int totalTime;
    vector<Job>jobs;
    Machine();
    Machine(int n);
    ~Machine();
    void insert(Job &job);
    void displayJobs();
};

Machine::Machine() {
    no = 0;
    totalTime = 0;
    vector<Job> jos;
    jobs = jos;
}

Machine::Machine(int n) {
    no = n;
    totalTime = 0;
    vector<Job> jos;
    jobs = jos;
}

Machine::~Machine() {
    no = 0;
    totalTime = 0;
    jobs.clear();
}

void Machine::insert(Job &job) {
    totalTime += job.time;
    jobs.push_back(job);
}

void Machine::displayJobs() {
    cout<<endl;
    cout<<"Machine "<<no<<" :";
    string temp;
    for (int i=0; i<jobs.size();i++) {
        if (i%7 == 0) {
            cout<<endl;
        }
        cout<<setw(16)<<setiosflags(ios::left)<< setfill(' ');
        temp = "[job "+to_string(jobs.at(i).no)+", "+to_string(jobs.at(i).time)+"s]";
        cout<<temp;
    }
    cout<<endl;
    cout<<"Total: "<<jobs.size()<<" jobs, "<<totalTime<<" seconds."<<endl;
}

vector<Job> importJobs(string filename) {
    ifstream in(filename);
    if (!in) {
        cout<<"错误：文件找不到"<<endl;
        vector<Job> jobs;
        return jobs;
    }
    string temp;
    if (!getline(in, temp)) {
        cout<<"错误：空文件"<<endl;
        vector<Job> jobs;
        return jobs;
    }
    int jobCount = 0;
    for (; getline(in, temp);jobCount++) {}
    if (jobCount == 0) {
        cout<<"错误：至少一个任务"<<endl;
        vector<Job> jobs;
        return jobs;
    }

    // 重新读取
    vector<Job> jobs;
    in.close();
    ifstream in2(filename);
    getline(in2, temp);
    for (int i=0; getline(in2, temp);i++) {
        if (temp.empty()) {
            continue;
        }
        Job job;
        job.no = i;
        job.time = str2int(temp);
        jobs.push_back(job);
    }
    return jobs;
}

vector<Machine> generateMachines(int num) {
    if (num < 1) {
        cout<<"错误：至少一台处理器"<<endl;
        vector<Machine> machines;
        return machines;
    }
    vector<Machine> machines;
    for (int i=0; i<num; i++) {
        Machine m(i);
        machines.push_back(m);
    }
    return machines;
}

bool greater_job(const Job &j1, const Job &j2) {
        return j1.time > j2.time;
}

bool less_time_machine(const Machine &m1, const Machine &m2) {
    return m1.totalTime < m2.totalTime;
}

void start(vector<Machine> machs, vector<Job> jobs) {
    if (jobs.size() <= machs.size()) {
        for (int i=0; i<jobs.size(); i++) {
            machs.at(i).insert(jobs.at(i));
        }
    }else {
        sort(jobs.begin(), jobs.end(), greater_job);
        for (int i=0; i<machs.size(); i++) {
            machs.at(i).insert(jobs.at(i));
        }
        for (int i=machs.size(); i<jobs.size();i++) {
            machs.at(0).insert(jobs.at(i));
            sort(machs.begin(),machs.end(),less_time_machine);
        }
    }

    // start to print handle information
    cout<<"Machines handle information (second):"<<endl;
    for (int i=0; i<machs.size();i++) {
        machs.at(i).displayJobs();
    }
    cout<<endl;
    cout<<"====> All job can be finished in "<<machs.at(machs.size()-1).totalTime<<" seconds."<<endl;
}

int main()
{
    vector<Job> jobs = importJobs("./multiMachiSched.txt");
    vector<Machine> machs = generateMachines(NUM_MACHINE);
    start(machs, jobs);
    return 0;
}
