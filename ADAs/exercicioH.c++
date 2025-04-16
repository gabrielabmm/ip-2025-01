#include <bits/stdc++.h>
using namespace std;

int main(){
int t, i, n, a, valorMin, valorNovo, multiplicacao, prod;
vector <int> A;

cin >> t;

    for (i=1; i<= t; i++){
cin >> n;

A.clear();

for (int j=1; j<= n; j++){
cin >> a;

A.push_back(a);

}
valorMin = A[0];

for (int l=0; l< A.size(); l++){
if (A[l] == 0){
A[l] = A[l] + 1;
valorMin = A[l];
}

else if (valorMin > A[l]){
    valorMin = A[l];

} else {
}
} if (valorMin > 1){
}else {
    valorNovo = valorMin + 1;
}


multiplicacao = 1;
for (int k=0; k< A.size(); k++){
multiplicacao *= A[k];
}

prod = (multiplicacao / valorMin) * valorNovo;

cout << prod;}

    }
