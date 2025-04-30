#include <bits/stdc++.h>
using namespace std;

int main(){
int n, an, i, contador, j, valor;
vector<int> S;
vector<int> Sn;

cin >> n;

for (i=0; i< n; i++){
 cin >> an;

S.push_back(an);
contador++;
}

for (i=0; i< S.size(); i++){
    for (j=0; contador - j < i; j++){
if (S[i] > S[contador-j]){
valor = S[i] - S[contador-j];
S[i]= S[i] - valor;
S[contador-j]= S[contador-j] + valor;

}

}

}

for (i=0; i< S.size(); i++){
cout << S[i] << " ";
}

return 0;}
