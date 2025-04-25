#include <bits/stdc++.h>
using namespace std;

int main(){
int N, i, j, nota;

cin >> N;

vector<int> Rank(N);
vector<int> Rank2(N);
vector<int> Rankf(N);

for (i=0; i < N; i++){
cin >> nota;

Rank.push_back(nota);
Rank2.push_back(nota);
}

sort(Rank.begin(), Rank.end(), [](int a, int b){
return a > b;
});

for (i=0; i < N; i++){
    for (j=0; j < N; j++){
if (Rank2[i] == Rank[j]){


    Rankf.push_back(j);
    break;
}

}
}

for (i=0; i< Rankf.size(); i++){
cout << Rankf[i] << " "; 
}

cout << endl;

return 0;

}
